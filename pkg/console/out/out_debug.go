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

	// Debug prints a debug message without a new line
	// Debug prints a debug message.
	Debug = func(data ...interface{}) {
		if DebugMessages {
			consolePrint("Debug", Theme["Debug"], data...)
		}
	}

	// Debugf prints a debug message with a new line.
	Debugf = func(format string, data ...interface{}) {
		if DebugMessages {
			consolePrintf("Debug", Theme["Debug"], format, data...)
		}
	}

	// Debugln prints a debug message with a new line.
	Debugln = func(data ...interface{}) {
		if DebugMessages {
			consolePrintln("Debug", Theme["Debug"], data...)
		}
	}
)
