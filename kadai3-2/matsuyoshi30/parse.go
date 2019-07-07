package godl

import (
	"errors"
	"flag"
	"fmt"
)

type Option struct {
	Rt     uint
	URL    string
	Output string
}

const USAGE = "usage: godl -url url [-r set num of goroutine] [-o output filename]"

func ParseFlag(args ...string) (*Option, error) {
	fs := flag.NewFlagSet("godl", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(USAGE)
	}

	url := fs.String("url", "", "set URL for download")
	r := fs.Uint("r", 2, "set number of goroutine")
	o := fs.String("o", "output", "set output filename")
	fs.Parse(args)

	if *url == "" {
		fs.Usage()
		return nil, errors.New("set URL")
	}

	return &Option{Rt: *r, URL: *url, Output: *o}, nil
}
