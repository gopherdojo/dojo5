package quiz

import (
	"bufio"
	"context"
	"io"
	"math/rand"
	"time"
)

// FIXME: inputCh does not close because scanner.Scan() is blocking.
func MakeAnswerChannel(ctx context.Context, input io.Reader) <-chan Answer {
	inputCh := make(chan Answer)
	scanner := bufio.NewScanner(input)

	go func() {
		//defer fmt.Println("closed AnswerChannel")
		defer close(inputCh)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				scanner.Scan()
				inputCh <- Answer{Timestamp: time.Now(), Text: scanner.Text()}
			}
		}
	}()
	return inputCh
}

type Quiz struct {
	Timestamp time.Time
	Text      string
	Answers   []string
}

type Answer struct {
	Timestamp time.Time
	Text      string
}

func MakeQuizChannel(ctx context.Context, next <-chan interface{}, quizLoader Loader) <-chan Quiz {
	quizCh := make(chan Quiz)
	go func() {
		//defer fmt.Println("closed QuizChannel")
		defer close(quizCh)
		for {
			select {
			case <-next:
				quizCh <- quizLoader.Next()
			case <-ctx.Done():
				return
			}
		}
	}()
	return quizCh
}

type Loader interface {
	Next() Quiz
}

type DummyLoader struct{}

func (l *DummyLoader) Next() Quiz {
	return Quiz{Timestamp: time.Now(), Text: "abc", Answers: []string{"abc"}}
}

type CSVLoader struct {
	QuizList []Quiz
	random   *rand.Rand
}

func (l *CSVLoader) Next() Quiz {
	n := len(l.QuizList)
	return l.QuizList[l.random.Intn(n)]

}

func NewCSVLoader(path string, randSeed int64) (*CSVLoader, error) {
	l := &CSVLoader{}
	l.random = rand.New(rand.NewSource(randSeed))
	// TODO: load csv
	return l, nil
}
