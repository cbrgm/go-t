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

// LookupStatusOptions represents options for the lookup action
type LookupStatusOptions struct {
	Output      string
	AsList      bool
	isConfirmed bool
	File        string
	ID          []int64
	TrimUser    bool
}

// LookupTweet returns a slice of tweets for a list of IDs
func (c *Client) LookupTweet(opts *LookupStatusOptions) {
	params := twitter.StatusLookupParams{
		ID:        opts.ID,
		TrimUser:  &opts.TrimUser,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	tweets, _, err := c.api.Statuses.Lookup(opts.ID, &params)
	if err != nil {
		out.Fatalln(err)
	}

	printer.PrintTweets(tweets, printOpts)
}

// RemoveStatusOptions represents options for the remove action
type RemoveStatusOptions struct {
	Output   string
	AsList   bool
	SkipYes  bool
	File     bool
	ID       int64
	TrimUser bool
	Verbose  bool
}

// RemoveTweet deletes a tweet
func (c *Client) RemoveTweet(opts *RemoveStatusOptions) {
	params := twitter.StatusDestroyParams{
		ID:        opts.ID,
		TrimUser:  &opts.TrimUser,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	tweet, _, err := c.api.Statuses.Destroy(opts.ID, &params)
	if err != nil {
		out.Fatalln(err)
	}

	out.Infoln("✔ tweet successfully removed!")
	if opts.Verbose {
		printer.PrintTweet(*tweet, printOpts)
	}
}

// RetweetStatusOptions represents options for the retweet action
type RetweetStatusOptions struct {
	Output   string
	AsList   bool
	SkipYes  bool
	File     bool
	ID       int64
	TrimUser bool
	Verbose  bool
}

// RetweetTweet represents options for the retweet action
func (c *Client) RetweetTweet(opts *RetweetStatusOptions) {
	params := twitter.StatusRetweetParams{
		ID:        opts.ID,
		TrimUser:  &opts.TrimUser,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	tweet, _, err := c.api.Statuses.Retweet(opts.ID, &params)
	if err != nil {
		out.Fatalln(err)
	}

	out.Infoln("✔ tweet successfully retweeted!")
	if opts.Verbose {
		printer.PrintTweet(*tweet, printOpts)
	}

}

// RetweetsStatusOptions represents options for the retweet action
type RetweetsStatusOptions struct {
	Output   string
	AsList   bool
	File     bool
	ID       int64
	TrimUser bool
	Count    int
	Head     int
	Tail     int
	Sort     string
}

// RetweetsOfTweet shows all retweets of a given tweet
func (c *Client) RetweetsOfTweet(opts *RetweetsStatusOptions) {

	params := twitter.StatusRetweetsParams{
		Count:     opts.Count,
		ID:        opts.ID,
		TrimUser:  &opts.TrimUser,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Tail:   opts.Tail,
		Head:   opts.Head,
		Count:  opts.Count,
		Sort:   opts.Sort,
	}

	tweets, _, err := c.api.Statuses.Retweets(opts.ID, &params)
	if err != nil {
		out.Fatalln(err)
	}

	printer.PrintTweets(tweets, printOpts)

}

// UnretweetStatusOptions represents options for the unretweet action
type UnretweetStatusOptions struct {
	Output   string
	AsList   bool
	File     bool
	ID       int64
	TrimUser bool
	SkipYes  bool
	Verbose  bool
}

// UnretweetTweet unretweets a tweet
func (c *Client) UnretweetTweet(opts *UnretweetStatusOptions) {

	params := twitter.StatusUnretweetParams{
		ID:        opts.ID,
		TrimUser:  &opts.TrimUser,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	tweet, _, err := c.api.Statuses.Unretweet(opts.ID, &params)
	if err != nil {
		out.Fatalln(err)
	}

	out.Infoln("✔ tweet successfully unretweeted!")
	if opts.Verbose {
		printer.PrintTweet(*tweet, printOpts)
	}
}

// UpdateStatusOptions represents options for the update action
type UpdateStatusOptions struct {
	Output  string
	AsList  bool
	SkipYes bool
	Status  string
	Verbose bool
}

//UpdateTweet sends a new tweet
func (c *Client) UpdateTweet(opts *UpdateStatusOptions) {

	params := twitter.StatusUpdateParams{
		Status:    opts.Status,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	tweet, _, err := c.api.Statuses.Update(opts.Status, &params)
	if err != nil {
		out.Fatalln(err)
	}

	out.Infoln("✔ tweet successfully updated!")
	if opts.Verbose {
		printer.PrintTweet(*tweet, printOpts)
	}
}

// TweetSearchOptions represents search options for different tweets
type TweetSearchOptions struct {
	Output string
	AsList bool
	File   bool
	Query  string
	Count  int
	Head   int
	Tail   int
}

// SearchTweet searches for tweets by a given query
func (c *Client) SearchTweet(opts *TweetSearchOptions) {

	params := twitter.SearchTweetParams{
		Count:     opts.Count,
		Query:     opts.Query,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Head:   opts.Head,
		Tail:   opts.Tail,
		Count:  opts.Count,
	}

	search, _, err := c.api.Search.Tweets(&params)
	if err != nil {
		out.Fatalln(err)
	}

	printer.PrintTweets(search.Statuses, printOpts)
}
