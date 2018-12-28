/*
 * Copyright 2018 Christian Bargmann <chris@cbrgm.net>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package printer

import (
	"encoding/json"
	"fmt"
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/dghubble/go-twitter/twitter"
	"gopkg.in/yaml.v2"
	"os"
	"text/tabwriter"
)

// RelationshipPrinter represents an relationship printer
type RelationshipPrinter interface {
	PrintAll(relationships []twitter.Relationship)
	Print(relationship twitter.Relationship)
}

// PrintRelationships prints an array of relationships
func PrintRelationships(relationships []twitter.Relationship, opts *PrintOptions) {

	if opts.AsList {
		opts.Output = "list"
	}

	var printer RelationshipPrinter
	switch opts.Output {
	case "json":
		printer = new(RelationshipJSONPrinter)
	case "yaml":
		printer = new(RelationshipYAMLPrinter)
	case "list":
		printer = new(RelationshipListPrinter)
	default:
		printer = new(RelationshipTextPrinter)
	}

	// print tail if set
	if opts.Tail != 0 && opts.Tail > 0 {
		relationships = relationships[len(relationships)-opts.Tail:]
		printer.PrintAll(relationships)
		return
	}

	// print head if set
	if opts.Head != 0 && opts.Head > 0 {
		relationships = relationships[:opts.Head]
		printer.PrintAll(relationships)
		return
	}

	// print all
	printer.PrintAll(relationships)
}

// PrintRelationship prints a relationship
func PrintRelationship(relationship twitter.Relationship, opts *PrintOptions) {

	if opts.AsList {
		opts.Output = "list"
	}

	var printer RelationshipPrinter
	switch opts.Output {
	case "json":
		printer = new(RelationshipJSONPrinter)
	case "yaml":
		printer = new(RelationshipYAMLPrinter)
	case "list":
		printer = new(RelationshipListPrinter)
	default:
		printer = new(RelationshipTextPrinter)
	}
	printer.Print(relationship)
}

// RelationshipTextPrinter represents an relationship text printer
type RelationshipTextPrinter struct{}

// PrintAll prints an array of relationships
func (t *RelationshipTextPrinter) PrintAll(relationships []twitter.Relationship) {
	for _, relationship := range relationships {
		t.Print(relationship)
	}
}

// Print prints a relationship
func (t *RelationshipTextPrinter) Print(relationship twitter.Relationship) {
	out.Printf("%s is following %s: %t\n", relationship.Source.ScreenName, relationship.Target.ScreenName, relationship.Source.Following)
	out.Printf("%s is following %s: %t\n", relationship.Target.ScreenName, relationship.Source.ScreenName, relationship.Target.Following)
	out.Printf("%s is blocking %s: %t\n", relationship.Source.ScreenName, relationship.Target.ScreenName, relationship.Source.Blocking)
	out.Printf("%s is muting %s: %t\n", relationship.Source.ScreenName, relationship.Target.ScreenName, relationship.Source.Blocking)
	out.Printf("%s is able to send direct messages to %s: %t\n", relationship.Source.ScreenName, relationship.Target.ScreenName, relationship.Source.CanDM)
}

// RelationshipJSONPrinter represents an relationship json printer
type RelationshipJSONPrinter struct{}

// PrintAll prints an array of relationships
func (t *RelationshipJSONPrinter) PrintAll(relationships []twitter.Relationship) {
	b, err := json.MarshalIndent(relationships, "", "  ")
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// Print prints a relationship
func (t *RelationshipJSONPrinter) Print(relationship twitter.Relationship) {
	b, err := json.MarshalIndent(relationship, "", "  ")
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// RelationshipYAMLPrinter represents an relationship yaml printer
type RelationshipYAMLPrinter struct{}

// PrintAll prints an array of relationships
func (t *RelationshipYAMLPrinter) PrintAll(relationships []twitter.Relationship) {
	b, err := yaml.Marshal(relationships)
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// Print prints a relationship
func (t *RelationshipYAMLPrinter) Print(relationship twitter.Relationship) {
	b, err := yaml.Marshal(relationship)
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// RelationshipListPrinter represents an relationship list printer
type RelationshipListPrinter struct{}

// PrintAll prints an array of relationships
func (t *RelationshipListPrinter) PrintAll(relationships []twitter.Relationship) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	for _, relationship := range relationships {
		s := fmt.Sprintf("%s\t%s\t", relationship.Source.ScreenName, relationship.Target.ScreenName)
		fmt.Fprintln(w, s)
	}
	w.Flush()
}

// Print prints a relationship
func (t *RelationshipListPrinter) Print(relationship twitter.Relationship) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	s := fmt.Sprintf("%s\t%s\t", relationship.Source.ScreenName, relationship.Target.ScreenName)
	fmt.Fprintln(w, s)
	w.Flush()
}
