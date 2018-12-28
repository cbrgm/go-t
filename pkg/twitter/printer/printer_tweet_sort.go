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
	"github.com/dghubble/go-twitter/twitter"
	"sort"
	"strings"
)

func sortTweets(tweets []twitter.Tweet, arg string) {
	if arg == "" {
		return
	}

	sorting, order := parseSortString(arg)

	switch sorting {
	case "likes":
		sort.Slice(tweets, func(i, j int) bool {
			if order == "asc" {
				return tweets[i].FavoriteCount > tweets[j].FavoriteCount
			}
			return tweets[i].FavoriteCount < tweets[j].FavoriteCount
		})
	case "retweets":
		sort.Slice(tweets, func(i, j int) bool {
			if order == "asc" {
				return tweets[i].RetweetCount > tweets[j].RetweetCount
			}
			return tweets[i].RetweetCount < tweets[j].RetweetCount
		})
	case "usernames":
		sort.Slice(tweets, func(i, j int) bool {
			var si, sj = tweets[i].User.ScreenName, tweets[j].User.ScreenName
			var sil, sjl = strings.ToLower(si), strings.ToLower(sj)
			if sil == sjl {
				if order == "asc" {
					return si < sj
				}
				return si > sj
			}
			if order == "asc" {
				return sil < sjl
			}
			return sil > sjl
		})
	case "tweeted":
		sort.Slice(tweets, func(i, j int) bool {
			ti, err := createdAtTime(tweets[i].CreatedAt)
			if err != nil {
				return false
			}
			tj, err := createdAtTime(tweets[j].CreatedAt)

			if order == "asc" {
				return ti.After(tj)
			}
			return ti.Before(tj)
		})
	default:
		return
	}
}
