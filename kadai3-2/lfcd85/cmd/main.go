package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"

	"github.com/gopherdojo/dojo5/kadai3-2/lfcd85/mypget"
)

func main() {
	flag.Parse()
	urlStr := flag.Arg(0)
	if urlStr == "" {
		err := errors.New("URL is not inputted")
		fmt.Println(err) // TODO: beautify error handling
		return
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = mypget.New(url).Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
