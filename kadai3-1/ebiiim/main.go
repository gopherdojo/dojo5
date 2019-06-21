package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/quiz"
	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/typing"
	"github.com/pkg/errors"
)

func game(quizFilePath string, timeSec int, randSeed int64) {
	bc := context.Background()
	ctx, cancel := context.WithCancel(bc)
	defer cancel()
	nextQuizCh := make(chan struct{})

	// load quiz
	reader, err := os.Open(quizFilePath)
	if err != nil {
		err = errors.Wrap(err, "failed to open file")
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	quizLoader, err := quiz.NewJSONLoader(reader, randSeed)
	if err != nil {
		err = errors.Wrap(err, "failed to generate JSONLoader")
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	tg := typing.NewTypingGame(ctx, nextQuizCh, quizLoader, os.Stdin, timeSec)

	// request the first quiz
	nextQuizCh <- struct{}{}
QuizLoop:
	for {
		select {
		case q := <-tg.QuizChannel:
			fmt.Fprintf(os.Stdout, "Please type: %v\n", q.Text)
			tg.QuizList = append(tg.QuizList, q)
		case str := <-tg.AnswerChannel:
			tg.AnswerList = append(tg.AnswerList, str)
			nextQuizCh <- struct{}{}
		case <-tg.TimerChannel:
			fmt.Fprintf(os.Stdout, "Time is up!\n")
			break QuizLoop
		}
	}

	// show the result
	grade := tg.CalcGrade()
	fmt.Fprintf(os.Stdout, "\nTotal Time: %d s\nAnswered: %d\nCorrects: %d\nCorrect Rate: %.0f%%\nWPM: %.0f\n", grade.TotalTime, grade.Answered, grade.Corrects, 100*grade.CorrectRate, grade.WordsPerMinute)
}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-quiz] [-time]\n", os.Args[0])
		flags.PrintDefaults()
	}
	var (
		timeSec      = flags.Int("time", 30, "Seconds the game takes.")
		quizFilePath = flags.String("quiz", "pkg/testdata/quiz.json", "The path of quiz file.")
	)
	err := flags.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	game(*quizFilePath, *timeSec, time.Now().UnixNano())
}
