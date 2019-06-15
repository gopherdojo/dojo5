package parse

import (
	"flag"
	"fmt"
	"os"
)

const USAGE = `
NAME:
   gotp - Simple Typing Game in Go

USAGE:
   gotp [-lt limit time (sec)] [-wf source word file]

VERSION:
   0.1.0

GLOBAL OPTIONS:
  -lt             set limit time (sec) [default: 20]
  -wf             specify source word file [default: default.txt]
  --help, -h      show help
`

type Option struct {
	LimitTime int
	WordFile  string
}

func ParseFlag(args ...string) *Option {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(USAGE)
	}

	lt := fs.Int("lt", 20, "set limit time")
	wf := fs.String("wf", "default.txt", "specify source word file")
	fs.Parse(args)

	return &Option{LimitTime: *lt, WordFile: *wf}
}
