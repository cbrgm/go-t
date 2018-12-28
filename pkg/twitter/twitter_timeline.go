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

package twitter

import (
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/cbrgm/go-t/pkg/twitter/printer"
	"github.com/dghubble/go-twitter/twitter"
)

// HomeTimelineOptions represents show hometimeline options
type HomeTimelineOptions struct {
	Tail           int
	Head           int
	AsList         bool
	Output         string
	File           string
	Follow         bool
	Count          int
	TrimUser       bool
	ExcludeReplies bool
	Sort           string
}

// HomeTimeline shows your home timeline
func (c *Client) HomeTimeline(opts *HomeTimelineOptions) {
	params := &twitter.HomeTimelineParams{
		Count:          opts.Count,
		TrimUser:       &opts.TrimUser,
		ExcludeReplies: &opts.ExcludeReplies,
		TweetMode:      "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Count:  opts.Count,
		Head:   opts.Head,
		Tail:   opts.Tail,
		Sort:   opts.Sort,
	}

	tweets, _, err := c.api.Timelines.HomeTimeline(params)
	if err != nil {
		out.Fatalf("Failed %s", err)
	}

	printer.PrintTweets(tweets, printOpts)
}

// MentionsTimelineOptions represents mentions timeline options
type MentionsTimelineOptions struct {
	Tail     int
	Head     int
	AsList   bool
	Output   string
	Follow   bool
	Count    int
	TrimUser bool
	Sort     string
}

// MentionsTimeline returns latest mentions of you by other users
func (c *Client) MentionsTimeline(opts *MentionsTimelineOptions) {
	params := &twitter.MentionTimelineParams{
		Count:     opts.Count,
		TrimUser:  &opts.TrimUser,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Count:  opts.Count,
		Head:   opts.Head,
		Tail:   opts.Tail,
		Sort:   opts.Sort,
	}

	tweets, _, err := c.api.Timelines.MentionTimeline(params)
	if err != nil {
		out.Fatalf("Failed %s", err)
	}

	printer.PrintTweets(tweets, printOpts)
}

// UserTimelineOptions represents user timeline options
type UserTimelineOptions struct {
	Tail            int
	Head            int
	AsList          bool
	Sort            string
	Output          string
	File            string
	Follow          bool
	Count           int
	TrimUser        bool
	ScreenName      string
	UserID          int64
	ExcludeReplies  bool
	IncludeRetweets bool
}

// UserTimeline returns the timeline of a given user
func (c *Client) UserTimeline(opts *UserTimelineOptions) {
	params := &twitter.UserTimelineParams{
		ScreenName:      opts.ScreenName,
		UserID:          opts.UserID,
		ExcludeReplies:  &opts.ExcludeReplies,
		IncludeRetweets: &opts.IncludeRetweets,
		Count:           opts.Count,
		TrimUser:        &opts.TrimUser,
		TweetMode:       "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Count:  opts.Count,
		Head:   opts.Head,
		Tail:   opts.Tail,
		Sort:   opts.Sort,
	}

	tweets, _, err := c.api.Timelines.UserTimeline(params)
	if err != nil {
		out.Fatalf("Failed %s", err)
	}

	printer.PrintTweets(tweets, printOpts)
}

// RetweetsOfMeTimelineOptions represents retweets of your tweets display options
type RetweetsOfMeTimelineOptions struct {
	Tail     int
	Head     int
	AsList   bool
	Output   string
	Follow   bool
	Count    int
	TrimUser bool
	Sort     string
}

// RetweetsOfMeTimeline shows retweets of your tweets
func (c *Client) RetweetsOfMeTimeline(opts *RetweetsOfMeTimelineOptions) {
	params := &twitter.RetweetsOfMeTimelineParams{
		Count:     opts.Count,
		TrimUser:  &opts.TrimUser,
		TweetMode: "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Count:  opts.Count,
		Head:   opts.Head,
		Tail:   opts.Tail,
		Sort:   opts.Sort,
	}

	tweets, _, err := c.api.Timelines.RetweetsOfMeTimeline(params)
	if err != nil {
		out.Fatalf("Failed %s", err)
	}

	printer.PrintTweets(tweets, printOpts)
}
