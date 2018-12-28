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

package in

import (
	"fmt"
	"strconv"
	"strings"
)

const maxTweetStatusLength = 280
const maxTweetIDLength = 19

// ParseTweetFromArgs parses tweet id from stdin
func ParseTweetFromArgs(arg string, isFile bool) (int64, error) {
	if isFile {
		return tweetFromStdin(arg)
	}

	var tweetID int64
	if IsValidTweetID(arg) {
		id, _ := strconv.ParseInt(arg, 10, 64)
		tweetID = id
	}
	if tweetID == 0 {
		return 0, ErrParseTweetFromStdin
	}
	return tweetID, nil
}

func tweetFromStdin(filename string) (int64, error) {
	var tweetID int64

	// read from stdin
	b, err := readFile(filename)
	if err != nil {
		fmt.Print(err)
		return 0, ErrParseTweetFromStdin
	}

	// get first element
	s := strings.Fields(string(b))[0]
	if len(s) == 0 {
		return 0, ErrParseTweetFromStdin
	}

	if IsValidTweetID(s) {
		id, _ := strconv.ParseInt(s, 10, 64)
		tweetID = id
	}
	if tweetID == 0 {
		return 0, ErrParseTweetFromStdin
	}

	return tweetID, nil
}

// IsValidTweetID checks wether the tweet id is valid or not
func IsValidTweetID(s string) bool {
	if s == "" || len(s) > maxTweetIDLength {
		return false
	}
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	return true
}

// ParseTweetStatusFromArgs parses tweet text from args
func ParseTweetStatusFromArgs(arg string, isFile bool) (status string, err error) {
	if isFile {
		return tweetStatusFromStdin(arg)
	}
	return arg, nil
}

func tweetStatusFromStdin(filename string) (string, error) {
	var status string

	// read from stdin
	b, err := readFile(filename)
	if err != nil {
		return "", ErrParseTweetFromStdin
	}

	s := string(b)
	if IsValidStatus(s) {
		status = s
	}
	if status == "" {
		return "", ErrParseTweetFromStdin
	}

	return status, nil
}

// IsValidStatus returns wether the status is valid
func IsValidStatus(s string) bool {
	if s == "" {
		return false
	}
	return len(s) <= maxTweetStatusLength
}
