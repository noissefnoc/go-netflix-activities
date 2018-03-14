package main

import (
	"fmt"
	"os"
	"io"
	"flag"
)

var (
	// Version is variable for build version
	Version string
	// Revision is variable for build revision
	Revision string
)

var (
	c = flag.String("conf", "settings.toml", "configuration file path")
	d = flag.String("data", "viewing_history.json", "viewing history data file path")
	l = flag.Int("limit", 10, "viewing history display record number limit")
	e = flag.Int("expire", 60, "viewing history expire duration (minutes)")
	f = flag.String("format", "simple", "print format(simple/csv/table. default: simple)")
	v = flag.Bool("version", false, "print version")
	h = flag.Bool("help", false, "print usage")
)

func printVersion() {
	fmt.Printf("version: %s(%s)\n", Version, Revision)
}

func run(w io.Writer) (int) {
	config := &Config{}
	if err := config.Read(*c); err != nil {
		fmt.Printf("failed to read config:%v\n", err)
		return 1
	}

	vh := &ViewingHistory{}
	filePath := *d

	if vh.ExistData(filePath) && !vh.Expire(filePath, *e) {
		if err := vh.LoadFromFile(filePath); err != nil {
			fmt.Printf("failed to parse viewing history(local):%v\n", err)
			return 1
		}
	} else {
		netflix := &Netflix{
			LoginURL:          "https://netflix.com/jp/login",
			ViewingHistoryURL: "https://www.netflix.com/wiviewingactivity"}

		if err := netflix.FetchViewingHistory(config.Auth.Email, config.Auth.Password);
			err != nil {
			fmt.Printf("failed to fetch viewing history:%v\n", err)
			return 1
		}

		if err := vh.LoadFromHTML(netflix.ViewingHistoryHTML); err != nil {
			fmt.Printf("failed to parse viewing history(remote):%v\n", err)
			return 1
		}
	}

	if err := vh.SaveData(filePath); err != nil {
		// even if save failed, print result
		fmt.Printf("failed to save viewing history:%v\n", err)
	}

	vh.Print(*l, *f, w)
	return 0
}

func main() {
	flag.Parse()

	if *v {
		printVersion()
		os.Exit(0)
	}

	if *h {
		flag.Usage()
		printVersion()
		os.Exit(0)
	}

	ret := run(os.Stdout)
	os.Exit(ret)
}
