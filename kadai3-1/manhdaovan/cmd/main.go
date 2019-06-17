package main

import (
	"flag"
	"fmt"

	"github.com/gopherdojo/dojo5/kadai3-1/manhdaovan/pkg/typinggame"
)

func main() {
	ca := parseCliArgs()
	game := typinggame.TypingGame{
		Duration: ca.duration,
	}
	done := game.Start()
	select {
	case reason := <-done:
		correct := game.CorrectSentences()
		fmt.Println("\n")
		fmt.Println(reason)
		fmt.Println("\n ----- Result ----- ")
		fmt.Println("correct %d sentences", correct)
	}
}

type cliArgs struct {
	duration uint64 // seconds
}

func parseCliArgs() *cliArgs {
	ca := cliArgs{}

	flag.Uint64Var(&ca.duration, "-t", uint64(5), "duration for game in seconds")
	flag.Parse()

	return &ca
}
