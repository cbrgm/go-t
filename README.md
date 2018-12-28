# go-`t`

<img src=".img/go_t_logo_312x156.png" alt="go-t logo" title="go-t logo" />

[![GoDoc](https://godoc.org/github.com/cbrgm/go-t?status.svg)](https://godoc.org/github.com/cbrgm/go-t)
[![Build Status](https://drone.cbrgm.net/api/badges/cbrgm/go-t/status.svg)](https://drone.cbrgm.net/cbrgm/go-t)
[![Go Report Card](https://goreportcard.com/badge/github.com/cbrgm/go-t)](https://goreportcard.com/report/github.com/cbrgm/go-t)
[![](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://github.com/cbrgm/go-t/blob/master/LICENSE)
[![](https://img.shields.io/github/release/cbrgm/go-t.svg)](https://github.com/cbrgm/go-t/releases)

 **A blazing-fast, simple and easy to use command-line client for [Twitter](https://twitter.com) written in [Go](https://golang.org/).**
🚀📨

## Overview

1.  [Demo](#demo)
2.  [Features](#features-&-roadmap)
3.  [Installation](#installation)
4.  [Configuration](#configuration)
5.  [Usage Examples](#usage-examples)
6.  [Credit & License](#credit-&-license)
7.  [Contributing](#contributing)
8.  [Acknowledgements](#acknowledgements)

## Demo

<p align="center">
<img src=".img/demo.gif" alt="go-t demo" title="go-t demo" />
</p>

## Features & Roadmap

Please have a look at [docs/features.md](https://github.com/cbrgm/go-t/blob/master/docs/features.md) for an extensive list of all features along with several usage examples.

| **Feature**           | **Status**   | **Description**                                                                                   |
| --------------------- | ------------ | ------------------------------------------------------------------------------------------------- |
| Multi-account support | ✔ [beta]     | Manage multiple twitter accounts                                                                  |
| Tweets                | ✔ [beta]     | Post, (un)retweet tweets, list retweets from command line or stdin                                |
| Favorites             | ✔ [beta]     | (Dis)like tweets from the command line or stdin                                                   |
| Friendships           | ✔ [beta]     | (Un)follow users, show friendship satus from command line or stdin from the command line or stdin |
| Timelines             | ✔ [beta]     | List tweets of your personal-, mentions- and retweets- timeline                                   |
| Users                 | ✔  [beta]    | Retrieve user information, Mute users, search users from command line or stdin                    |
| Followers             | ✔  [alpha]   | List friends, groupies and leaders                                                                |
| Ouput formats         | ✔ [alpha]    | Support different output formats for commands such as JSON, YAML, lists or CSV                    |
| Sorting               | ✔ [alpha]    | Support sorting for output                                                                        |
| Trends                | 🚧 [planned] | Retrieve trend information from different countries                                               |
| Searching             | ✔  [alpha]  | Search for specifific tweet content , hashtags, links, ... keywords via regex                     |
| Direct Messages       | 🚧 [planned] | List, send and delete direct messages from command line or stdin                                  |
| List support          | 🚧 [planned] | Create/Delete from command line or stdin                                                          |
| Auto-Completion       | 🚧 [planned] | Auto-completion different shells like bash, zsh, fish, ...                                        |

-   📝 [planned] = Feature is planned, a concept is still being worked on or not yet started.
-   🚧 [alpha] = Commands / Flags / Arguments can change incompatibly with a new release.
-   ✔ [beta] = Commands / Flags / Arguments are considered to be mostly stable and backwards compatible to an earlier released version.
-   ✔ [stable] = Commands / Flags / Arguments are stable and will not change with a new release.

## Installation

In case you have [Go](https://golang.org/) 1.11+ installed:

```bash
go get github.com/cbrgm/go-t
```

You can also download precompiled binaries. See [Releases](https://github.com/cbrgm/go-t/releases).

## Configuration

Create a configuration directory with a default configuration for go-`t`.

    t init

The configuration file can be found at `~/.trc/config.json`.

[go-`t`](https://github.com/cbrgm/go-t) communicates with the Twitter API v1.1 to retrieve information. You will need to create an account on the Twitter developer website and credentials to use the Twitter API. Twitters API requires OAuth for all of its functionality, so you'll need a registered Twitter application.

Creating a developer account is quickly done.

1.  Go to <https://dev.twitter.com/user/login> and log in with your Twitter username and password. If you don't have a Twitter account yet, click on the Sign up link below the Username field.

2.  Go to the Twitter application page at <https://dev.twitter.com/apps> and click `Create a new application`. Follow the instructions on the screen. Enter an application name, description and website.

3.  Click the `Key and Access Tokens` tab to collect the credentials for your Twitter developer account. Click `Create my access token` at the bottom of the page. The following values are needed to add your Twitter account to go-`t`:

-   `Consumer Key`
-   `Consumer Secret`
-   `Access Key`
-   `Access Secret`

With the following command you add an account with the alias `foo` to be used by go-`t` (Important: `foo` is **not** your Twitter account name, but just an alias used by go-`t` to identify your credentials).

Follow the setup instructions and add your account:

    t accounts add foo

## Usage Examples

Here you will find some examples as an introduction how to use go-`t`. You can find more examples at [docs/examples](github.com/chttps://github.com/cbrgm/go-t/blob/master/docs/examples) .

Do you have any other examples of using go-`t`? Please share them with others! See the [Contributing Guide](https://github.com/cbrgm/go-t/blob/master/CONTRIBUTING.md).

**Send a tweet**

    t status update "First tweet with go-`t`! Whoop!"

**Send a tweet from stdin**

    echo "First tweet with go-`t` from stdin! Whoop!" | t status update -f -

**Send a tweet and like it after beeing published**

    echo "Instant like!" | t status update -v -l -y -f - | awk '{print $1}' | t fav like -y -f -

**Delete a tweet**

    t status rm <tweet id>

**Retweet a tweet**

    t status retweet <tweet id>

**Show all retweets of a tweet**

    t status retweets <tweet id>

**like a tweet**

    t favorites like <tweet id>

**dislike a tweet**

    t favorites dislike <tweet id>

**follow users**

    t friendships follow @foo

**Show all likes of user `foo` as a list**

    t favorites list @foo

**Like the last 4 tweets liked by user `foo`**

     t favorites list @foo -l -c 4 | awk '{print $1}' | xargs -I tweet t favorites like -y tweet

**Show your latest timeline**

    t timeline

**Show your latest timeline as list, sorted by favorites count, descending**

    t timeline -l --sort likes,asc

**Show your latest mentions**

    t timeline mentions

**Favorite the last 10 tweets that mention you**

    t timeline mentions -c 10 -l | awk '{print $1}' | xargs -I tweet t fav like tweet

**Show latest retweets**

    t timeline retweets

**Show the 50 latest tweets of user `foo`'s timeline as list**

    t timeline user @foo -l -c 50

**Search for term `foo` in your timeline**

    t timeline -l | grep @foo

**Get detailed user information**

    t users whois foo

**Search for users with name `golang` and output json**

    t users search "golang" -o json

**list all followers of user `foo`**

    t followers @foo

**list all followers of user `foo` as list, sorted by follower's recent tweet activity**

    t followers @foo -l --sort tweeted

**Unfollow the last 10 persons you are following, who tweeted less then all your other followers**

    t followers @foo -l --sort tweeted | awk '{print $1}' | xargs -I user t users unfollow user -y

**list all followers of user `foo` as list, sorted by follower's registration date**

    t followers @foo -l --sort tweeted

## Credit & License

go-`t` is open-source and is developed under the terms of the [Apache 2.0 License](https://github.com/cbrgm/go-t/blob/master/LICENSE).

Maintainer of this repository is:

-   [@cbrgm](https://github.com/cbrgm) | Christian Bargmann <mailto:chris@cbrgm.net>

Please refer to the git commit log for a complete list of contributors.

## Contributing

See the [Contributing Guide](https://github.com/cbrgm/go-t/blob/master/CONTRIBUTING.md).

## Acknowledgements

go-`t` was initially started by Christian Bargmann.

This project is inspired by [sferik/t](https://github.com/sferik/t), a command-line power tool for Twitter. go-`t` uses the [dghubble/go-twitter](https://github.com/dghubble/go-twitter) library to communicate with the Twitter API. Many thanks to both projects, you are doing a great job!
