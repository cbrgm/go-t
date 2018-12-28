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
	"github.com/urfave/cli"
)

// Commands represents the account commands.
func Commands() cli.Command {
	return cli.Command{
		Name:  "accounts",
		Usage: "manage your twitter accounts",
		Subcommands: []cli.Command{
			{
				Name:   "add",
				Usage:  "add a new account",
				Action: addAccountAction,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "skip-set-context",
						Usage: "skip set account as current context",
					},
					cli.StringFlag{
						Name:  "consumer-key",
						Usage: "twitter accounts consumer key",
					},
					cli.StringFlag{
						Name:  "consumer-secret",
						Usage: "twitter accounts consumer secret",
					},
					cli.StringFlag{
						Name:  "access-key",
						Usage: "twitter accounts access key",
					},
					cli.StringFlag{
						Name:  "access-secret",
						Usage: "twitter accounts access secret",
					},
				},
			},
			{
				Name: "remove",
				Aliases: []string{
					"rm",
				},
				Usage:  "remove an existing account",
				Action: removeAccountAction,
			},
			{
				Name: "list",
				Aliases: []string{
					"ls",
				},
				Usage:  "list account details",
				Action: listAccountAction,
			},
			{
				Name:   "set",
				Usage:  "set account as current context",
				Action: setAccountAction,
			},
			{
				Name: "move",
				Aliases: []string{
					"mv",
				},
				Usage:  "update account alias",
				Action: moveAccountAction,
			},
		},
	}
}
