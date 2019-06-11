package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/matsuyoshi30/dojo5/kadai2/matsuyoshi30/conv"
	"github.com/matsuyoshi30/dojo5/kadai2/matsuyoshi30/opt"
)

// usage
const usage = `
NAME:
   goconv - image format converter written in Go

USAGE:
   goconv [-from before image format] [-to after image format] path/to/dir

VERSION:
   0.1.0

GLOBAL OPTIONS:
   -from           specify format before converted (jpg, png, gif) [DEFAULT: jpg]
   -to             specify format after converted  (png, jpg, gif) [DEFAULT: png]
   --help, -h      show help
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	var from, to string
	var sh bool
	flag.StringVar(&from, "from", "jpeg", "Choose format before converted")
	flag.StringVar(&to, "to", "png", "Choose format after converted")
	flag.BoolVar(&sh, "h", false, "Show help")
	flag.Parse()

	if sh {
		fmt.Println(usage)
		return
	}

	fromtype, err := conv.SelectFormat(from)
	if err != nil {
		logError(err, true)
	}
	totype, err := conv.SelectFormat(to)
	if err != nil {
		logError(err, true)
	}
	dirlist := opt.OptPath(flag.Args())

	if len(dirlist) < 1 {
		fmt.Println(usage)
	} else {
		for _, d := range dirlist {
			err := conv.Imgconv(fromtype, totype, d)
			logError(err, false)
		}
	}
}

func logError(err error, stop bool) {
	if err != nil {
		log.Fatal(err)
		if stop {
			os.Exit(1)
		}
	}
}
