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

package config

import (
	"reflect"
	"testing"
)

func TestNewTypesConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "test new config not nil",
			want: &Config{
				Version:  globalConfigVersion,
				Accounts: make(map[string]AccountConfig),
				Context:  "",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := newConfig(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("newConfig() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestLoadDefaults(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "test load default config",
			want: &Config{
				Version: globalConfigVersion,
				Context: "",
				Accounts: map[string]AccountConfig{
					"example": AccountConfig{
						ConsumerKey:    "",
						ConsumerSecret: "",
						AccessToken:    "",
						AccessSecret:   "",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := newConfig()

			if c.loadDefaults(); !reflect.DeepEqual(c, test.want) {
				t.Errorf("loadDefaults() = %v, want %v", c, test.want)
			}
		})
	}
}

func TestSetAccount(t *testing.T) {
	tests := []struct {
		name  string
		alias string
		want  *Config
	}{
		{
			name:  "test insert new account alias if not exists",
			alias: "test",
			want: &Config{
				Version: globalConfigVersion,
				Context: "",
				Accounts: map[string]AccountConfig{
					"example": {
						ConsumerKey:    "",
						ConsumerSecret: "",
						AccessToken:    "",
						AccessSecret:   "",
					},
					"test": {
						ConsumerKey:    "",
						ConsumerSecret: "",
						AccessToken:    "",
						AccessSecret:   "",
					},
				},
			},
		},
		{
			name:  "test do not insert new account alias if exists",
			alias: "test",
			want: &Config{
				Version: globalConfigVersion,
				Context: "",
				Accounts: map[string]AccountConfig{
					"example": {
						ConsumerKey:    "",
						ConsumerSecret: "",
						AccessToken:    "",
						AccessSecret:   "",
					},
					"test": {
						ConsumerKey:    "",
						ConsumerSecret: "",
						AccessToken:    "",
						AccessSecret:   "",
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := newConfig()
			c.loadDefaults()

			c.setAccount(test.alias, AccountConfig{
				ConsumerKey:    "",
				ConsumerSecret: "",
				AccessToken:    "",
				AccessSecret:   "",
			})

			c.setAccount(test.alias, AccountConfig{
				ConsumerKey:    "updated",
				ConsumerSecret: "updated",
				AccessToken:    "updated",
				AccessSecret:   "updated",
			})

			if c.loadDefaults(); !reflect.DeepEqual(c, test.want) {
				t.Errorf("setAccount() = %v, want %v", c, test.want)
			}
		})
	}
}
