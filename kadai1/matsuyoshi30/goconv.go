package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/matsuyoshi30/dojo5/kadai1/matsuyoshi30/conv"
)

// usage
const usage = `
NAME:
   goconv - image format converter written in Go

USAGE:
   goconv [-b before image format] [-a after image format] path/to/dir

VERSION:
   0.1.0

GLOBAL OPTIONS:
   -b              specify format before converted (jpg, png, gif) [DEFAULT: jpg]
   -a              specify format after converted  (png, jpg, gif) [DEFAULT: png]
   --help, -h      show help
`

func main() {
	// 引数の長さを確認する
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	// フラグを読み込む
	var a, b string
	var sh bool
	flag.StringVar(&b, "b", "jpg", "Choose format before converted")
	flag.StringVar(&a, "a", "png", "Choose format after converted")
	flag.BoolVar(&sh, "h", false, "Show help")
	flag.Parse()

	// ヘルプフラグがオンの場合、ヘルプを表示して終了する
	if sh {
		fmt.Println(usage)
		return
	}

	// フラグで指定されたフォーマットを ImageType 型に置き換える
	bf := selectFormat(b)
	af := selectFormat(a)

	dirlist := OptPath(flag.Args())

	if len(dirlist) < 1 {
		fmt.Println(usage)
	} else {
		for _, d := range dirlist {
			err := conv.Imgconv(bf, af, d)
			logError(err, false)
		}
	}
}

// selectFormat はフラグに指定されたフォーマットの文字列を見て ImageType 型を返す
func selectFormat(f string) conv.ImageType {
	switch f {
	case "jpeg":
		return conv.JPEG
	case "jpg":
		return conv.JPG
	case "png":
		return conv.PNG
	case "gif":
		return conv.GIF
	default:
		log.Fatal("Unknown format")
	}
	return ""
}

func OptPath(paths []string) []string {
	dirlist := make([]string, 0)
	for _, p := range paths {
		if !contains(dirlist, p) {
			dirlist = append(dirlist, p)
		}
	}

	return dirlist
}

func contains(s []string, e string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if strings.HasPrefix(v, e) || strings.HasPrefix(e, v) {
			return true
		}
	}
	return false
}

func logError(err error, stop bool) {
	if err != nil {
		log.Fatal(err)
		if stop {
			os.Exit(1)
		}
	}
}
