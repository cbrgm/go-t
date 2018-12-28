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

// Package initialize contains various commands and actions
// to be executed via command line interface
package initialize

import (
	"github.com/cbrgm/go-t/pkg/config"
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/urfave/cli"
)

// Commands represents the initialize commands.
func Commands() cli.Command {
	return cli.Command{
		Name:   "init",
		Usage:  "initialize go-t configuration directory and file",
		Action: initAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "overwrite",
				Usage: "force overwriting the config file",
			},
			cli.StringFlag{
				Name:  "configDir, d",
				Usage: "custom config directory location",
			},
		},
	}
}

func initAction(c *cli.Context) error {
	var (
		override  = c.Bool("overwrite")
		configDir = c.String("configDir")
	)

	if configDir != "" {
		config.SetConfigDir(configDir)
	}

	path, err := config.GetConfigPath()
	if err != nil {
		out.Fatalf("%s", err)
	}

	if !config.IsConfigExists() || override {
		out.Infof("Initializing default config at %s\n", path)

		cfg := config.NewConfig()
		err := config.SaveConfig(cfg)
		if err != nil {
			return err
		}
		out.PrintC("Successful. In a next step please add a twitter account using `t config add`\n")
	} else {
		out.Errorf("Config already exists at %s.\nSkipping initialization. Use `--overwrite` to force overwriting \n", path)
	}
	return nil
}
