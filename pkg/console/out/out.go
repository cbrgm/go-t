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
	"fmt"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"os"
	"path/filepath"
)

const (
	errorPrefix = "error:"
	debugPrefix = "debug:"
)

var (

	// DebugMessages turns on/off console debug printing.
	DebugMessages = false

	// stderrColoredOutput represents an instance of Writer which handle escape sequence for stderr.
	stderrColoredOutput = colorable.NewColorableStderr()

	// Colorize prints message in a colorized form, dictated by the corresponding tag argument.
	Colorize = func(tag string, data interface{}) string {
		if isatty.IsTerminal(os.Stdout.Fd()) {
			colorized, ok := Theme[tag]
			if ok {
				return colorized.SprintFunc()(data)
			} // else: No theme found. Return as string.
		}
		return fmt.Sprint(data)
	}

	// Eraseline Print in new line and adjust to top so that we don't print over the ongoing progress bar.
	Eraseline = func() {
		consolePrintf("Print", Theme["Print"], "%c[2K\n", 27)
		consolePrintf("Print", Theme["Print"], "%c[A", 27)
	}
)

// wrap around standard fmt functions.
// consolePrint prints a message prefixed with message type and program name.
func consolePrint(tag string, c *color.Color, a ...interface{}) {
	switch tag {
	case "Debug":
		// if no arguments are given do not invoke debug printer.
		if len(a) == 0 {
			return
		}
		output := color.Output
		color.Output = stderrColoredOutput
		if isatty.IsTerminal(os.Stderr.Fd()) {
			c.Print(ProgramName() + ": " + debugPrefix + " ")
			c.Print(a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": "+debugPrefix+" ")
			fmt.Fprint(color.Output, a...)
		}
		color.Output = output
	case "Fatal":
		fallthrough
	case "Error":
		// if no arguments are given do not invoke fatal and error printer.
		if len(a) == 0 {
			return
		}
		output := color.Output
		color.Output = stderrColoredOutput
		if isatty.IsTerminal(os.Stderr.Fd()) {
			c.Print(ProgramName() + ": " + errorPrefix + " ")
			c.Print(a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": "+errorPrefix+" ")
			fmt.Fprint(color.Output, a...)
		}
		color.Output = output
	case "Info":
		// if no arguments are given do not invoke info printer.
		if len(a) == 0 {
			return
		}
		if isatty.IsTerminal(os.Stdout.Fd()) {
			c.Print(ProgramName() + ": ")
			c.Print(a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": ")
			fmt.Fprint(color.Output, a...)
		}
	default:
		if isatty.IsTerminal(os.Stdout.Fd()) {
			c.Print(a...)
		} else {
			fmt.Fprint(color.Output, a...)
		}
	}
}

// consolePrintf same as print with a new line.
func consolePrintf(tag string, c *color.Color, format string, a ...interface{}) {

	switch tag {
	case "Debug":
		// if no arguments are given do not invoke debug printer.
		if len(a) == 0 {
			return
		}
		output := color.Output
		color.Output = stderrColoredOutput
		if isatty.IsTerminal(os.Stderr.Fd()) {
			c.Print(ProgramName() + ": " + debugPrefix + " ")
			c.Printf(format, a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": "+debugPrefix+" ")
			fmt.Fprintf(color.Output, format, a...)
		}
		color.Output = output
	case "Fatal":
		fallthrough
	case "Error":
		// if no arguments are given do not invoke fatal and error printer.
		if len(a) == 0 {
			return
		}
		output := color.Output
		color.Output = stderrColoredOutput
		if isatty.IsTerminal(os.Stderr.Fd()) {
			c.Print(ProgramName() + ": " + errorPrefix + " ")
			c.Printf(format, a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": "+errorPrefix+" ")
			fmt.Fprintf(color.Output, format, a...)
		}
		color.Output = output
	case "Info":
		// if no arguments are given do not invoke info printer.
		if len(a) == 0 {
			return
		}
		if isatty.IsTerminal(os.Stdout.Fd()) {
			c.Print(ProgramName() + ": ")
			c.Printf(format, a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": ")
			fmt.Fprintf(color.Output, format, a...)
		}
	default:
		if isatty.IsTerminal(os.Stdout.Fd()) {
			c.Printf(format, a...)
		} else {
			fmt.Fprintf(color.Output, format, a...)
		}
	}
}

// consolePrintln - same as print with a new line.
func consolePrintln(tag string, c *color.Color, a ...interface{}) {
	switch tag {
	case "Debug":
		// if no arguments are given do not invoke debug printer.
		if len(a) == 0 {
			return
		}
		output := color.Output
		color.Output = stderrColoredOutput
		if isatty.IsTerminal(os.Stderr.Fd()) {
			c.Print(ProgramName() + ": " + debugPrefix + " ")
			c.Println(a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": "+debugPrefix+" ")
			fmt.Fprintln(color.Output, a...)
		}
		color.Output = output
	case "Fatal":
		fallthrough
	case "Error":
		// if no arguments are given do not invoke fatal and error printer.
		if len(a) == 0 {
			return
		}
		output := color.Output
		color.Output = stderrColoredOutput
		if isatty.IsTerminal(os.Stderr.Fd()) {
			c.Print(ProgramName() + ": " + errorPrefix + " ")
			c.Println(a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": "+errorPrefix+" ")
			fmt.Fprintln(color.Output, a...)
		}
		color.Output = output
	case "Info":
		// if no arguments are given do not invoke info printer.
		if len(a) == 0 {
			return
		}
		if isatty.IsTerminal(os.Stdout.Fd()) {
			c.Print(ProgramName() + ": ")
			c.Println(a...)
		} else {
			fmt.Fprint(color.Output, ProgramName()+": ")
			fmt.Fprintln(color.Output, a...)
		}
	default:
		if isatty.IsTerminal(os.Stdout.Fd()) {
			c.Println(a...)
		} else {
			fmt.Fprintln(color.Output, a...)
		}
	}
}

// ProgramName return the name of the executable program.
func ProgramName() string {
	_, progName := filepath.Split(os.Args[0])
	return progName
}
