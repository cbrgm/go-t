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

package util

import "testing"

func TestIsValidArgsLength(t *testing.T) {

	tests := []struct {
		name   string
		want   bool
		number int
		args   []string
	}{
		{
			name:   "test with correct params",
			want:   true,
			number: 3,
			args: []string{
				"arg0",
				"arg1",
				"arg2",
			},
		},
		{
			name:   "test with incorrect number count",
			want:   false,
			number: 4,
			args: []string{
				"arg0",
				"arg1",
				"arg2",
			},
		},
		{
			name:   "test with incorrect number count",
			want:   false,
			number: 2,
			args: []string{
				"arg0",
				"arg1",
				"arg2",
			},
		},
		{
			name:   "test with negative number count",
			want:   false,
			number: -1,
			args: []string{
				"arg0",
				"arg1",
				"arg2",
			},
		},
		{
			name:   "test with 0 number count",
			want:   false,
			number: 0,
			args: []string{
				"arg0",
				"arg1",
				"arg2",
			},
		},
		{
			name:   "test with empty args",
			want:   true,
			number: 0,
			args:   []string{},
		},
		{
			name:   "test with nil",
			want:   false,
			number: 1,
			args:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsValidArgsLength(test.args, test.number); test.want != got {
				t.Errorf("%s failed, want: %t, got: %t", test.name, test.want, got)
			}
		})
	}
}
