/*
 * Copyright 2018. Christian Bargmann <chris@cbrgm.net>
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

package main

import (
	"fmt"
	"github.com/cbrgm/go-t/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
	"runtime"
)

var (
	// Version of go-t.
	Version string
	// Revision or Commit this binary was built from.
	Revision string
	// GoVersion running this binary.
	GoVersion = runtime.Version()
)

func main() {
	app := cli.NewApp()
	app.Name = "go-t"
	app.Author = "Christian Bargmann"
	app.Email = "chris@cbrgm.net"
	app.Usage = "A fast, simple and easy to use command-line client for Twitter written in Go."
	app.Description = "Twitter actions from your command line"
	app.Version = fmt.Sprintf("%s revision %s %s", Version, Revision, GoVersion)

	// set cli commands and global flags
	app.Commands = cmd.Commands()
	app.Flags = cmd.GlobalFlags()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
