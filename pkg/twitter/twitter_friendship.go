/*
 * Copyright 2018 Christian Bargmann <chris@cbrgm.net>
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

package twitter

import (
	"github.com/cbrgm/go-t/pkg/console/out"
	"github.com/cbrgm/go-t/pkg/twitter/printer"
	"github.com/dghubble/go-twitter/twitter"
)

// CreateFriendshipOptions represents user follow options
type CreateFriendshipOptions struct {
	Output     string
	SkipYes    bool
	File       bool
	UserID     int64
	ScreenName string
	Follow     bool
	Verbose    bool
}

// CreateFriendship follows a user
func (c *Client) CreateFriendship(opts *CreateFriendshipOptions) {
	params := twitter.FriendshipCreateParams{
		UserID:     opts.UserID,
		ScreenName: opts.ScreenName,
		Follow:     &opts.Follow,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
	}

	user, _, err := c.api.Friendships.Create(&params)
	if err != nil {
		out.Fatalln(err)
	}

	out.Infof("✔ successfully followed %s!\n", user.Name)
	if opts.Verbose {
		printer.PrintUser(*user, printOpts)
	}
}

// IncomingFriendshipOptions represents incoming friendship options
type IncomingFriendshipOptions struct {
	Output  string
	AsList  bool
	SkipYes bool
	File    bool
}

// IncomingFriendship shows incoming friendships
func (c *Client) IncomingFriendship(opts *IncomingFriendshipOptions) {
	/*params := twitter.FriendshipPendingParams{
		Cursor: 20,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	ids, _, err := c.api.Friendships.Incoming(&params)
	if err != nil {
		console.Fatalln(err)
	}

	printer.PrintRelationships(relationships, printOpts)*/
}

// OutgoingFriendshipOptions represents outgoing friendship options
type OutgoingFriendshipOptions struct {
	Output  string
	AsList  bool
	SkipYes bool
	File    bool
}

// OutgoingFriendships shows outgoing friendship relations
func (c *Client) OutgoingFriendships(opts *OutgoingFriendshipOptions) {
	/*params := twitter.FriendshipPendingParams{
		Cursor: 20,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	relationships, _, err := c.api.Friendships.Outgoing(&params)
	if err != nil {
		console.Fatalln(err)
	}

	printer.PrintRelationships(relationships, printOpts)*/
}

// RemoveFriendshipOptions represents remove friendship options
type RemoveFriendshipOptions struct {
	Output     string
	AsList     bool
	SkipYes    bool
	File       bool
	UserID     int64
	ScreenName string
	Verbose    bool
}

// RemoveFriendship unfollows a user
func (c *Client) RemoveFriendship(opts *RemoveFriendshipOptions) {
	params := twitter.FriendshipDestroyParams{
		ScreenName: opts.ScreenName,
		UserID:     opts.UserID,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
		AsList: opts.AsList,
	}

	user, _, err := c.api.Friendships.Destroy(&params)
	if err != nil {
		out.Fatalln(err)
	}

	out.Infof("✔ successfully unfollowed %s!\n", user.Name)
	if opts.Verbose {
		printer.PrintUser(*user, printOpts)
	}
}

// ShowFriendshipOptions represents show friendship options
type ShowFriendshipOptions struct {
	Output           string
	SourceID         int64
	TargetID         int64
	SourceScreenName string
	SourceTargetName string
}

// ShowFriendship shows friendship information
func (c *Client) ShowFriendship(opts *ShowFriendshipOptions) {
	params := twitter.FriendshipShowParams{
		SourceID:         opts.SourceID,
		TargetID:         opts.TargetID,
		SourceScreenName: opts.SourceScreenName,
		TargetScreenName: opts.SourceTargetName,
	}

	printOpts := &printer.PrintOptions{
		Output: opts.Output,
	}

	relationship, _, err := c.api.Friendships.Show(&params)
	if err != nil {
		out.Fatalln(err)
	}

	printer.PrintRelationship(*relationship, printOpts)
}
