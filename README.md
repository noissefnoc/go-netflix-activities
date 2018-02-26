netflix-activities
==================

`netflix-activities` fetches Netflix User's Own Activities by Scraping.


Usage
-----

To use `netflix-activities` is simple. After create Netflix authentication settings (see more on [Netflix Authentication Settings](https://github.com/noissefnoc/netflix-activities#netflix-authentication-settings) section) file, and run the following command on the same directory:

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
email = <YOUR_EMAIL>
password = <YOUR_PASSWORD>
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

If you are macOS user, you need to install ChromeDriver 2.13+ via [Homebrew](https://brew.sh/) first.

```
$ brew install chromedriver
```

And then, you can use `go get`

```
$ go get -u github.com/noissefnoc/netflix-activities
```

I currently prepare download binary on release page and `homebrew`.


Author
------

[Kota Saito](https://github.com/noissefnoc)