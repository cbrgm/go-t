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

package friendship

import "github.com/urfave/cli"

// Commands represents the friendship commands.
func Commands() cli.Command {
	return cli.Command{
		Name: "friendships",
		Aliases: []string{
			"fs",
		},
		Usage: "friendship related commands",
		Subcommands: []cli.Command{
			{
				Name:   "follow",
				Usage:  "create a new friendship e.g. follow someone",
				Action: createFriendshipAction,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
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
						Name:  "verbose, v",
						Usage: "verbose output",
					},
				},
			},
			{
				Name:   "unfollow",
				Action: removeFriendshipAction,
				Usage:  "retweet a twitter status",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
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
						Name:  "verbose, v",
						Usage: "verbose output",
					},
				},
			},
			{
				Name:   "show",
				Action: showFriendshipAction,
				Usage:  "show friendship relationship between users",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
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
		},
	}
}
