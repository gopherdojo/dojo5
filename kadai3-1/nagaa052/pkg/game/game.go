/*
Typing game main logic
*/
package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/questions"
)

const (
	ExitOK = iota
	ExitError
)

// Game is manages game information
type Game struct {
	qs *questions.Questions
	*Result
	Options
	inStream             io.Reader
	outStream, errStream io.Writer
}

// Options is a specifiable option.
type Options struct {
	TimeUpSecond int
	IsColor      bool
}

// DefaultOptions is the default value of Options.
var DefaultOptions = Options{
	TimeUpSecond: 30,
	IsColor:      false,
}

// Result is manages result information.
type Result struct {
	Questions    []*questions.Question
	CorrectCount int
}

// Print is Output the result.
func (r *Result) Print(out io.Writer) {
	cr := float64(r.CorrectCount) / float64(len(r.Questions)) * 100
	fmt.Fprintln(out, "========================")
	fmt.Fprintf(out, "Correct Count: %d\n", r.CorrectCount)
	fmt.Fprintf(out, "Correct Rate: %.1f％\n", cr)
}

// New is Generate a new game.
func New(opt Options, inStream io.Reader, outStream, errStream io.Writer) (*Game, error) {
	if opt.TimeUpSecond <= 0 {
		opt.TimeUpSecond = DefaultOptions.TimeUpSecond
	}

	qs, err := questions.New()
	if err != nil {
		return nil, err
	}

	return &Game{
		qs:        qs,
		Result:    &Result{},
		Options:   opt,
		inStream:  inStream,
		outStream: outStream,
		errStream: errStream,
	}, nil
}

// Start is Start Game
func (g *Game) Start() int {
	g.printStart()

	bc := context.Background()
	ctx, cannel := context.WithTimeout(bc, time.Duration(g.TimeUpSecond)*time.Second)
	defer cannel()

	dst := g.startScanner()

	return func() int {
		for {
			q, err := g.getQuestion()
			if err != nil {
				fmt.Fprintf(g.errStream, "%v\n", err.Error())
				return ExitError
			}
			g.printQuestion(q.Word)

			select {
			case <-ctx.Done():
				g.printTimeOut()
				return ExitOK
			case input := <-dst:
				if q.IsCorrect(input) {
					g.Result.CorrectCount++
				}
			}
		}
	}()
}

func (g *Game) startScanner() <-chan string {

	dst := make(chan string)

	go func() {
		scanner := bufio.NewScanner(g.inStream)
		defer close(dst)

		for scanner.Scan() {
			dst <- scanner.Text()
		}
	}()

	return dst
}

func (g *Game) getQuestion() (*questions.Question, error) {

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(g.qs.GetSize())

	q, err := g.qs.GetOne(index)
	if err != nil {
		return nil, err
	}

	g.Result.Questions = append(g.Result.Questions, q)
	return q, nil
}

func (g *Game) printStart() {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(g.outStream, "%d", i)
		time.Sleep(250 * time.Millisecond)

		fmt.Fprint(g.outStream, ".")
		time.Sleep(250 * time.Millisecond)
		fmt.Fprint(g.outStream, ".")
		time.Sleep(250 * time.Millisecond)
		fmt.Fprint(g.outStream, ".")
		time.Sleep(250 * time.Millisecond)
	}
	fmt.Fprintln(g.outStream, "Start!!")
	fmt.Fprintln(g.outStream, "========================")
	time.Sleep(500 * time.Millisecond)
}

func (g *Game) printQuestion(word string) {
	if g.IsColor {
		cFPrint(green, g.outStream, fmt.Sprintf("%s", word))
	} else {
		fmt.Fprintf(g.outStream, "%s", word)
	}
}

func (g *Game) printTimeOut() {
	if g.IsColor {
		cFPrint(red, g.outStream, "\nTimeUp!!!")
	} else {
		fmt.Fprintf(g.outStream, "\nTimeUp!!!")
	}

	time.Sleep(1 * time.Second)
	g.Result.Print(g.outStream)
}
