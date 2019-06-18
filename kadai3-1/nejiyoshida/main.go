package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var random_strs = []string{"archive", "bufio", "builtin", "bytes", "cmd", "compress", "container", "context",
	"crypto", "database", "debug", "encoding", "errors", "flag", "fmt", "index", "internal",
	"io", "log", "math", "mime", "net", "os", "path", "plugin", "reflect", "regexp", "runtime",
	"sort", "strconv", "strings", "sync", "syscall", "testing", "text", "time", "unicode", "unsafe", "vendor"}

func main() {
	rand.Seed(time.Now().UnixNano())
	correct := 0
	incorrect := 0
	var t = 20 * time.Second
	//empty context キャンセルされず値も持たず、タイムアウトもない
	bc := context.Background()

	//タイムアウト時間を設定。cancelfuncを呼ぶとctxとctxの子のcontextがキャンセルされる
	ctx, cancelfunc := context.WithTimeout(bc, t)
	defer cancelfunc()

	//標準入力を受け続け、chに送信する
	ch := inputstr(os.Stdin, ctx)

GAME:
	for {
		random_num := rand.Intn(len(random_strs))
		ans := random_strs[random_num]
		fmt.Print("word : ")
		fmt.Println(ans)

		select {

		case <-ctx.Done():
			fmt.Println("game finished")
			break GAME

		default:
			typed := <-ch
			if typed == ans {
				fmt.Println("correct!!!")
				correct++
			} else if typed == "" {
				//時間切れの時にゼロ値が送られてくるため
				continue
			} else {
				incorrect++
			}

		}
	}
	fmt.Printf("corecct: %d, incorrect: %d\n", correct, incorrect)
}

func inputstr(input io.Reader, ctx context.Context) chan string {

	ch := make(chan string)
	go func() {
		str := bufio.NewScanner(input)

		// 標準入力を読み続けるのでこのゴルーチンがずっと続く
		for str.Scan() {
			select {
			//Doneを受け取るというか、値が入っている場合チャネルを閉じてゲーム終了。このゴルーチンも閉じる
			case <-ctx.Done():
				close(ch)
			//そうでなければ入力をchに送る
			default:
				ch <- str.Text()
			}

		}
	}()
	//inputstr()のメインルーチンはchを戻す
	return ch
}
