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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
)

const (
	globalConfigVersion    = "1"
	globalConfigDir        = ".trc/"
	globalConfigWindowsDir = "trc\\"
	globalConfigFile       = "config.json"
)

// customConfigDir contains the whole path to config dir. Only access via get/set functions.
var customConfigDir string

// SetConfigDir sets a custom go-t client config folder.
func SetConfigDir(configDir string) {
	customConfigDir = configDir
}

// GetConfigDir constructs go-t client config folder.
func GetConfigDir() (string, error) {
	if customConfigDir != "" {
		return customConfigDir, nil
	}
	homeDir, e := homedir.Dir()
	if e != nil {
		return "", ErrUnableLocateHomeDir
	}
	var configDir string
	// For windows the path is slightly different
	if runtime.GOOS == "windows" {
		configDir = filepath.Join(homeDir, globalConfigWindowsDir)
	} else {
		configDir = filepath.Join(homeDir, globalConfigDir)
	}
	return configDir, nil
}

// GetConfigPath onstructs go-t client configuration path.
func GetConfigPath() (string, error) {
	if customConfigDir != "" {
		return filepath.Join(customConfigDir, globalConfigFile), nil
	}
	dir, err := GetConfigDir()
	if err != nil {
		return "", ErrUnableLocateHomeDir
	}
	return filepath.Join(dir, globalConfigFile), nil
}

// NewConfig initializes a new config.
func NewConfig() *Config {
	config := newConfig()
	config.loadDefaults()
	return config
}

// SaveConfig saves the configuration file and returns error if any.
func SaveConfig(config *Config) error {
	if config == nil {
		return ErrInvalidArgument
	}

	err := createConfigDir()
	if err != nil {
		return err
	}

	// Save the config.
	file, err := GetConfigPath()
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("%s, %s", ErrUnableSafeConfig, err)
	}

	err = ioutil.WriteFile(file, b, 0700)
	if err != nil {
		return fmt.Errorf("%s, %s", ErrUnableSafeConfig, err)
	}

	return nil
}

// LoadConfig loads the config file
func LoadConfig() (*Config, error) {
	file, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s, %s", ErrUnableLoadConfig, err)
	}

	// Load config
	var config = &Config{}
	err = json.Unmarshal(b, config)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", ErrUnableLoadConfig, err)
	}

	return config, nil
}

// createConfigDir creates go-t client config folder
func createConfigDir() error {
	p, err := GetConfigDir()
	if err != nil {
		return err
	}
	if e := os.MkdirAll(p, 0700); e != nil {
		return err
	}
	return nil
}

// IsConfigExists returns false if config doesn't exist.
func IsConfigExists() bool {
	configFile, err := GetConfigPath()
	if err != nil {
		return false
	}
	if _, e := os.Stat(configFile); e != nil {
		return false
	}
	return true
}

// IsValidAlias Check if an account alias valid.
func IsValidAlias(alias string) bool {
	return regexp.MustCompile("[a-zA-Z0-9-_]+$").MatchString(alias)
}

// GetAccountConfig retrieves account specific configuration.
func GetAccountConfig(alias string) (*AccountConfig, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	// if account is exact return quickly.
	if _, ok := config.Accounts[alias]; ok {
		accountConfig := config.Accounts[alias]
		return &accountConfig, nil
	}

	// return error if cannot be matched.
	return nil, fmt.Errorf("%s, %s", ErrAliasNotFound, alias)
}

// GetAccountConfig retrieves account specific configuration.
func (c *Config) GetAccountConfig(alias string) (*AccountConfig, error) {

	// if account is exact return quickly.
	if _, ok := c.Accounts[alias]; ok {
		accountConfig := c.Accounts[alias]
		return &accountConfig, nil
	}

	// return error if cannot be matched.
	return nil, fmt.Errorf("%s, %s", ErrAliasNotFound, alias)
}

// GetAccountByContext retrieves account specific configuration.
func (c *Config) GetAccountByContext() (*AccountConfig, error) {

	// if account is exact return quickly.
	if _, ok := c.Accounts[c.Context]; ok {
		accountConfig := c.Accounts[c.Context]
		return &accountConfig, nil
	}

	// return error if cannot be matched.
	return nil, fmt.Errorf("%s, %s", ErrAliasNotFound, c.Context)
}
