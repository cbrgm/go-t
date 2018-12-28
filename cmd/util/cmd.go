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

package util

import (
	"github.com/cbrgm/go-t/cmd/global"
	"github.com/cbrgm/go-t/pkg/config"
)

// GetAccountConfig returns an AccountConfig either from context defined in config file or from globalFlags
func GetAccountConfig(cfg *config.Config, g *global.Flags) (*config.AccountConfig, error) {
	var accCfg *config.AccountConfig
	var err error

	if g.IsUserSet() {
		accCfg, err = cfg.GetAccountConfig(g.User)
		if err != nil {
			return nil, err
		}
	} else {
		accCfg, err = cfg.GetAccountByContext()
		if err != nil {
			return nil, err
		}
	}
	return accCfg, nil
}

// IsValidArgsLength checks if arguments passed via command line matches the given count
func IsValidArgsLength(args []string, n int) bool {
	if args == nil && n == 0 {
		return true
	}
	if args == nil {
		return false
	}

	if n < 0 {
		return false
	}

	argsNr := len(args)
	if argsNr < n || argsNr > n {
		return false
	}
	return true
}
