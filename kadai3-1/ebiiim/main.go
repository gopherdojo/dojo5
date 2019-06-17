package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/quiz"
	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/typing"
	"github.com/pkg/errors"
)

func main() {
	bc := context.Background()
	ctx, cancel := context.WithCancel(bc)
	defer cancel()
	nextQuizCh := make(chan interface{})

	// load quiz
	// quizLoader := &quiz.DummyLoader{}
	reader, err := os.Open("pkg/testdata/quiz.json")
	if err != nil {
		err = errors.Wrap(err, "failed to open file")
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	quizLoader, err := quiz.NewJSONLoader(reader, 1)
	if err != nil {
		err = errors.Wrap(err, "failed to generate JSONLoader")
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	tg := typing.NewTypingGame(ctx, nextQuizCh, quizLoader, os.Stdin, 5)

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
	fmt.Fprintf(os.Stdout, "WordsPerMinute: %.2f\nCorrectRate: %.2f\n", grade.WordsPerMinute, grade.CorrectRate)
}
