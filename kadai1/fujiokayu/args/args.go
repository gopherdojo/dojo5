package args

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

//Args struct has parsed args.
type Args struct {
	DecodeType     string
	EncodeType     string
	RootFolderName []string
}

func usage() {
	_, err := io.WriteString(os.Stderr, usageString)
	if err != nil {
		log.Fatal(err)
	}
	flag.PrintDefaults()
}

const usageString = `Usage of myConverter: 
  # convert
  ./myConverter [-from ext] [-to ext] directory

  # example 
  ./myConverter -from png -to jpg testdir
  
  # args
`

// ParseArgs is the constructor of struct "args"
func ParseArgs() (*Args, error) {
	flag.Usage = usage
	arg1 := flag.String("from", "jpg", "original file type to convert")
	arg2 := flag.String("to", "png", "file type you want to convert")

	flag.Parse()

	// フォルダが指定されているかチェックする
	var err error
	folder := flag.Args()
	if len(folder) == 0 {
		err = fmt.Errorf("specify target directory")
	}

	newArgs := &Args{
		DecodeType:     *arg1,
		EncodeType:     *arg2,
		RootFolderName: flag.Args(),
	}

	return newArgs, err
}
