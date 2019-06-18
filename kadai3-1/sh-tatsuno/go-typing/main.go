package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/sh-tatsuno/go-typing/typing_game"
	"github.com/gopherdojo/dojo5/kadai3-1/sh-tatsuno/go-typing/words"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

func main() {

	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	timeLimit := f.Int("t", 60, "Time limit of the game (sec): default 60sec")
	wordPath := f.String("f", "animals", "type of words. details in words directory: default animals")
	f.Parse(args)

	filePath := "words/wordlist/" + *wordPath + ".txt"
	gameWords, err := words.Import(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "words import error. err: %v", err)
		return ExitCodeError
	}

	g := typing_game.NewTypingGame(gameWords, time.Duration(*timeLimit)*time.Second)
	if err := g.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Execute error. err: %v", err)
		return ExitCodeError
	}

	return ExitCodeOK
}
