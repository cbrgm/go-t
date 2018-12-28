/*
 * Copyright 2018. Christian Bargmann <chris@cbrgm.net>
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

package cmd

import (
	"github.com/cbrgm/go-t/cmd/account"
	"github.com/cbrgm/go-t/cmd/favorite"
	"github.com/cbrgm/go-t/cmd/follower"
	"github.com/cbrgm/go-t/cmd/friendship"
	"github.com/cbrgm/go-t/cmd/global"
	"github.com/cbrgm/go-t/cmd/initialize"
	"github.com/cbrgm/go-t/cmd/search"
	"github.com/cbrgm/go-t/cmd/status"
	"github.com/cbrgm/go-t/cmd/timeline"
	"github.com/cbrgm/go-t/cmd/user"
	"github.com/urfave/cli"
)

// Commands returns all go-t commands a user can execute.
func Commands() []cli.Command {
	return cli.Commands{
		// accounts commands package
		account.Commands(),
		// initialize commands package
		initialize.Commands(),
		// timeline commands package
		timeline.Commands(),
		// status commands package
		status.Commands(),
		// user commands package
		user.Commands(),
		// trend commands package
		// trend.Commands(), - TODO: not yet implemented
		search.Commands(),

		// friendship commands package
		friendship.Commands(),
		// favorite commands package
		favorite.Commands(),
		// followers commands package
		follower.Commands(),
	}
}

// GlobalFlags returns all go-t global flags.
func GlobalFlags() []cli.Flag {
	return global.GetFlags()
}
