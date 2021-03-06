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

package timeline

import (
	"github.com/cbrgm/go-t/cmd/global"
	"github.com/cbrgm/go-t/cmd/util"
	"github.com/cbrgm/go-t/pkg/config"
	"github.com/cbrgm/go-t/pkg/console/in"
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/cbrgm/go-t/pkg/twitter"
	"github.com/urfave/cli"
)

func userTimelineAction(c *cli.Context) {
	if !util.IsValidArgsLength(c.Args(), 1) {
		out.Fatalln("invalid number of arguments passed for `t timeline user`")
	}

	userID, name, err := in.ParseUserFromArgs(c.Args()[0], c.Bool("file"))
	if err != nil {
		out.Fatal(err)
	}

	opts := &twitter.UserTimelineOptions{
		ScreenName:      name,
		UserID:          userID,
		Count:           c.Int("count"),
		Follow:          c.Bool("follow"),
		Head:            c.Int("head"),
		Tail:            c.Int("tail"),
		AsList:          c.Bool("list"),
		TrimUser:        c.Bool("trim-user"),
		Output:          c.String("output"),
		Sort:            c.String("sort"),
		IncludeRetweets: c.Bool("include-retweets"),
		ExcludeReplies:  c.Bool("exclude-retweets"),
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
	client.UserTimeline(opts)
}
