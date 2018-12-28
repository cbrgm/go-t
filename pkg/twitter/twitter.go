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

package twitter

import (
	"github.com/cbrgm/go-t/pkg/config"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Client represents the twitter api client
type Client struct {
	api   *twitter.Client
	debug bool
}

// NewFromConfig returns a new twitter Client from accountConfig
func NewFromConfig(accCfg *config.AccountConfig, debug bool) *Client {
	config := oauth1.NewConfig(accCfg.ConsumerKey, accCfg.ConsumerSecret)
	token := oauth1.NewToken(accCfg.AccessToken, accCfg.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	api := twitter.NewClient(httpClient)

	return &Client{
		api:   api,
		debug: debug,
	}
}
