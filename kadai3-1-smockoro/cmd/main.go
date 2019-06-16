package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var LIMIT = 5 * time.Second

func startup() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("タイピングゲームです\n制限時間%sで何個問題が解けるか競います\n間違えても制限時間リセットされますが、スコアは悪くなります\n開始する場合はYesを、止める場合はそれ以外をタイプしてください\n", LIMIT)
	for scanner.Scan() {
		if scanner.Text() == "Yes" {
			break
		} else {
			return false
		}
	}
	return true
}

func main() {
	problemChan := make(chan string, 1)
	okChan := make(chan struct{}, 1)
	nextChan := make(chan struct{}, 1)

	// startup
	// ゲーム内容を説明して同意させる
	// 同意しなければ、チャネルを閉じて終了する。
	if !startup() {
		close(okChan)
		close(problemChan)
		close(nextChan)
		return
	}

	nextChan <- struct{}{} // 一番最初の問題を出させる

	// Problem gorutine
	// 問題を問題リストから出し続ける
	// Timerから次の問題を出すように言われたら問題を
	// stdin goroutineに対して渡す。
	problemList := []string{"Apple", "Orange", "Pineapple", "Mango", "Coffee", "Tea", "Water", "Miso", "Soy Source"}
	go func(problemChan chan<- string) {
		for _, problem := range problemList {
			<-nextChan
			problemChan <- problem
		}
	}(problemChan)

	// stdin goroutine
	// 標準入力を受け取って、問題に対して正解かを判定する
	// 正解ならTimerに対して、正解を告げる
	go func(problemChan <-chan string) {
		for {
			problem := <-problemChan
			fmt.Println("Typing Problem : ", problem)
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				fmt.Println("Typed text : ", scanner.Text())
				if scanner.Text() == problem {
					fmt.Println("OK")
					okChan <- struct{}{}
					break
				} else {
					fmt.Println("No matched")
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Err:", err)
			}
		}
	}(problemChan)

	// Timer
	// 時間計測をし続ける
	// stdin goroutineから次の問題を出すように言われたら
	// problem goroutineに問題を出すように伝える
	// 時間切れになったらすべてのチャネルを止めて終了する。
	for {
		select {
		case ok := <-okChan:
			nextChan <- ok
		case <-time.After(LIMIT):
			fmt.Println("\n=========end Game==========")
			close(okChan)
			close(problemChan)
			close(nextChan)
			return
		}
	}
}
