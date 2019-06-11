package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gopherdojo/dojo5/kadai2/nagaa052/internal/gocon"
)

var imageFormat string
var destDir string
var outStream io.Writer = os.Stdout
var errStream io.Writer = os.Stderr

func init() {
	flag.StringVar(&imageFormat,
		"f",
		gocon.DefaultOptions.FromFormat.String()+":"+gocon.DefaultOptions.ToFormat.String(),
		"Convert image format. The input format is [In]:[Out].")
	flag.StringVar(&destDir,
		"d",
		gocon.DefaultOptions.DestDir,
		"Destination directory.")
}

func main() {

	flag.Usage = usage
	flag.Parse()

	if opt, err := parseOptions(); err != nil {
		fmt.Printf("%s\n", err.Error())
		usage()
	} else {
		for _, srcDir := range flag.Args() {
			gc, err := gocon.New(srcDir, *opt, outStream, errStream)
			if err != nil {
				fmt.Printf("%v\n", err.Error())
				os.Exit(1)
			}

			os.Exit(gc.Run())
		}
	}
}

func parseOptions() (*gocon.Options, error) {
	format := strings.Split(imageFormat, ":")

	if len(format) != 2 {
		return nil, fmt.Errorf("invalid format")
	}

	from := gocon.ImgFormat(format[0])
	to := gocon.ImgFormat(format[1])

	if !from.Exist() || !to.Exist() {
		return nil, fmt.Errorf("Format that does not exist is an error")
	}

	if from == to {
		return nil, fmt.Errorf("The same format specification is an error")
	}

	return &gocon.Options{
		FromFormat: from,
		ToFormat:   to,
		DestDir:    destDir,
	}, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `
gocon is a tool for ...
Usage:
  gocon [option] <directory path>
Options:
`)
	flag.PrintDefaults()
}
