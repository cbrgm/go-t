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
	"fmt"
	"github.com/cbrgm/go-t/cmd/util"
	"github.com/cbrgm/go-t/pkg/config"
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/urfave/cli"
)

func listAccountAction(c *cli.Context) {
	var accounts = map[string]config.AccountConfig{}

	if ok := util.IsValidArgsLength(c.Args(), 0); ok {
		accounts = getAccounts()
	}

	if ok := util.IsValidArgsLength(c.Args(), 1); ok {
		alias := c.Args()[0]
		accounts = getAccount(alias)
	}

	listAccounts(accounts, c.String("format"))
}

func listAccounts(accounts map[string]config.AccountConfig, format string) {
	for k := range accounts {
		s := fmt.Sprintf("%s\n", k)
		out.PrintC(s)
	}
}

func getAccounts() map[string]config.AccountConfig {
	cfg, err := config.LoadConfig()
	if err != nil {
		out.Fatalf("Unable to load config, %s", err)
	}

	acc := cfg.Accounts
	if len(acc) == 0 {
		out.Fatal("No account have been added. Please use `t config add` to add an account to your config")
	}
	return acc
}

func getAccount(alias string) map[string]config.AccountConfig {
	cfg, err := config.LoadConfig()
	if err != nil {
		out.Fatalf("Unable to load config, %s", err)
	}

	acc, ok := cfg.Accounts[alias]
	if !ok {
		out.Fatalf("Failed to show account details for `%s`, account does not exist", alias)
	}

	result := map[string]config.AccountConfig{
		alias: acc,
	}
	return result
}
