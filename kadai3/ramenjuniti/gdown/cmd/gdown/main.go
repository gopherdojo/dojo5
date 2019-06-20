package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/gopherdojo/dojo5/kadai3/ramenjuniti/gdown"
)

func main() {
	p := flag.Int("p", runtime.NumCPU(), "並列数")
	flag.Parse()

	if *p < 1 {
		fmt.Fprintln(os.Stderr, "p cannot be less than 1")
		os.Exit(1)
	}

	url := flag.Arg(0)
	if url == "" {
		fmt.Fprintln(os.Stderr, "please input URL")
		os.Exit(1)
	}

	c, err := gdown.New(url, *p)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err = c.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
