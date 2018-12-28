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

func removeAccountAction(c *cli.Context) {
	validateRemoveAccountAction(c)
	alias := c.Args()[0]
	removeAccount(alias)
	out.Infof("Account `%s` has been removed", alias)
}

func removeAccount(alias string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		out.Fatalf("Unable to load config, %s", err)
	}

	delete(cfg.Accounts, alias)

	// unset context if context was deleted account
	if cfg.Context == alias {
		cfg.Context = ""
	}

	err = config.SaveConfig(cfg)
	if err != nil {
		out.Fatalf("Unable to safe config, %s", err)
	}
}

func validateRemoveAccountAction(c *cli.Context) {
	if ok := util.IsValidArgsLength(c.Args(), 1); !ok {
		out.Fatal("Incorrect number of arguments for account `remove` command")
	}

	args := c.Args()
	alias := args[0]

	if !config.IsValidAlias(alias) {
		out.Fatal("Invalid alias name, please use only letters and numbers.")
	}

	_, err := config.GetAccountConfig(args[0])
	if err != nil {
		out.Fatalf("Failed to remove account `%s` from config, %s", args[0], err)
	}
}
