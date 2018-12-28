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

import "errors"

var (
	// ErrAliasNotFound is returned when the account alias does not exist
	ErrAliasNotFound = errors.New("alias does not exist")
	// ErrUnableLoadConfig is returned when the config cannot be loaded from config dir
	ErrUnableLoadConfig = errors.New("unable to load config from config directory")
	// ErrUnableSafeConfig is returned when the config cannot be safed to config dir
	ErrUnableSafeConfig = errors.New("unable to safe config from config directory")
	// ErrUnableLocateHomeDir is returned when the home dir cannot be located
	ErrUnableLocateHomeDir = errors.New("unable to locate home directory")
	// ErrInvalidArgument is returned when invalid arguments have been passed to function
	ErrInvalidArgument = errors.New("argument is invalid or was nil")
)
