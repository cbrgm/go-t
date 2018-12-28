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

func sortUsers(users []twitter.User, arg string) {
	if arg == "" {
		return
	}

	sorting, order := parseSortString(arg)

	switch sorting {
	case "likes":
		sort.Slice(users, func(i, j int) bool {
			if order == "asc" {
				return users[i].FavouritesCount > users[j].FavouritesCount
			}
			return users[i].FavouritesCount < users[j].FavouritesCount
		})
	case "followers":
		sort.Slice(users, func(i, j int) bool {
			if order == "asc" {
				return users[i].FollowersCount > users[j].FollowersCount
			}
			return users[i].FollowersCount < users[j].FollowersCount
		})
	case "tweets":
		sort.Slice(users, func(i, j int) bool {
			if order == "asc" {
				return users[i].StatusesCount > users[j].StatusesCount
			}
			return users[i].StatusesCount < users[j].StatusesCount
		})
	case "usernames":
		sort.Slice(users, func(i, j int) bool {
			var si, sj = users[i].ScreenName, users[j].ScreenName
			var sil, sjl = strings.ToLower(si), strings.ToLower(sj)
			if sil == sjl {
				if order == "asc" {
					return si > sj
				}
				return si < sj
			}
			if order == "asc" {
				return sil > sjl
			}
			return sil < sjl
		})
	case "tweeted":
		sort.Slice(users, func(i, j int) bool {
			if users[i].Status == nil || users[j].Status == nil {
				return false
			}

			ti, err := createdAtTime(users[i].Status.CreatedAt)
			if err != nil {
				return false
			}
			tj, err := createdAtTime(users[j].Status.CreatedAt)

			if order == "asc" {
				return ti.After(tj)
			}
			return ti.Before(tj)
		})
	case "joined":
		sort.Slice(users, func(i, j int) bool {
			if users[i].Status == nil || users[j].Status == nil {
				return false
			}

			ti, err := createdAtTime(users[i].CreatedAt)
			if err != nil {
				return false
			}
			tj, err := createdAtTime(users[j].CreatedAt)

			if order == "asc" {
				return ti.After(tj)
			}
			return ti.Before(tj)
		})
	default:
		return
	}
}
