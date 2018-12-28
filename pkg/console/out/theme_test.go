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

package out

import (
	"testing"

	"github.com/fatih/color"
)

func TestSetColor(t *testing.T) {
	type args struct {
		tag string
		cl  *color.Color
	}
	tests := []struct {
		name string
		want bool
		args args
	}{
		{
			name: "test set color",
			want: true,
			args: args{
				tag: "test",
				cl:  color.New(color.BgBlue),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SetColor(test.args.tag, test.args.cl)
			if _, ok := Theme[test.args.tag]; !ok {
				t.Errorf("Failed to set color, want %t, got %t", test.want, ok)
			}
		})
	}
}
