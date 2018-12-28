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

// Package trend contains various commands and actions
// to be executed via command line interface
package trend

import "github.com/urfave/cli"

// Commands represents the trend commands.
func Commands() cli.Command {
	return cli.Command{
		Name:  "trends",
		Usage: "trend related commands",
		Subcommands: []cli.Command{
			{
				Name:   "available",
				Usage:  "retrieve detailed user information",
				Action: availableTrendAction,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display user information as a list",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
				},
			},
		},
	}
}
