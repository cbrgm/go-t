/*
 * Copyright 2018 Christian Bargmann <chris@cbrgm.net>
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package twitter

import (
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/cbrgm/go-t/pkg/twitter/printer"
	"github.com/dghubble/go-twitter/twitter"
)

// UserSearchOptions represents the user search options
type UserSearchOptions struct {
	Output string
	AsList bool
	File   bool
	Query  string
	Page   int
	Count  int
	Head   int
	Tail   int
}

// SearchUser searches for users by a given query
func (c *Client) SearchUser(opts *UserSearchOptions) {
	params := twitter.UserSearchParams{
		Count: opts.Count,
		Query: opts.Query,
		Page:  opts.Page,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Head:   opts.Head,
		Tail:   opts.Tail,
		Count:  opts.Count,
	}

	users, _, err := c.api.Users.Search(opts.Query, &params)
	if err != nil {
		out.Fatalln(err)
	}

	printer.PrintUsers(users, printOpts)
}

// WhoisUserOptions represents whois user options
type WhoisUserOptions struct {
	Output     string
	AsList     bool
	File       bool
	UserID     []int64
	ScreenName []string
	Head       int
	Tail       int
	Count      int
}

// WhoisUser provides detailed user information
func (c *Client) WhoisUser(opts *WhoisUserOptions) {
	params := twitter.UserLookupParams{
		UserID:     opts.UserID,
		ScreenName: opts.ScreenName,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Head:   opts.Head,
		Tail:   opts.Tail,
		Count:  opts.Count,
	}

	users, _, err := c.api.Users.Lookup(&params)
	if err != nil {
		out.Fatalln(err)
	}

	printer.PrintUsers(users, printOpts)
}
