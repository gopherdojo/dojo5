package typing

import (
	"context"
	"io"

	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/quiz"
	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/timer"
)

// Game struct holds channels and lists for the typing game.
type Game struct {
	// receive the next quiz
	QuizChannel <-chan *quiz.Quiz
	// scan input
	AnswerChannel <-chan *quiz.Answer
	// receive time's up
	TimerChannel <-chan interface{}
	// a list of quizzes which appeared
	QuizList []*quiz.Quiz
	// a list of input texts with timestamps
	AnswerList []*quiz.Answer
	timeSec    int
}

// Grade struct holds the result of the typing game.
type Grade struct {
	// total time (sec)
	TotalTime int
	// number of answered quizzes
	Answered int
	// number of correct answers
	Corrects int
	// the number of correctly answered words per a minute
	WordsPerMinute float64
	// percentage of correct answers (0.0 -- 1.0)
	CorrectRate float64
}

// NewTypingGame initializes Game.
func NewTypingGame(ctx context.Context, nextQuizCh <-chan interface{}, quizLoader quiz.Loader, reader io.Reader, timeSec int) *Game {
	g := &Game{timeSec: timeSec}
	g.QuizChannel = quiz.MakeQuizChannel(ctx, nextQuizCh, quizLoader)
	g.AnswerChannel = quiz.MakeAnswerChannel(ctx, reader)
	g.TimerChannel = timer.MakeChannel(timeSec)
	return g
}

// CalcGrade calculates the current score and returns Grade.
func (g *Game) CalcGrade() *Grade {
	totalSec := g.timeSec
	answered := len(g.AnswerList)
	numCorrects := g.countCorrectAnswers()
	numMinutes := float64(totalSec) / 60.0
	wpm := float64(numCorrects) / float64(numMinutes)
	cr := float64(numCorrects) / float64(answered)
	gr := &Grade{TotalTime: totalSec, Answered: answered, Corrects: numCorrects, WordsPerMinute: wpm, CorrectRate: cr}
	return gr
}

func (g *Game) countCorrectAnswers() (c int) {
	for idx, ans := range g.AnswerList {
		if isContain(g.QuizList[idx].Answers, ans.Text) {
			c++
		}
	}
	return
}

func isContain(ss []string, s string) bool {
	for _, e := range ss {
		if s == e {
			return true
		}
	}
	return false
}
