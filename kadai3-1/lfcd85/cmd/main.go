package main

import (
	"fmt"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/typinggame"
	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/words"
)

func main() {
	path := "./testdata/go_standard_library.txt" // FIXME: move to options
	words, err := words.Import(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	g := typinggame.Game{
		Words:     words,
		TimeLimit: 30 * time.Second,
	}

	if err := typinggame.Execute(g); err != nil {
		fmt.Println("error:", err)
		return
	}
}
