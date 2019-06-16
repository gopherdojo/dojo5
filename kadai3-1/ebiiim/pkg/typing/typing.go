package typing

import (
	"context"
	"io"

	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/quiz"
	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/timer"
)

type Game struct {
	// receive the next quiz
	QuizChannel <-chan quiz.Quiz
	// scan input
	AnswerChannel <-chan quiz.Answer
	// receive time's up
	TimerChannel <-chan interface{}
	// a list of quizzes which appeared
	QuizList []quiz.Quiz
	// a list of input texts with timestamps
	AnswerList []quiz.Answer
	timeSec    int
}

type Grade struct {
	WordsPerMinute float64
	CorrectRate    float64
}

func NewTypingGame(ctx context.Context, nextQuizCh <-chan interface{}, quizLoader quiz.Loader, reader io.Reader, timeSec int) *Game {
	g := &Game{timeSec: timeSec}
	g.QuizChannel = quiz.MakeQuizChannel(ctx, nextQuizCh, quizLoader)
	g.AnswerChannel = quiz.MakeAnswerChannel(ctx, reader)
	g.TimerChannel = timer.MakeChannel(timeSec)
	return g
}

func (g *Game) CalcGrade() *Grade {
	numCorrects := float64(g.countCorrectAnswers())
	numMinutes := float64(g.timeSec) / 60.0
	cr := numCorrects / float64(len(g.AnswerList))
	wpm := numCorrects / numMinutes
	gr := &Grade{WordsPerMinute: wpm, CorrectRate: cr}
	return gr
}

func (g *Game) countCorrectAnswers() (c int) {
	for idx, ans := range g.AnswerList {
		// TODO: check ans.Text in Quiz.Answers
		if ans.Text == g.QuizList[idx].Answers[0] {
			c++
		}
	}
	return
}
