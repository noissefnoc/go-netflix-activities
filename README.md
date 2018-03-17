netflix-activities
==================

[![GitHub release](http://img.shields.io/github/release/noissefnoc/go-netflix-activities.svg?style=flat-square)][release]
[![CircleCI](https://circleci.com/gh/noissefnoc/go-netflix-activities/tree/master.svg?style=svg)](https://circleci.com/gh/noissefnoc/go-netflix-activities/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/noissefnoc/go-netflix-activities/badge.svg?branch=master)](https://coveralls.io/github/noissefnoc/go-netflix-activities?branch=master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/noissefnoc/go-netflix-activities/releases
[license]: https://github.com/noissefnoc/embulk-output-gsheets/blob/master/LICENSE.txt

`netflix-activities` fetches Netflix User's Own Activities by Scraping.

**NOTE**: THIS PROGRAM IS NOT WORKING NOW(2018/03/17-). Maybe Netflix login page has some change. I'll fix next version(v0.0.7).

Usage
-----

To use `netflix-activities` is simple. After create Netflix authentication settings (see more on [Netflix Authentication Settings](https://github.com/noissefnoc/go-netflix-activities#netflix-authentication-settings) section) file, and run the following command on the same directory:

```
$ netflix-activities [option]
```

and you can get activities(currently viewing history only) as follow:

```
# View Date   Video Title   Video URL
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: もっと強引に https://www.netflix.com/title/80132376  
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: 地下迷路 https://www.netflix.com/title/80132376
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: 壁マニア https://www.netflix.com/title/80132375
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: 迷いの中で    https://www.netflix.com/title/80132374
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: 新たな地平    https://www.netflix.com/title/80132373
... 
```

I'm sorry for illustrating in Japanese but I can only see Japanese version.

**NOTE**: `Video URL` added from v0.0.4

### Netflix Authentication Settings

To use `netflix-activities`, you need to write Netflix Authentication Settings(email and password) by configuration file or environmental variables.

#### Configuration file

Configuration file format is [toml](https://github.com/toml-lang/toml) as follow:

```settings.toml
[Auth]
email = "YOUR_EMAIL"
password = "YOUR_PASSWORD"
```

#### Environmental variables

Environmental variable definition as follow:

```
export NF_EMAIL="YOUR_EMAIL"
export NF_PASSWORD="YOUR_PASSORD"
```

It's convenient for running app some PaaS.

Options
--------

You can set some options:

```
$ netflix-activities \
    -conf PATH  \        # Set config file path (default: settings.toml)
    -data PATH  \        # Set data file path (default: viewing_history.json)
    -limit NUM  \        # Set amount of displaying viewing history (default: 10)
    -expire NUM \        # Set amount of minutes keeping viewing history before you get (default: 60)
    -format TYPE \       # Set display type (simple/csv/table. default: simple)
```

you can also display version and usage:

```
# version
$ netflix-activities -version

# usage
$ netflix-activities -help
```

Install
--------

If you are macOS user, you can install via [Homebrew](https://brew.sh/).

```
$ brew tap noissefnoc/homebrew-go-netflix-activities
$ brew install go-netflix-activities
```

Other OS users can download binary from [release page](https://github.com/noissefnoc/go-netflix-activities/releases).

And you can also use `go get`

```
$ go get -u github.com/noissefnoc/go-netflix-activities
```


Author
------

[Kota Saito](https://github.com/noissefnoc)
