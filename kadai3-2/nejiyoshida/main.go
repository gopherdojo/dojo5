package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"./downloader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(1)
	}

	url := os.Args[1]

	tmp := strings.Split(url, "/")

	filename := tmp[len(tmp)-1]
	fp, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	defer fp.Close()

	size, dltype, err := downloader.CheckHead(url)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	//cancel時に他のゴルーチンもとじる
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := start(ctx, fp, url, dltype, size)

	sig := make(chan os.Signal)
	//ctrl+cを受けるようにする
	signal.Notify(sig, syscall.SIGINT)

loop:
	for {
		select {
		case err = <-res:
			// cancel()が実行されるかエラーが戻ってくるとループを抜ける
			break loop

		case <-sig:
			//ctrl+cなどで中断
			fmt.Fprintln(os.Stderr, "ctrl+c received")
			cancel()
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

}

func start(ctx context.Context, fp *os.File, url, dltype string, size int64) <-chan error {
	ch := make(chan error)

	go func() {
		defer close(ch)

		var err error

		switch dltype {
		case "bytes":
			err = downloader.ParallelDownload(ctx, fp, url, size)
		default:
			err = downloader.Download(ctx, fp, url, size)
		}

		ch <- err
	}()

	return ch
}
