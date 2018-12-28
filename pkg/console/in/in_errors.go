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

import "errors"

var (
	// ErrParseUserFromStdin is returned when the user alias from stdin cannot be parsed
	ErrParseUserFromStdin = errors.New("invalid user id or account name from stdin")
	// ErrParseTweetFromStdin is returned when the tweet id from stdin cannot be parsed
	ErrParseTweetFromStdin = errors.New("invalid tweet id from stdin")
	// ErrParseInput is returned if the input from stdin cannot be parsed
	ErrParseInput = errors.New("failed parse input from stdin")
	// ErrAbortInput is returned if the user aborted input action
	ErrAbortInput = errors.New("waiting for input was aborted")
	// ErrUnknownInput is returned if the user entered an unknown action
	ErrUnknownInput = errors.New("unknown answer entered")
)
