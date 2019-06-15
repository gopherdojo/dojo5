package main

import (
	"fmt"

	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/typinggame"
)

func main() {
	// create words from input file (abstracted by io.Reader if possible)

	// start time counting by time.After or context.WithTimeout

	// output word to stdout
	// get one line from stdin
	// judge whether two words are the same
	// if so, add count of correct answers

	// when time limit has come, show the count of correct answers

	if err := typinggame.Execute(); err != nil {
		fmt.Println("error:", err)
		return
	}
}
