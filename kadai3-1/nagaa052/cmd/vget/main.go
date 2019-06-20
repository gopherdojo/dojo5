package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var timeout int
	flag.IntVar(&timeout, "t", 30, "Timeout Seconds")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, `
tgame is a Typing Game
Usage:
  tgame [option]
Options:
`)
	flag.PrintDefaults()
}
