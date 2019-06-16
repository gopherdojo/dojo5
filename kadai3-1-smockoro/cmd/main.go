package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// stdin goroutine
	problemChan := make(chan string, 1)

	go func(problemChan chan<- string) {
		problem := "Apple" // sample problem
		problemChan <- problem
	}(problemChan)

	go func(problemChan <-chan string) {
		problem := <-problemChan
		fmt.Println("Typing Problem : ", problem)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())

		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Err:", err)
		}
	}(problemChan)

	// Timer go goroutine

	time.Sleep(10 * time.Second)
}
