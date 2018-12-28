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

// TweetPrinter represents a tweet printer
type TweetPrinter interface {
	PrintAll(tweets []twitter.Tweet)
	Print(tweets twitter.Tweet)
}

// PrintTweets prints an array of tweets
func PrintTweets(tweets []twitter.Tweet, opts *PrintOptions) {

	if opts.AsList {
		opts.Output = "list"
	}

	sortTweets(tweets, opts.Sort)

	var printer TweetPrinter
	switch opts.Output {
	case "json":
		printer = new(TweetJSONPrinter)
	case "yaml":
		printer = new(TweetYAMLPrinter)
	case "list":
		printer = new(TweetListPrinter)
	default:
		printer = new(TweetTextPrinter)
	}

	// print tail if set
	if opts.Tail != 0 && opts.Tail > 0 {
		tweets = tweets[len(tweets)-opts.Tail:]
		printer.PrintAll(tweets)
		return
	}

	// print head if set
	if opts.Head != 0 && opts.Head > 0 {
		tweets = tweets[:opts.Head]
		printer.PrintAll(tweets)
		return
	}

	// print all
	printer.PrintAll(tweets)
}

// PrintTweet prints a tweet
func PrintTweet(tweet twitter.Tweet, opts *PrintOptions) {

	if opts.AsList {
		opts.Output = "list"
	}

	var printer TweetPrinter
	switch opts.Output {
	case "json":
		printer = new(TweetJSONPrinter)
	case "yaml":
		printer = new(TweetYAMLPrinter)
	case "list":
		printer = new(TweetListPrinter)
	default:
		printer = new(TweetTextPrinter)
	}
	printer.Print(tweet)
}

// TweetTextPrinter represents a tweet text printer
type TweetTextPrinter struct{}

// PrintAll prints an array of tweets
func (t *TweetTextPrinter) PrintAll(tweets []twitter.Tweet) {

	// newest first
	for i := len(tweets)/2 - 1; i >= 0; i-- {
		opp := len(tweets) - 1 - i
		tweets[i], tweets[opp] = tweets[opp], tweets[i]
	}

	for _, tweet := range tweets {
		t.Print(tweet)
	}
}

// Print prints a tweet
func (t *TweetTextPrinter) Print(tweet twitter.Tweet) {
	heading := fmt.Sprintf("@%s (%s), %s # %d \n❤️ %d, %d RTs\n", tweet.User.ScreenName, tweet.User.Name, toLocalTime(tweet.CreatedAt), tweet.ID, tweet.FavoriteCount, tweet.RetweetCount)
	heading = out.Colorize("Info", heading)

	// Print output
	out.Println(heading + tweet.FullText + "\n")
}

// TweetJSONPrinter represents a tweet json printer
type TweetJSONPrinter struct{}

// PrintAll prints an array of tweets
func (t *TweetJSONPrinter) PrintAll(tweets []twitter.Tweet) {
	b, err := json.MarshalIndent(tweets, "", "  ")
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// Print prints a tweet
func (t *TweetJSONPrinter) Print(tweet twitter.Tweet) {
	b, err := json.MarshalIndent(tweet, "", "  ")
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// TweetYAMLPrinter represents a tweet yaml printer
type TweetYAMLPrinter struct{}

// PrintAll prints an array of tweets
func (t *TweetYAMLPrinter) PrintAll(tweets []twitter.Tweet) {
	b, err := yaml.Marshal(tweets)
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// Print prints a tweet
func (t *TweetYAMLPrinter) Print(tweet twitter.Tweet) {
	b, err := yaml.Marshal(tweet)
	if err != nil {
		// error
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

// TweetListPrinter represents a tweet list printer
type TweetListPrinter struct{}

// PrintAll prints an array of tweets
func (t *TweetListPrinter) PrintAll(tweets []twitter.Tweet) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, tweet := range tweets {
		s := fmt.Sprintf("%d\t%s\t@%s\t%s\t❤️ %d\t%d RTs", tweet.ID, toLocalTime(tweet.CreatedAt), tweet.User.ScreenName, abbreviate(tweet.FullText, 30), tweet.FavoriteCount, tweet.RetweetCount)
		fmt.Fprintln(w, s)
	}
	w.Flush()
}

// Print prints a tweet
func (t *TweetListPrinter) Print(tweet twitter.Tweet) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	s := fmt.Sprintf("%d\t%s\t%s\t@%s\t❤️ %d\t%d RTs", tweet.ID, toLocalTime(tweet.CreatedAt), tweet.User.ScreenName, abbreviate(tweet.FullText, 30), tweet.FavoriteCount, tweet.RetweetCount)
	fmt.Fprintln(w, s)
	w.Flush()
}
