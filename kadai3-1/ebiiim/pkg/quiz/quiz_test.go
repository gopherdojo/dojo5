package quiz_test

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/quiz"
)

//func TestMakeAnswerChannel(t *testing.T) {
//
//}

func TestMakeQuizChannel(t *testing.T) {
	var (
		loader1   = &quiz.DummyLoader{}
		dummyQuiz = quiz.Quiz{Text: "abc", Answers: []string{"abc"}}
	)
	cases := []struct {
		name   string
		loader quiz.Loader
		want1  quiz.Quiz
		want2  quiz.Quiz
	}{
		{name: "normal", loader: loader1, want1: dummyQuiz, want2: dummyQuiz},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(context.Background())
			nextQuizCh := make(chan interface{})
			quizCh := quiz.MakeQuizChannel(ctx, nextQuizCh, c.loader)

			// get the 1st quiz
			nextQuizCh <- struct{}{}
			q, ok := <-quizCh
			if !ok {
				t.Error("invalid channel status")
			}
			if !cmp.Equal(q, c.want1, cmpopts.IgnoreFields(q, "Timestamp")) {
				t.Errorf("invalid answers: want %v got %v", c.want1, q)
			}

			// get the 2nd quiz
			nextQuizCh <- struct{}{}
			q, ok = <-quizCh
			if !ok {
				t.Error("invalid channel status")
			}
			if !cmp.Equal(q, c.want2, cmpopts.IgnoreFields(q, "Timestamp")) {
				t.Errorf("invalid answers: want %v got %v", c.want2, q)
			}

			// close the channel
			cancel()
			_, ok = <-quizCh
			if ok {
				t.Error("failed to close the channel")
			}
		})
	}
}

func TestDummyLoader_Next(t *testing.T) {
	cases := []struct {
		name    string
		quiz    string
		answers []string
	}{
		{name: "normal", quiz: "abc", answers: []string{"abc"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			l := &quiz.DummyLoader{}
			q := l.Next()
			if q.Text != c.quiz {
				t.Errorf("invalid quiz text: want %v got %v", c.quiz, q.Text)
			}
			if !cmp.Equal(q.Answers, c.answers) {
				t.Errorf("invalid answers: want %v got %v", c.answers, q.Answers)
			}
		})
	}
}

//func TestNewJSONLoader(t *testing.T) {
//
//}

// TODO: divide this test into `TestNewJSONLoader` and `TestJSONLoader_Next`
func TestJSONLoader_Next(t *testing.T) {
	const (
		jQuiz        = `[{"text": "hello", "answers": ["hello"]}, {"text": "world", "answers": ["world", "world!"]}]`
		jInvalidQuiz = `[{"A": "hello", "B": ["hello"]}, {"A": "world", "B": ["world", "world!"]}]`
	)
	var (
		qr1      = strings.NewReader(jQuiz)
		qr2      = strings.NewReader(jInvalidQuiz)
		qr3, err = os.Open("../testdata/quiz.json")
	)
	if err != nil {
		t.Fatal("failed to load file")
	}

	cases := []struct {
		name    string
		reader  io.Reader
		seed    int64
		quiz    string
		answers []string
		isErr   bool
	}{
		{name: "normal_str", reader: qr1, seed: 1, quiz: "world", answers: []string{"world", "world!"}, isErr: false},
		{name: "invalid_json", reader: qr2, seed: 1, quiz: "", answers: nil, isErr: true},
		{name: "normal_file", reader: qr3, seed: 1, quiz: "world", answers: []string{"world"}, isErr: false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			l, err := quiz.NewJSONLoader(c.reader, c.seed)
			if err != nil {
				// non-normal cases
				if !c.isErr {
					t.Errorf("invalid error status %v(c.isErr) %v(err)", c.isErr, err)
				} else {
					// do nothing
				}
			} else {
				// normal cases
				if c.isErr {
					t.Errorf("invalid error status %v(c.isErr) %v(err)", c.isErr, err)
				} else {
					q := l.Next()
					if q.Text != c.quiz {
						t.Errorf("invalid quiz text: want %v got %v", c.quiz, q.Text)
					}
					if !cmp.Equal(q.Answers, c.answers) {
						t.Errorf("invalid answers: want %v got %v", c.answers, q.Answers)
					}
				}
			}
		})
	}
}
