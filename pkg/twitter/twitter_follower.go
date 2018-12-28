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

// ListFollowersOptions represents list follower options
type ListFollowersOptions struct {
	Output     string
	AsList     bool
	File       bool
	UserID     int64
	ScreenName string
	Head       int
	Tail       int
	Sort       string
}

// ListFollowers lists followersof a given user
func (c *Client) ListFollowers(opts *ListFollowersOptions) {

	// see: https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list.html
	params := &twitter.FollowerListParams{
		Cursor:     -1,
		Count:      200,
		ScreenName: opts.ScreenName,
		UserID:     opts.UserID,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
		Tail:   opts.Tail,
		Head:   opts.Tail,
		Sort:   opts.Sort,
	}

	result := make([]twitter.User, 0)

	for params.Cursor != 0 {
		followers, _, err := c.api.Followers.List(params)
		if err != nil {
			out.Fatalln(err)
		}

		result = append(result, followers.Users...)
		params.Cursor = followers.NextCursor
	}

	printer.PrintUsers(result, printOpts)
}
