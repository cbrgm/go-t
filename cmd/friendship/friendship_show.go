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

import (
	"github.com/cbrgm/go-t/cmd/global"
	"github.com/cbrgm/go-t/cmd/util"
	"github.com/cbrgm/go-t/pkg/config"
	"github.com/cbrgm/go-t/pkg/console/in"
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/cbrgm/go-t/pkg/twitter"
	"github.com/urfave/cli"
)

func showFriendshipAction(c *cli.Context) {
	if !util.IsValidArgsLength(c.Args(), 2) {
		out.Fatalln("invalid number of arguments passed for `t friendship show`")
	}

	sourceUserID, sourceUsername, err := in.ParseUserFromArgs(c.Args()[0], false)
	targetUserID, targetUsername, err := in.ParseUserFromArgs(c.Args()[1], c.Bool("file"))

	if err != nil {
		out.Fatal(err)
	}

	opts := &twitter.ShowFriendshipOptions{
		SourceID:         sourceUserID,
		SourceScreenName: sourceUsername,
		TargetID:         targetUserID,
		SourceTargetName: targetUsername,
		Output:           c.String("output"),
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
	client.ShowFriendship(opts)
}
