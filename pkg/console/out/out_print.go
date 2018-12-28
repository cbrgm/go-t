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

var (

	// Print prints a message.
	Print = func(data ...interface{}) {
		consolePrint("Print", Theme["Print"], data...)
		return
	}

	// PrintC prints a message with color.
	PrintC = func(data ...interface{}) {
		consolePrint("PrintC", Theme["PrintC"], data...)
		return
	}

	// Printf prints a formatted message.
	Printf = func(format string, data ...interface{}) {
		consolePrintf("Print", Theme["Print"], format, data...)
		return
	}

	// Println prints a message with a newline.
	Println = func(data ...interface{}) {
		consolePrintln("Print", Theme["Print"], data...)
		return
	}
)
