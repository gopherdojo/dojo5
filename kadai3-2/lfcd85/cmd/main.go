package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/gopherdojo/dojo5/kadai3-2/lfcd85/mypget"
)

func main() {
	splitNum := flag.Int("n", 8, "Number of splitted downloads")
	flag.Parse()
	urlStr := flag.Arg(0)
	if urlStr == "" {
		err := errors.New("URL is not inputted")
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}

	if err := mypget.New(url, *splitNum).Execute(nil); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}
