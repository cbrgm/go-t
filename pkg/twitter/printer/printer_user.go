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

// UserPrinter represents a user printer
type UserPrinter interface {
	PrintAll(users []twitter.User)
	Print(user twitter.User)
}

// PrintUsers prints an array of users
func PrintUsers(users []twitter.User, opts *PrintOptions) {

	if opts.AsList {
		opts.Output = "list"
	}

	sortUsers(users, opts.Sort)

	var printer UserPrinter
	switch opts.Output {
	case "json":
		printer = new(UserJSONPrinter)
	case "yaml":
		printer = new(UserYAMLPrinter)
	case "list":
		printer = new(UserListPrinter)
	default:
		printer = new(UserTextPrinter)
	}

	// print tail if set
	if opts.Tail != 0 && opts.Tail > 0 {
		users = users[len(users)-opts.Tail:]
		printer.PrintAll(users)
		return
	}

	// print head if set
	if opts.Head != 0 && opts.Head > 0 {
		users = users[:opts.Head]
		printer.PrintAll(users)
		return
	}

	// print all
	printer.PrintAll(users)
}

// PrintUser prints a user
func PrintUser(user twitter.User, opts *PrintOptions) {

	if opts.AsList {
		opts.Output = "list"
	}

	var printer UserPrinter
	switch opts.Output {
	case "json":
		printer = new(UserJSONPrinter)
	case "yaml":
		printer = new(UserYAMLPrinter)
	case "list":
		printer = new(UserListPrinter)
	default:
		printer = new(UserTextPrinter)
	}
	printer.Print(user)
}

// UserTextPrinter represents a user text printer
type UserTextPrinter struct{}

// PrintAll prints an array of users
func (t *UserTextPrinter) PrintAll(users []twitter.User) {
	for _, tweet := range users {
		t.Print(tweet)
	}
}

// Print prints a user
func (t *UserTextPrinter) Print(user twitter.User) {
	heading := fmt.Sprintf("@%s, %d Tweets, ❤️ %d, %d Followers\n%s", user.Name, user.StatusesCount, user.FavouritesCount, user.FollowersCount, user.Description)
	heading = out.Colorize("Info", heading)

	// Print output
	out.Println(heading + "\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	lastTweeted := ""
	if user.Status != nil {
		lastTweeted = "last tweeted: \t" + toLocalTime(user.Status.CreatedAt)
	}

	fmt.Fprintf(w, "joined:\t%s\t\n", toLocalTime(user.CreatedAt))
	fmt.Fprintf(w, "%s\t\n", lastTweeted)
	fmt.Fprintf(w, "screen name:\t%s\t\n", user.ScreenName)
	fmt.Fprintf(w, "email:\t%s\t\n", user.Email)
	fmt.Fprintf(w, "homepage:\t%s\t\n", user.URL)
	fmt.Fprintf(w, "location:\t%s\t\n", user.Location)
	fmt.Fprintf(w, "language:\t%s\t\n", user.Lang)
	fmt.Fprintf(w, "verified:\t%t\t\n", user.Verified)
	fmt.Fprintf(w, "following:\t%t\t\n\n", user.Following)
	w.Flush()
}

// UserJSONPrinter represents a user json printer
type UserJSONPrinter struct{}

// PrintAll prints an array of users
func (t *UserJSONPrinter) PrintAll(users []twitter.User) {
	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// Print prints a user
func (t *UserJSONPrinter) Print(users twitter.User) {
	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// UserYAMLPrinter represents a user yaml printer
type UserYAMLPrinter struct{}

// PrintAll an array of users
func (t *UserYAMLPrinter) PrintAll(users []twitter.User) {
	b, err := yaml.Marshal(users)
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// Print prints a user
func (t *UserYAMLPrinter) Print(users twitter.User) {
	b, err := yaml.Marshal(users)
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// UserListPrinter represents a user list printer
type UserListPrinter struct{}

// PrintAll prints an array of users
func (t *UserListPrinter) PrintAll(users []twitter.User) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, user := range users {
		lastTweeted := ""
		if user.Status != nil {
			lastTweeted = "tweeted: " + toLocalTime(user.Status.CreatedAt)
		}
		s := fmt.Sprintf("%d\t%s\t@%s\t%d tweets\t❤️ %d\t%d followers\t%s\t", user.ID, toLocalTime(user.CreatedAt), user.ScreenName, user.StatusesCount, user.FavouritesCount, user.FollowersCount, lastTweeted)
		fmt.Fprintln(w, s)
	}
	w.Flush()
}

// Print prints a user
func (t *UserListPrinter) Print(user twitter.User) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	lastTweeted := ""
	if user.Status != nil {
		lastTweeted = "tweeted: " + toLocalTime(user.Status.CreatedAt)
	}
	s := fmt.Sprintf("%d\t%s\t@%s\t%d tweets\t❤️ %d\t%d followers\t%s\t", user.ID, toLocalTime(user.CreatedAt), user.ScreenName, user.StatusesCount, user.FavouritesCount, user.FollowersCount, lastTweeted)
	fmt.Fprintln(w, s)
	w.Flush()
}
