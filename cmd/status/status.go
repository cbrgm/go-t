/*
 * Copyright 2018 Christian Bargmann <chris@cbrgm.net>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package status contains various commands and actions
// to be executed via command line interface
package status

import (
	"github.com/urfave/cli"
)

// Commands represents the status commands.
func Commands() cli.Command {
	return cli.Command{
		Name:  "status",
		Usage: "twitter status related commands",
		Subcommands: []cli.Command{
			{
				Name:   "update",
				Usage:  "update your twitter status",
				Action: updateStatusAction,
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
				Name:   "lookup",
				Usage:  "lookup a twitter status",
				Action: lookupStatusAction,
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
				},
			},
			{
				Name:   "remove",
				Action: removeStatusAction,
				Usage:  "remove your twitter status",
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
				Name:   "retweet",
				Action: retweetStatusAction,
				Usage:  "retweet a twitter status",
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
				Name:   "retweets",
				Action: retweetsStatusAction,
				Usage:  "show retweets of a twitter status",
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
			{
				Name:   "unretweet",
				Action: unretweetStatusAction,
				Usage:  "unretweet a twitter status",
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
		},
	}
}
