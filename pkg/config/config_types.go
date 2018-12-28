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

package config

// Config represents the cli config
type Config struct {
	Version  string                   `json:"version"`
	Context  string                   `json:"context"`
	Accounts map[string]AccountConfig `json:"account"`
}

// AccountConfig represents twitter app auth credentials
type AccountConfig struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

func newConfig() *Config {
	config := new(Config)
	config.Version = globalConfigVersion
	config.Accounts = make(map[string]AccountConfig)
	config.Context = ""

	return config
}

// loadDefaults sets the configs default configuration
func (c *Config) loadDefaults() {
	c.Version = globalConfigVersion
	c.Context = ""

	c.setAccount("example", AccountConfig{
		ConsumerKey:    "",
		ConsumerSecret: "",
		AccessToken:    "",
		AccessSecret:   "",
	})
}

// setAccount sets account config if not empty.
func (c *Config) setAccount(alias string, cfg AccountConfig) {
	if _, ok := c.Accounts[alias]; !ok {
		c.Accounts[alias] = cfg
	}
}
