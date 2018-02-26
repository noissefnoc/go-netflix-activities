netflix-activities
==================

[![CircleCI](https://circleci.com/gh/noissefnoc/go-netflix-activities/tree/master.svg?style=svg)](https://circleci.com/gh/noissefnoc/go-netflix-activities/tree/master)

`netflix-activities` fetches Netflix User's Own Activities by Scraping.


Usage
-----

To use `netflix-activities` is simple. After create Netflix authentication settings (see more on [Netflix Authentication Settings](https://github.com/noissefnoc/go-netflix-activities#netflix-authentication-settings) section) file, and run the following command on the same directory:

```
$ netflix-activities [option]
```

and you can get activities(currently viewing history only) as follow:

```
# View Date   Video Title
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: もっと強引に
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: 地下迷路
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: 壁マニア
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: 迷いの中で
2018/02/25      私立探偵ダーク・ジェントリー: シーズン1: 新たな地平
... 
```

I'm sorry for illustrating in Japanese but I can only see Japanese version.



### Netflix Authentication Settings

To use `netflix-activities`, you need to write Netflix Authentication Settings(email and password) at config file.

The config file format is [toml](https://github.com/toml-lang/toml) as follow:

```settings.toml
[Auth]
email = "YOUR_EMAIL"
password = "YOUR_PASSWORD"
```


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
$ go get -u github.com/noissefnoc/netflix-activities
```


Author
------

[Kota Saito](https://github.com/noissefnoc)