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

import (
	"os"
	"reflect"
	"testing"
)

func TestSetConfigDir(t *testing.T) {
	tests := []struct {
		name      string
		configDir string
		want      string
	}{
		{
			name:      "set custom config directory",
			configDir: "~/home/go-t",
			want:      "~/home/go-t",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SetConfigDir(test.configDir)
			if test.want != customConfigDir {
				t.Errorf("GetConfigDir() = %v, want %v", customConfigDir, test.want)
			}

			customConfigDir = ""
		})
	}
}

func TestGetConfigDir(t *testing.T) {
	tests := []struct {
		name            string
		want            string
		customConfigDir string
		wantErr         bool
	}{
		{
			name:            "test get custom config dir",
			want:            "/home/test/.trc",
			customConfigDir: "/home/test/.trc",
			wantErr:         false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			customConfigDir = test.customConfigDir

			got, err := GetConfigDir()
			if (err != nil) != test.wantErr {
				t.Errorf("GetConfigDir() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("GetConfigDir() = %v, want %v", got, test.want)
			}

			customConfigDir = ""
		})
	}
}

func TestGetConfigPath(t *testing.T) {
	tests := []struct {
		name            string
		want            string
		customConfigDir string
		wantErr         bool
	}{
		{
			name:            "test get custom config path",
			want:            "/home/test/.trc/test/config.json",
			customConfigDir: "/home/test/.trc/test",
			wantErr:         false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			customConfigDir = test.customConfigDir
			got, err := GetConfigPath()
			if (err != nil) != test.wantErr {
				t.Errorf("GetConfigPath() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("GetConfigPath() = %v, want %v", got, test.want)
			}

			customConfigDir = ""
		})
	}
}

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "test new config",
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
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := NewConfig(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewConfig() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSaveAndLoadConfig(t *testing.T) {
	tests := []struct {
		name            string
		customConfigDir string
		want            *Config
		wantErr         bool
	}{
		{
			name:            "test save and load config",
			customConfigDir: "~/.trc/test",
			want:            NewConfig(),
			wantErr:         false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setupTestConfig(t, test.customConfigDir)

			err := SaveConfig(test.want)
			if (err != nil) != test.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			got, err := LoadConfig()
			if (err != nil) != test.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, test.want)
			}

			cleanupTestConfig(t, test.customConfigDir)
		})
	}
}

func TestCreateConfigDir(t *testing.T) {
	tests := []struct {
		name            string
		customConfigDir string
		wantErr         bool
	}{
		{
			name:            "test create config dir",
			customConfigDir: "~/.trc/test",
			wantErr:         false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			customConfigDir = test.customConfigDir

			if err := createConfigDir(); (err != nil) != test.wantErr {
				t.Errorf("createConfigDir() error = %v, wantErr %v", err, test.wantErr)
			}

			cleanupTestConfig(t, test.customConfigDir)
		})
	}
}

func TestIsConfigExists(t *testing.T) {
	tests := []struct {
		name            string
		customConfigDir string
		want            bool
	}{
		{
			name:            "test config dir exists",
			customConfigDir: "~/.trc/test",
			want:            true,
		},
		{
			name:            "test config dir not exists",
			customConfigDir: "~/.trc/failed",
			want:            false,
		},
	}

	// Make sure config dir exists
	setupTestConfig(t, "~/.trc/test")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			customConfigDir = test.customConfigDir
			if got := IsConfigExists(); got != test.want {
				t.Errorf("IsConfigExists() = %v, want %v", got, test.want)
			}
		})
	}

	cleanupTestConfig(t, "~/.trc/test")
}

