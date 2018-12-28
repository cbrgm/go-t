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
	"strings"
	"time"
	"unicode"
)

// TimeLayout represents the timelayout
const TimeLayout = "Mon Jan 02 15:04:05 -0700 2006"

// PrintOptions represents several display options passed via cli flags
type PrintOptions struct {
	Output string
	AsList bool
	Count  int
	Head   int
	Tail   int
	Sort   string
}

// toLocalTime converts a given date string to users local time
func toLocalTime(timeStr string) string {
	timeValue, err := time.Parse(TimeLayout, timeStr)
	if err != nil {
		return timeStr
	}
	return timeValue.Local().Format(TimeLayout)
}

// createdAtTime returns the time an object was created.
func createdAtTime(createdAt string) (time.Time, error) {
	return time.Parse(time.RubyDate, createdAt)
}

// abbreviate shortens a string to a given length and adds `...` to the strings suffix
func abbreviate(str string, num int) string {
	result := ""
	str = minify(str)
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		result = str[0:num] + "..."
	}
	return result
}

// minify removes double whitespaces and linebreaks from a given string
func minify(str string) (out string) {
	white := false
	for _, c := range str {
		if unicode.IsSpace(c) {
			if !white {
				out = out + " "
			}
			white = true
		} else {
			out = out + string(c)
			white = false
		}
	}
	return
}

func parseSortString(arg string) (sorting, order string) {
	if arg == "" {
		return "unknown", "asc"
	}
	s := strings.Split(arg, ",")
	if len(s) != 2 {
		return s[0], "asc"
	}

	sorting = s[0]
	order = s[1]

	if order == "asc" || order == "desc" {
		return sorting, order
	}
	return sorting, "asc"
}
