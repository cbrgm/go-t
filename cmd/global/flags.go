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

package global

import "github.com/urfave/cli"

// Flags represents global cli flags passed to go-t
type Flags struct {
	User   string
	Config string
	Debug  bool
}

// GetFlags returns all global flags
func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "set account as current context",
			EnvVar: "TWITTER_USER",
		},
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "set custom config location",
			EnvVar: "TWITTER_CONFIG",
		},
		cli.StringFlag{
			Name:  "debug",
			Usage: "show debug messages",
		},
	}
}

// ParseFlags parses a cli context and returns a Flags struct
func ParseFlags(c *cli.Context) *Flags {
	return &Flags{
		User:   c.GlobalString("user"),
		Config: c.GlobalString("config"),
		Debug:  c.GlobalBool("debug"),
	}
}

// IsUserSet indicates if the global user flag is set
func (g *Flags) IsUserSet() bool {
	return g.User != ""
}

// IsConfigSet indicates if the global config flag is set
func (g *Flags) IsConfigSet() bool {
	return g.Config != ""
}
