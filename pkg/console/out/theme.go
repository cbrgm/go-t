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

import "github.com/fatih/color"

var (
	// Theme contains default color mapping.
	Theme = map[string]*color.Color{
		"Debug":  color.New(color.FgWhite, color.Faint, color.Italic),
		"Fatal":  color.New(color.FgRed, color.Italic, color.Bold),
		"Error":  color.New(color.FgYellow, color.Italic),
		"Info":   color.New(color.FgGreen, color.Bold),
		"Print":  color.New(),
		"PrintC": color.New(color.FgGreen, color.Bold),
	}
)

// SetColorOff disables coloring for the entire session.
func SetColorOff() {
	color.NoColor = true
}

// SetColorOn enables coloring for the entire session.
func SetColorOn() {
	color.NoColor = false
}

// SetColor sets a color for a particular tag.
// Can be used to set your own individual theme.
func SetColor(tag string, cl *color.Color) {
	Theme[tag] = cl
}
