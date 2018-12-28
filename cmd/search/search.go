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

package search

import (
	"github.com/cbrgm/go-t/cmd/global"
	"github.com/cbrgm/go-t/cmd/util"
	"github.com/cbrgm/go-t/pkg/config"
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/cbrgm/go-t/pkg/twitter"
	"github.com/urfave/cli"
)

// Commands returns search related commands
func Commands() cli.Command {
	return cli.Command{
		Name:   "search",
		Usage:  "search for tweets",
		Action: searchTweetAction,
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
	}
}

func searchTweetAction(c *cli.Context) {
	if !util.IsValidArgsLength(c.Args(), 1) {
		out.Fatalln("invalid number of arguments passed for `t search`")
	}

	opts := &twitter.TweetSearchOptions{
		Query:  c.Args()[0],
		AsList: c.Bool("list"),
		Output: c.String("output"),
		File:   c.Bool("file"),
		Head:   c.Int("head"),
		Tail:   c.Int("tail"),
		Count:  c.Int("count"),
	}

	// parse global flags
	g := global.ParseFlags(c)

	if g.IsConfigSet() {
		config.SetConfigDir(g.Config)
	}

	// load the config
	cfg, err := config.LoadConfig()
	if err != nil {
		out.Fatal(err)
	}

	accCfg, err := util.GetAccountConfig(cfg, g)
	if err != nil {
		out.Fatal(err)
	}

	client := twitter.NewFromConfig(accCfg, g.Debug)
	client.SearchTweet(opts)
}
