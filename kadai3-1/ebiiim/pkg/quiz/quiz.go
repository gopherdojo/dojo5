package quiz

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

// MakeAnswerChannel returns a read-only channel that passes data scanned from the given io.Reader.
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

// Quiz struct holds a quiz (with answers) and the time when the quiz is presented.
type Quiz struct {
	Timestamp time.Time
	Text      string
	Answers   []string
}

// Answer struct holds a answer text (entered by the user) and the time when it was created.
type Answer struct {
	Timestamp time.Time
	Text      string
}

// MakeQuizChannel returns a read-only channel, that provides Quiz when some data passed from the `next` channel.
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

// Loader interface provides a Next function that provides quizzes from a quiz database.
type Loader interface {
	Next() Quiz
}

// DummyLoader struct provides a dummy Loader for testing.
type DummyLoader struct{}

// Next returns a Quiz: {Timestamp: NOW, Text: "abc", Answers: ["abc"]}
func (l *DummyLoader) Next() Quiz {
	return Quiz{Timestamp: time.Now(), Text: "abc", Answers: []string{"abc"}}
}

// JSONLoader struct provides a Loader that uses JSON for the database.
type JSONLoader struct {
	QuizList []Quiz
	random   *rand.Rand
}

// Next returns a randomly loaded Quiz from QuizList.
func (l *JSONLoader) Next() Quiz {
	n := len(l.QuizList)
	q := l.QuizList[l.random.Intn(n)]
	q.Timestamp = time.Now()
	return q
}

// NewJSONLoader initializes JSONLoader.
func NewJSONLoader(reader io.Reader, randSeed int64) (*JSONLoader, error) {
	l := &JSONLoader{}
	l.random = rand.New(rand.NewSource(randSeed))
	jsonDecoder := json.NewDecoder(reader)
	err := jsonDecoder.Decode(&l.QuizList)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode json")
	}
	return l, nil
}
