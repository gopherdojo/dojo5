package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gopherdojo/dojo5/kadai1/sh-tatsuno/gophoto/conv"
	"github.com/gopherdojo/dojo5/kadai1/sh-tatsuno/gophoto/dir"
)

const (
	// ExitCodeOK : exit code
	ExitCodeOK = iota

	// ExitCodeError : error code
	ExitCodeError
)

var (
	// SEP : separator of each line
	SEP = []byte("\n")
)

func usage() {
	io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}

const usageText = `this is image convert library by go.

In normal usage, you should set -d for directory and -i for input extension.
You also have to set output extension by -o.
You can also set maximum nuber you want to convert by set n.
current available extensions are jpg, jpeg, png, and gif.

Example:
	gophoto -d dir -i .png -o .jpeg -n 10

`

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	var n int
	var dirName, input, output string

	// args
	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)
	flags.Usage = usage
	flags.IntVar(&n, "n", 10, "number of maximum images to convert; default is 10.")
	flags.StringVar(&dirName, "d", "", "directory path.")
	flags.StringVar(&input, "i", "", "input extension.")
	flags.StringVar(&output, "o", "", "output extension.")
	flags.Parse(args)

	if dirName == "" {
		usage()
		fmt.Printf("Expected dir path\n")
		return ExitCodeError
	}

	// lookup
	var pathList []string
	pathList, err := dir.Lookup(dirName, input, pathList)
	if err != nil {
		fmt.Printf("can not open file, %v\n", err)
		return ExitCodeError
	}

	// convert
	for i, path := range pathList {
		if i > n {
			break
		}

		// convert in each file
		img, err := conv.NewImg(path)
		if err != nil {
			fmt.Printf("can not generate img instance, %v\n", err)
			return ExitCodeError
		}

		if err := img.Replace(output); err != nil {
			fmt.Printf("can not convert file, %v\n", err)
			return ExitCodeError
		}

	}

	return ExitCodeOK

}
