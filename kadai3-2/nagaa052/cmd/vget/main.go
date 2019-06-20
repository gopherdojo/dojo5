package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	vget "github.com/gopherdojo/dojo5/kadai3-2/nagaa052"
)

var outStream io.Writer = os.Stdout
var errStream io.Writer = os.Stderr

func main() {
	var timeout int
	flag.IntVar(&timeout, "t", 30, "Timeout Seconds")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	opt := vget.Options{
		TimeOut: time.Duration(timeout) * time.Second,
	}

	v, err := vget.New(args[0], opt, outStream, errStream)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		os.Exit(1)
	}
	os.Exit(v.Download())
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Parallel download
Usage:
  tgame [option]
Options:
`)
	flag.PrintDefaults()
}
