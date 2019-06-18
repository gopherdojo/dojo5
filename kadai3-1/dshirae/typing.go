package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

// result
type resultList struct {
	questions int
	collects  int
}

// test slice
func makeQuestionsList(file string) ([]string, error) {
	var list []string

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		list = append(list, s.Text())
	}
	return list, nil
}

// main
func main() {
	qtime := flag.Int("t", 10, "set second")
	qfile := flag.String("f", "test.txt", "word list")
	flag.Parse()

	// test file open
	qlist, err := makeQuestionsList(*qfile)
	if err != nil {
		fmt.Printf("file not exist:%s", *qfile)
		return
	}

	// 結果を準備
	rslt := &resultList{}
	timeup := time.After(time.Duration(*qtime) * time.Second)
	keyin := make(chan string)

	// keyboard input
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			keyin <- s.Text()
		}
	}()

	fmt.Printf("TYPING start\n\n")

	// check
LOOP:
	for _, question := range qlist {
		// Disp question
		rslt.questions++
		fmt.Printf("-> %s\n", question)

		select {
		case in := <-keyin:
			if in == question {
				rslt.collects++
			}
		case <-timeup:
			fmt.Println("Time UP!\n")ß
			break LOOP
		}
	}

	fmt.Println("Typing Result\n")
	fmt.Printf("- TOTAL  : %d\n", rslt.questions)
	fmt.Printf("- COLLECT: %d\n", rslt.collects)
	fmt.Printf("- Second : %d\n", *qtime)
}
