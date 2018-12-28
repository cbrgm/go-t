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

// Package user contains various commands and actions
// to be executed via command line interface
package user

import "github.com/urfave/cli"

// Commands represents the account commands.
func Commands() cli.Command {
	return cli.Command{
		Name:  "users",
		Usage: "user related commands",
		Subcommands: []cli.Command{
			{
				Name:   "whois",
				Usage:  "retrieve detailed user information",
				Action: whoisUserAction,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "file, f",
						Usage: "get input from file or stdin",
					},
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display user information as a list",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
					cli.IntFlag{
						Name:  "tail",
						Usage: "show only the last n tweets of your timeline",
					},
					cli.IntFlag{
						Name:  "head",
						Usage: "show only the first n tweets of your timeline",
					},
					cli.IntFlag{
						Name:  "count,c",
						Usage: "show only the n tweets of your timeline",
					},
				},
			},
			{
				Name:   "search",
				Usage:  "search for users",
				Action: searchUserAction,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "file, f",
						Usage: "get input from file or stdin",
					},
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display user information as a list",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
					cli.IntFlag{
						Name:  "tail",
						Usage: "show only the last n tweets of your timeline",
					},
					cli.IntFlag{
						Name:  "head",
						Usage: "show only the first n tweets of your timeline",
					},
					cli.IntFlag{
						Name:  "count,c",
						Usage: "show only the n tweets of your timeline",
					},
				},
			},
		},
	}
}
