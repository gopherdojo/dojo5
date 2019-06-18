package typinggame

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/hashicorp/go-multierror"
)

// Words stores a slice of words which is used for the game.
type Words []string

// Game struct holds the words and the time limits of the game.
type Game struct {
	Words     Words
	TimeLimit time.Duration
}

// Execute starts the game using standard input and output.
func Execute(g Game) error {
	err := g.run(inputChannel(os.Stdin), os.Stdout)
	return err
}

func (g *Game) run(ch <-chan string, w io.Writer) error {
	var result error

	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, g.TimeLimit)
	defer cancel()

	result = printWithMultiErr(w, result, "Let's type the standard package names! ( Time limit:", g.TimeLimit, ")")

	var score int
	rand.Seed(time.Now().UnixNano())
	word := g.Words[rand.Intn(len(g.Words))]
LOOP:
	for {
		result = printWithMultiErr(w, result, ">", word)
		select {
		case input := <-ch:
			if input == word {
				score++
				result = printWithMultiErr(w, result, input, "... OK! current score:", score)
				word = g.Words[rand.Intn(len(g.Words))]
			} else {
				result = printWithMultiErr(w, result, input, "... NG: try again.")
			}
		case <-ctx.Done():
			result = printWithMultiErr(w, result)
			result = printWithMultiErr(w, result, g.TimeLimit, "has passed: you correctly typed", score, "package(s)!")
			break LOOP
		}
	}

	return result
}

func inputChannel(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

func printWithMultiErr(w io.Writer, result error, a ...interface{}) error {
	if _, err := fmt.Fprintln(w, a...); err != nil {
		result = multierror.Append(result, err)
	}
	return result
}
