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

package in

import (
	"github.com/cbrgm/go-t/pkg/console/out"
	"os"
	"strings"
)

// AskForString requires user text input
func AskForString(text, def string, skip bool) (string, error) {
	if skip {
		return def, nil
	}

	out.Infof("%s [%s]: ", text, def)
	input, err := NewReader(os.Stdin).ReadLine()
	if err != nil {
		return def, ErrParseInput
	}

	input = strings.TrimSpace(input)
	if input == "" {
		input = def
	}

	return input, nil
}

// AskForBool requires user bool input
func AskForBool(text string, def bool, skip bool) (bool, error) {
	if skip {
		return def, nil
	}

	choices := "y-N-q"
	if def {
		choices = "Y-n-q"
	}

	str, err := AskForString(text, choices, skip)
	if err != nil {
		return false, ErrParseInput
	}

	switch str {
	case "Y-n-q":
		return true, nil
	case "y-N-q":
		return false, nil
	}

	str = strings.ToLower(string(str[0]))
	switch str {
	case "y":
		return true, nil
	case "n":
		return false, nil
	case "q":
		return false, ErrAbortInput
	default:
		return false, ErrUnknownInput
	}
}

const maxTries = 10

// AskForConfirmation asks a yes/no question until the user replies yes or no
func AskForConfirmation(text string, skip bool) bool {
	if skip {
		return true
	}

	for i := 0; i < maxTries; i++ {
		choice, err := AskForBool(text, true, skip)
		if err == nil {
			return choice
		}
		if err == ErrAbortInput {
			return false
		}
	}
	return false
}
