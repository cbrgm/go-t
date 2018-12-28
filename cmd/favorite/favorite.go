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

package favorite

import "github.com/urfave/cli"

// Commands represents the friendship commands.
func Commands() cli.Command {
	return cli.Command{
		Name: "favorites",
		Aliases: []string{
			"fav",
		},
		Usage: "favorite related commands",
		Subcommands: []cli.Command{
			{
				Name:   "like",
				Usage:  "create a new favorite e.g. `like` a tweet",
				Action: createFavoriteAction,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "file, f",
						Usage: "get input from file or stdin",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "skip confirmation",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display status as a list",
					},
					cli.BoolFlag{
						Name:  "verbose, v",
						Usage: "verbose output",
					},
				},
			},
			{
				Name:   "dislike",
				Action: removeFavoriteAction,
				Usage:  "remove an existing favorite, e.g. `dislike` a tweet",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "file, f",
						Usage: "get input from file or stdin",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "skip confirmation",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display status as a list",
					},
					cli.BoolFlag{
						Name:  "verbose, v",
						Usage: "verbose output",
					},
				},
			},
			{
				Name: "list",
				Aliases: []string{
					"ls",
				},
				Action: listFavoriteAction,
				Usage:  "list all favorites of a user",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "file, f",
						Usage: "get input from file or stdin",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "skip confirmation",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display status as a list",
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
					cli.StringFlag{
						Name:  "sort, s",
						Usage: "change sorting of the results returned",
					},
				},
			},
		},
	}
}
