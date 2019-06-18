package typing_game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type TypingGameInterface interface {
	Execute() error
}

type TypingGame struct {
	Words     []string
	TimeLimit time.Duration
}

// result
type resultList struct {
	total    int
	corrects int
}

func NewTypingGame(ws []string, t time.Duration) TypingGame {
	tg := TypingGame{
		Words:     ws,
		TimeLimit: t,
	}
	return tg
}

func (t *TypingGame) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(t.Words) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		t.Words[i], t.Words[j] = t.Words[j], t.Words[i]
	}
}

func (t TypingGame) Execute() error {

	r := &resultList{}
	t.shuffle()

	keyin := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			keyin <- s.Text()
		}
	}()

	fmt.Print("TYPING start\n")
	time.Sleep(500)
	timeup := time.After(t.TimeLimit)

LOOP:
	for _, word := range t.Words {

		fmt.Printf("-> %s\n", word)

		select {
		case input := <-keyin:
			r.total++
			if input == word {
				fmt.Print("Correct!\n\n")
				r.corrects++
			} else {
				fmt.Print("Missed!\n\n")
			}
		case <-timeup:
			fmt.Print("Time UP!\n\n")
			break LOOP
		}
	}

	fmt.Print("Typing Result\n\n")
	fmt.Printf("- SCORE: %d points (total %d words)\n", r.corrects, r.total)
	fmt.Printf("- Second : %d\n", t.TimeLimit/time.Second)

	return nil
}
