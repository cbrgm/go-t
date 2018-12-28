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

// Package account contains various user related commands and actions
// to be executed via command line interface
package account

import (
	"github.com/cbrgm/go-t/cmd/util"
	"github.com/cbrgm/go-t/pkg/config"
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/urfave/cli"
)

func moveAccountAction(c *cli.Context) {
	validateMoveAccountAction(c)

	aliasFrom := c.Args()[0]
	aliasTo := c.Args()[1]

	moveAccount(aliasFrom, aliasTo)
	out.Infof("Moved account `%s` to `%s`", aliasFrom, aliasTo)
}

func moveAccount(aliasFrom string, aliasTo string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		out.Fatalf("Unable to load config, %s", err)
	}

	if _, ok := cfg.Accounts[aliasFrom]; !ok {
		out.Fatalf("Cannot move account `%s` to `%s`, account `%s` does not exist", aliasFrom, aliasTo, aliasFrom)
	}

	acc := cfg.Accounts[aliasFrom]
	delete(cfg.Accounts, aliasFrom)
	cfg.Accounts[aliasTo] = acc

	// update context if moved account was set as current context
	if cfg.Context == aliasFrom {
		cfg.Context = aliasTo
		out.Infof("Switched to account `%s` as current context`\n", aliasTo)
	}

	err = config.SaveConfig(cfg)
	if err != nil {
		out.Fatalf("Unable to safe config, %s", err)
	}
}

func validateMoveAccountAction(c *cli.Context) {
	if ok := util.IsValidArgsLength(c.Args(), 2); !ok {
		out.Fatal("Incorrect number of arguments for account `remove` command")
	}

	args := c.Args()
	aliasTo := args[1]

	if !config.IsValidAlias(aliasTo) {
		out.Fatal("Invalid alias name, please use only letters and numbers")
	}
}
