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
	"github.com/cbrgm/go-t/pkg/console/in"
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/urfave/cli"
)

func addAccountAction(c *cli.Context) {
	validateAddAccountAction(c)

	var (
		alias          = c.Args()[0]
		skipSetContext = c.Bool("skip-set-context")
		consumerKey    = c.String("consumer-key")
		consumerSecret = c.String("consumer-secret")
		accessToken    = c.String("access-token")
		accessSecret   = c.String("access-secret")
	)

	if consumerKey == "" {
		input, err := in.AskForString("Please enter your `consumer key`", "consumer key", false)
		if err != nil {
			out.Fatalln(err)
		}
		consumerKey = input
	}

	if consumerSecret == "" {
		input, err := in.AskForString("Please enter your `consumer secret`", "consumer secret", false)
		if err != nil {
			out.Fatalln(err)
		}
		consumerSecret = input
	}

	if accessToken == "" {
		input, err := in.AskForString("Please enter your `access token`:", "access token", false)
		if err != nil {
			out.Fatalln(err)
		}
		accessToken = input
	}

	if accessSecret == "" {
		input, err := in.AskForString("Please enter your `access secret`:", "access secret", false)
		if err != nil {
			out.Fatalln(err)
		}
		accessSecret = input
	}

	account := config.AccountConfig{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken:    accessToken,
		AccessSecret:   accessSecret,
	}

	addAccount(alias, account, skipSetContext)
	out.Infof("Account `%s` configured\n", alias)
}

func addAccount(alias string, account config.AccountConfig, skipSetContext bool) {
	cfg, err := config.LoadConfig()
	if err != nil {
		out.Fatalf("Unable to load config, %s", err)
	}

	cfg.Accounts[alias] = account

	// use added account as current context
	if !skipSetContext {
		out.Infof("Setting `%s` as current context\n", alias)
		cfg.Context = alias
	}

	err = config.SaveConfig(cfg)
	if err != nil {
		out.Fatalf("Unable to safe config, %s", err)
	}
}

func validateAddAccountAction(c *cli.Context) {
	if ok := util.IsValidArgsLength(c.Args(), 1); !ok {
		out.Fatal("Incorrect number of arguments for account `add` command")
	}

	args := c.Args()
	alias := args[0]

	if !config.IsValidAlias(alias) {
		out.Fatal("Invalid alias name, please use only letters and numbers")
	}
}
