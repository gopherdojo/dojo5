package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo5/kadai1/shuntakeuch1/ifconv"
)

// flagをつけてbefore format と after formatのデフォルトをきめる
var before = flag.String("b", "jpg", "変換前画像形式を指定してください")
var after = flag.String("a", "png", "変換後画像形式を指定してください")

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	// ディレクトリを聞く or 引数にないと強制終了
	if dir == "" {
		fmt.Println("ディレクトリを指定してください")
		os.Exit(0)
	}

	ifconv.Execute(dir, *before, *after)
}
