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

// CreateFavoriteOptions represents the create favorite command line options
type CreateFavoriteOptions struct {
	Output  string
	AsList  bool
	File    bool
	ID      int64
	SkipYes bool
	Verbose bool
}

// CreateFavorite likes a tweet
func (c *Client) CreateFavorite(opts *CreateFavoriteOptions) {
	params := twitter.FavoriteCreateParams{
		ID: opts.ID,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	favorite, _, err := c.api.Favorites.Create(&params)
	if err != nil {
		out.Fatalln(err)
	}

	out.Infoln("✔ successfully liked tweet!")
	if opts.Verbose {
		printer.PrintTweet(*favorite, printOpts)
	}
}

// RemoveFavoriteOptions remove favorite options
type RemoveFavoriteOptions struct {
	Output  string
	AsList  bool
	SkipYes bool
	File    bool
	ID      int64
	Verbose bool
}

// RemoveFavorite dislikes a tweet
func (c *Client) RemoveFavorite(opts *RemoveFavoriteOptions) {
	params := twitter.FavoriteDestroyParams{
		ID: opts.ID,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	favorite, _, err := c.api.Favorites.Destroy(&params)
	if err != nil {
		out.Fatalln(err)
	}

	out.Infoln("✔ successfully disliked tweet!")
	if opts.Verbose {
		printer.PrintTweet(*favorite, printOpts)
	}
}

// ListFavoriteOptions represents list favorites options
type ListFavoriteOptions struct {
	Output     string
	AsList     bool
	File       bool
	UserID     int64
	ScreenName string
	Count      int
	Head       int
	Tail       int
	Sort       string
}

// ListFavorites lists favorites
func (c *Client) ListFavorites(opts *ListFavoriteOptions) {
	params := twitter.FavoriteListParams{
		UserID:     opts.UserID,
		ScreenName: opts.ScreenName,
		Count:      opts.Count,
		TweetMode:  "extended",
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Count:  opts.Count,
		Head:   opts.Head,
		Tail:   opts.Tail,
		Sort:   opts.Sort,
	}

	favorites, _, err := c.api.Favorites.List(&params)
	if err != nil {
		out.Fatalln(err)
	}

	printer.PrintTweets(favorites, printOpts)
}
