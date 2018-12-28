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

// Package timeline contains various commands and actions
// to be executed via command line interface
package timeline

import (
	"github.com/urfave/cli"
)

// Commands represents the timeline commands.
func Commands() cli.Command {
	return cli.Command{
		Name:  "timeline",
		Usage: "timeline related commands",
		Aliases: []string{
			"tl",
		},
		Action: homeTimelineAction,
		Flags: []cli.Flag{
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
			cli.BoolFlag{
				Name:  "list, l",
				Usage: "display tweets of your timeline as a list",
			},
			cli.StringFlag{
				Name:  "output, o",
				Usage: "change output format e.g. json, yaml",
			},
			cli.BoolFlag{
				Name:  "follow",
				Usage: "Stream tweets of your timeline and watch for updates",
			},
			cli.StringFlag{
				Name:  "sort, s",
				Usage: "change sorting of the results returned",
			},
			cli.BoolFlag{
				Name:  "trim-user",
				Usage: "trim tweets by user",
			},
			cli.BoolFlag{
				Name:  "exclude-replies",
				Usage: "exclude replies",
			},
		},
		Subcommands: []cli.Command{
			{
				Name:   "user",
				Usage:  "user timeline related commands",
				Action: userTimelineAction,
				Flags: []cli.Flag{
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
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display tweets of your timeline as a list",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
					cli.BoolFlag{
						Name:  "follow",
						Usage: "Stream tweets of your timeline and watch for updates",
					},
					cli.StringFlag{
						Name:  "sort, s",
						Usage: "change sorting of the results returned",
					},
					cli.BoolFlag{
						Name:  "file, f",
						Usage: "get input from file or stdin",
					},
					cli.BoolFlag{
						Name:  "trim-user",
						Usage: "trim tweets by user",
					},
					cli.BoolFlag{
						Name:  "exclude-replies",
						Usage: "exclude replies",
					},
				},
			},
			{
				Name:   "mentions",
				Usage:  "mentions timeline related commands",
				Action: mentionsTimelineAction,
				Flags: []cli.Flag{
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
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display tweets of your timeline as a list",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
					cli.StringFlag{
						Name:  "sort, s",
						Usage: "change sorting of the results returned",
					},
					cli.BoolFlag{
						Name:  "follow",
						Usage: "Stream tweets of your timeline and watch for updates",
					},
					cli.BoolFlag{
						Name:  "trim-user",
						Usage: "trim tweets by user",
					},
				},
			},
			{
				Name:   "retweets",
				Usage:  "retweets of your tweets timeline related commands",
				Action: retweetsTimelineAction,
				Flags: []cli.Flag{
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
					cli.BoolFlag{
						Name:  "list, l",
						Usage: "display tweets of your timeline as a list",
					},
					cli.StringFlag{
						Name:  "output, o",
						Usage: "change output format e.g. json, yaml",
					},
					cli.StringFlag{
						Name:  "sort, s",
						Usage: "change sorting of the results returned",
					},
					cli.BoolFlag{
						Name:  "follow",
						Usage: "Stream tweets of your timeline and watch for updates",
					},
					cli.BoolFlag{
						Name:  "trim-user",
						Usage: "trim tweets by user",
					},
				},
			},
		},
	}
}
