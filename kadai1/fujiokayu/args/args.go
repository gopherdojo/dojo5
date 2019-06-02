package args

import (
	"flag"
	"log"
)

//Args struct has parsed args.
type Args struct {
	DecodeType     string
	EncodeType     string
	RootFolderName []string
}

// ParseArgs is the constructor of struct "args"
func ParseArgs() *Args {
	arg1 := flag.String("from", "jpg", "original file type to convert")
	arg2 := flag.String("to", "png", "file type you want to convert")

	flag.Parse()

	// フォルダが指定されているかチェックする
	folder := flag.Args()
	if len(folder) == 0 {
		log.Fatal("specify target directory")
	}

	newArgs := &Args{
		DecodeType:     *arg1,
		EncodeType:     *arg2,
		RootFolderName: flag.Args(),
	}

	return newArgs
}