func TestIsValidAlias(t *testing.T) {
	tests := []struct {
		name  string
		alias string
		want  bool
	}{
		{
			name:  "test alias only letters",
			alias: "test",
			want:  true,
		},
		{
			name:  "test alias with numbers at the end",
			alias: "test93",
			want:  true,
		},
		{
			name:  "test alias with numbers at the beginning",
			alias: "93test",
			want:  true,
		},
		{
			name:  "test alias with special chars at the end",
			alias: "test-93_",
			want:  true,
		},
		{
			name:  "test alias with special chars at the beginning",
			alias: "_test-93",
			want:  true,
		},
		{
			name:  "test alias with special chars",
			alias: "test%",
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsValidAlias(test.alias); got != test.want {
				t.Errorf("IsValidAlias() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestGetAccountConfig(t *testing.T) {
	tests := []struct {
		name    string
		alias   string
		want    *AccountConfig
		wantErr bool
	}{
		{
			name:  "test get account config exists",
			alias: "tester",
			want: &AccountConfig{
				ConsumerKey:    "123",
				ConsumerSecret: "123",
				AccessToken:    "123",
				AccessSecret:   "123",
			},
			wantErr: false,
		},
		{
			name:    "test get account config not exists",
			alias:   "unknown",
			wantErr: true,
		},
	}

	// Make sure config dir exists
	setupTestConfig(t, "~/.trc/test")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetAccountConfig(test.alias)
			if (err != nil) != test.wantErr {
				t.Errorf("GetAccountConfig() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("GetAccountConfig() = %v, want %v", got, test.want)
			}
		})
	}

	cleanupTestConfig(t, "~/.trc/test")
}

func TestConfig_GetAccountConfig(t *testing.T) {
	tests := []struct {
		name    string
		alias   string
		want    *AccountConfig
		wantErr bool
	}{
		{
			name:  "test get account config exists",
			alias: "tester",
			want: &AccountConfig{
				ConsumerKey:    "123",
				ConsumerSecret: "123",
				AccessToken:    "123",
				AccessSecret:   "123",
			},
			wantErr: false,
		},
		{
			name:    "test get account config not exists",
			alias:   "unknown",
			wantErr: true,
		},
	}

	// Make sure config dir exists
	setupTestConfig(t, "~/.trc/test")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, err := LoadConfig()
			if err != nil {
				t.Errorf("Failed to load test config, %s", err)
			}

			got, err := c.GetAccountConfig(test.alias)
			if (err != nil) != test.wantErr {
				t.Errorf("Config.GetAccountConfig() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Config.GetAccountConfig() = %v, want %v", got, test.want)
			}
		})
	}

	cleanupTestConfig(t, "~/.trc/test")
}

func TestConfig_GetAccountByContext(t *testing.T) {
	tests := []struct {
		name    string
		context string
		want    *AccountConfig
		wantErr bool
	}{
		{
			name:    "test get account config by context exists",
			context: "tester",
			want: &AccountConfig{
				ConsumerKey:    "123",
				ConsumerSecret: "123",
				AccessToken:    "123",
				AccessSecret:   "123",
			},
			wantErr: false,
		},
		{
			name:    "test get account config by context not exists",
			context: "unknown",
			wantErr: true,
		},
	}

	// Make sure config dir exists
	setupTestConfig(t, "~/.trc/test")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, err := LoadConfig()
			if err != nil {
				t.Errorf("Failed to load test config, %s", err)
			}

			// Set the context
			c.Context = test.context

			got, err := c.GetAccountByContext()
			if (err != nil) != test.wantErr {
				t.Errorf("Config.GetAccountByContext() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Config.GetAccountByContext() = %v, want %v", got, test.want)
			}
		})
	}

	cleanupTestConfig(t, "~/.trc/test")
}

func setupTestConfig(t *testing.T, path string) {

	customConfigDir = path

	c := NewConfig()

	c.setAccount("tester", AccountConfig{
		ConsumerKey:    "123",
		ConsumerSecret: "123",
		AccessToken:    "123",
		AccessSecret:   "123",
	})

	err := SaveConfig(c)
	if err != nil {
		t.Errorf("failed to set up test config directory, %s", err)
	}
}

func cleanupTestConfig(t *testing.T, path string) {
	customConfigDir = ""
	err := os.RemoveAll(path)
	if err != nil {
		t.Errorf("failed to clean up test config directory, %s", err)
	}
}
