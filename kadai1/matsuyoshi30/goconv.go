package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	flag.StringVar(&b, "b", "jpeg", "Choose format before converted")
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

	// TODO: 複数のディレクトリが指定されている場合は複数のディレクトリをリスト化する
	//       指定されている複数のディレクトリが階層的に重複している場合は、最上位のパスで最適化する
	dirlist := flag.Args()

	if len(dirlist) < 1 {
		fmt.Println(usage)
	} else {
		for _, d := range dirlist {
			conv.Imgconv(bf, af, d)
		}
	}
}

// selectFormat はフラグに指定されたフォーマットの文字列を見て ImageType 型を返す
func selectFormat(f string) conv.ImageType {
	switch f {
	case "jpeg":
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
