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
	a, b := downloader.Download(context.Background(), "./tmp/", "https://itsycal.s3.amazonaws.com/Itsycal.zip")
	fmt.Println("a ------ ", a)
	fmt.Println("b ------ ", b)
}
