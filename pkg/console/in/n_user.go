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
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const maxUserIDLength = 19

// ParseUserFromArgs parses userid or/and name from args
func ParseUserFromArgs(arg string, isFile bool) (userID int64, name string, err error) {
	if isFile {
		return userFromStdin(arg)
	}
	return userFromArg(arg)
}

// userFromStdin parses a user from file or stdin and returns the user id and screenname for a given filename
// filename can either be `-` for stdin or a concrete path like `/home/user/foo/username.txt`
func userFromStdin(filename string) (int64, string, error) {
	var userID int64
	var username string

	// read from stdin
	b, err := readFile(filename)
	if err != nil {
		return 0, "", ErrParseUserFromStdin
	}

	// get first element
	s := strings.Fields(string(b))[0]
	if len(s) == 0 {
		return 0, "", ErrParseUserFromStdin
	}

	if IsValidUsername(s) {
		username = s
	}
	if IsValidUserID(s) {
		id, _ := strconv.ParseInt(s, 10, 64)
		userID = id
	}
	if userID == 0 && username == "" {
		return 0, "", ErrParseUserFromStdin
	}

	username = strings.TrimPrefix(username, "@")
	return userID, username, nil
}

// userFromArg returns the user id and screenname for a given arg
func userFromArg(arg string) (int64, string, error) {
	var userID int64
	var username string

	if arg == "" {
		return 0, "", ErrParseUserFromStdin
	}

	if IsValidUsername(arg) {
		username = arg
	}
	if IsValidUserID(arg) {
		id, _ := strconv.ParseInt(arg, 10, 64)
		userID = id
	}
	if userID == 0 && username == "" {
		return 0, "", ErrParseUserFromStdin
	}

	username = strings.TrimPrefix(username, "@")
	return userID, username, nil
}

// IsValidUsername returns wether the provided user name is valid or not
func IsValidUsername(s string) bool {
	if s == "" {
		return false
	}

	// match twitter names like @name_123
	r, err := regexp.Compile(`^[@](\w){1,15}$`)
	if err != nil {
		return false
	}
	return r.MatchString(s)
}

// IsValidUserID checks wether the user id is valid or not
func IsValidUserID(s string) bool {
	if s == "" || len(s) > maxUserIDLength {
		return false
	}
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func readFile(filename string) ([]byte, error) {
	if filename == "-" {
		return ioutil.ReadAll(os.Stdin)
	}
	return ioutil.ReadFile(filename)
}
