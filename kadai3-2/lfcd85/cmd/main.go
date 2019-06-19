package main

import (
	"fmt"
	"net/url"

	"github.com/gopherdojo/dojo5/kadai3-2/lfcd85/mypget"
)

func main() {
	// FIXME: get URL from option
	url, err := url.Parse("https://golang.org/doc/gopher/frontpage.png")
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
