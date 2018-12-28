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
	"os"
)

func setAccountAction(c *cli.Context) {
	validateSetAccountAction(c)
	alias := c.Args()[0]
	setAccount(alias)
	out.Infof("Switched to account `%s`, context is now active", alias)
}

func setAccount(alias string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		out.Fatalf("Unable to load config, %s", err)
	}

	if cfg.Context == alias {
		out.Infof("Account `%s` is already set as current context", alias)
		os.Exit(1)
	}

	cfg.Context = alias

	err = config.SaveConfig(cfg)
	if err != nil {
		out.Fatalf("Unable to safe config, %s", err)
	}
}

func validateSetAccountAction(c *cli.Context) {
	if ok := util.IsValidArgsLength(c.Args(), 1); !ok {
		out.Fatal("Incorrect number of arguments for account `set` command")
	}

	args := c.Args()

	_, err := config.GetAccountConfig(args[0])
	if err != nil {
		out.Fatalf("Unable to switch to account `%s`, %s", args[0], err)
	}
}
