package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/cli"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dirconv"
)

func main() {
	dirconv.Logger = log.New(os.Stdout, fmt.Sprintf("%s ", os.Args[0]), log.LstdFlags)

	dc, err := cli.ParseArgs(os.Args)
	if err != nil {
		if cli.IsInvalidArgs(err) {
			// show usage if no dir specified
			fmt.Fprintf(os.Stdout, "%s\n", cli.Usage)
		} else {
			// unexpected error
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		os.Exit(0)
	}
	_, err = dc.Convert()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(0)
	}
}
