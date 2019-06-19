package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/gopherdojo/dojo5/kadai3-2/manhdaovan/pkg/mget"
)

func main() {
	downloader := mget.MGet{
		WorkerNum: runtime.NumCPU(),
	}
	a, b := downloader.Download(context.Background(), "./tmp", "http://releases.ubuntu.com/16.04/ubuntu-16.04.6-desktop-i386.iso")
	fmt.Println("a ------ ", a)
	fmt.Println("b ------ ", b)
}
