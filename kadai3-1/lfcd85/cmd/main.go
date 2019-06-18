package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/typinggame"
	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/words"
)

const path = "./testdata/go_standard_library.txt"

func main() {
	timeLimit := flag.Int("t", 30, "Time limit of the game (sec)")
	flag.Parse()

	words, err := words.Import(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	g := typinggame.Game{
		Words:     words,
		TimeLimit: time.Duration(*timeLimit) * time.Second,
	}

	if err := typinggame.Execute(g); err != nil {
		fmt.Println("error:", err)
		return
	}
}
