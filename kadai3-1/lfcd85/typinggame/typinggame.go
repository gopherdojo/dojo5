package typinggame

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
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
	g.run(inputChannel(os.Stdin), os.Stdout)
	return nil
}

func (g *Game) run(ch <-chan string, w io.Writer) {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, g.TimeLimit)
	defer cancel()

	fmt.Fprintln(w, "Let's type the standard package names! ( Time limit:", g.TimeLimit, ")")

	var score int
	rand.Seed(time.Now().UnixNano())
	word := g.Words[rand.Intn(len(g.Words))]
LOOP:
	for {
		fmt.Fprintln(w, ">", word)
		select {
		case input := <-ch:
			if input == word {
				score++
				fmt.Fprintln(w, input, "... OK! current score:", score)
				word = g.Words[rand.Intn(len(g.Words))]
			} else {
				fmt.Fprintln(w, input, "... NG: try again.")
			}
		case <-ctx.Done():
			fmt.Fprintln(w)
			fmt.Fprintln(w, g.TimeLimit, "has passed: you correctly typed", score, "package(s)!")
			break LOOP
		}
	}
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
