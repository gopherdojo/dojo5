package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/manhdaovan/pkg/typinggame"
	"github.com/pkg/errors"
)

var defaultSentences = []string{
	"We need to rent a room for our party.",
	"The waves were crashing on the shore; it was a lovely sight.",
	"Where do random thoughts come from?",
	"She borrowed the book from him many years ago and hasn't yet returned it.",
	"I will never be this young again. Ever. Oh damn… I just got older.",
}

func main() {
	ca := parseCliArgs()
	if err := ca.validate(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		printHelp()
		os.Exit(1)
	}

	sampleSentences, err := getSamplesFromFiles(ca.sentencesFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		printHelp()
		os.Exit(1)
	}

	game := typinggame.TypingGame{
		Duration:  time.Duration(ca.duration) * time.Second,
		Sentences: sampleSentences,
	}
	exitReason := game.Start()
	correct := game.CorrectSentences()

	fmt.Println("")
	fmt.Println("")
	fmt.Println(exitReason)
	fmt.Println("\n ----- Result ----- ")
	fmt.Printf("correct %d sentences", correct)
}

type cliArgs struct {
	duration      uint64 // seconds
	sentencesFile string
}

func (ca *cliArgs) validate() error {
	if ca.sentencesFile != "" {
		if _, err := os.Open(ca.sentencesFile); err != nil {
			return errors.Wrap(err, "invalid sample sentences file")
		}
	}

	return nil
}

func parseCliArgs() *cliArgs {
	ca := cliArgs{}

	flag.Uint64Var(&ca.duration, "t", uint64(30), "duration for game in seconds")
	flag.StringVar(&ca.sentencesFile, "f", "", "sample sentences file")
	flag.Parse()

	return &ca
}

func getSamplesFromFiles(file string) ([]string, error) {
	var sentences []string

	if file == "" { // no given file
		sentences = defaultSentences
		return sentences, nil
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}

	return sentences, nil
}

func printHelp() {
	helpStr := `
tpg -- typing game in limitation of time
usage: $./bin/tpg [-t] [-f]
options:
	-t uint64
	    Duration of game in seconds. Default is 30.
	-f string
		sample sentences file path. Default is blank.
`
	fmt.Println(helpStr)
}
