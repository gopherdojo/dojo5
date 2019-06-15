package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/matsuyoshi30/dojo5/kadai3-1/matsuyoshi30/parse"
	"github.com/matsuyoshi30/dojo5/kadai3-1/matsuyoshi30/word"
)

type Result struct {
	totalword   int
	collectword int
	totalchar   int
	time        int
}

func main() {
	o := parse.ParseFlag(os.Args[1:]...)
	fmt.Printf("<=== SIMPLE TYPING GAME ===>\n\n")

	gotp(o.LimitTime, o.WordFile)
}

func gotp(limit int, wordfile string) error {
	lt := time.After(time.Duration(limit) * time.Second)

	words, err := word.GenerateSource(wordfile)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())

	r := &Result{time: int(limit)}
	ch := inputString(os.Stdin)
LOOP:
	for {
		w := words[rand.Intn(len(words))]
		r.totalword++
		r.totalchar += len(w)

		fmt.Printf("WORD: %s\n> ", w)
		select {
		case input := <-ch:
			r.checkInput(input, w)
			fmt.Println()
		case <-lt:
			fmt.Println()
			break LOOP
		}

		time.Sleep(500 * time.Millisecond) // wait next question
	}

	r.printResult()

	return nil
}

func inputString(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()

	return ch
}

const (
	CollectColor   = "\033[1;34m%s\033[0m"
	IncollectColor = "\033[1;31m%s\033[0m"
)

func (r *Result) checkInput(input, word string) {
	col := IncollectColor
	res := "INCOLLECT.. "
	if input == word {
		r.collectword++
		col = CollectColor
		res = "COLLECT! "
	}

	fmt.Fprintf(os.Stdout, col, res+"=> "+input+"\n")
}

func (r *Result) printResult() {
	fmt.Println()

	fmt.Println("*** GAME FINISH! ***")
	fmt.Println("RESULT")
	fmt.Printf("  TOTAL: %d\n", r.totalword)
	fmt.Printf("  COLLECT: %d\n", r.collectword)
	fmt.Printf("  HIT per Second: %f\n", float64(r.totalchar)/float64(r.time))
}
