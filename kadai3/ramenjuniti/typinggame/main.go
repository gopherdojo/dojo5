package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	t := flag.Int("t", 10, "制限時間（秒）")
	flag.Parse()

	var score int
	l := time.After(time.Duration(*t) * time.Second)
	in := input(os.Stdin)

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Start!!")
L:
	for {
		q := words[rand.Intn(len(words))]
		fmt.Println(q)
		fmt.Print(">> ")
		select {
		case a := <-in:
			score += judge(q, a)
		case <-l:
			fmt.Println("finish!!")
			break L
		}
	}
	fmt.Printf("Score: %v, type/sec: %2.1f\n", score, float64(score)/float64(*t))
}

func judge(q, a string) int {
	var score int
	for i := 0; i < len(q); i++ {
		if i == len(a) {
			break
		}
		if q[i] == a[i] {
			score++
		}
	}
	return score
}

func input(r io.Reader) <-chan string {
	c := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			c <- s.Text()
		}
	}()
	return c
}
