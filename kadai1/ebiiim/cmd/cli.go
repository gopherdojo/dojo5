package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo5/kadai1/ebiiim/pkg/conv"
	"github.com/gopherdojo/dojo5/kadai1/ebiiim/pkg/dirconv"
)

const usageSrcExt = `source extension (jpg, png, tiff, bmp)`
const usageTgtExt = `target extension (jpg, png, tiff, bmp)`
const usage = `Usage:
  imgconv [-source_ext=<ext>] [-target_ext=<ext>] DIR
Arguments:
  -source_ext=<ext>` + "\t" + usageSrcExt + ` [default: jpg]
  -target_ext=<ext>` + "\t" + usageTgtExt + ` [default: png]`

// ParseArgs initializes a DirConv struct with given command line arguments.
func ParseArgs(args []string) *dirconv.DirConv {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		srcExt = flags.String("source_ext", "jpg", usageSrcExt)
		tgtExt = flags.String("target_ext", "png", usageTgtExt)
	)
	err := flags.Parse(args[1:])
	if err != nil {
		panic(err)
	}
	dir := flags.Arg(0) // get the first dir name only

	return &dirconv.DirConv{Dir: dir, SrcExt: conv.ParseImgExt(*srcExt), TgtExt: conv.ParseImgExt(*tgtExt)}
}

func main() {
	dc := ParseArgs(os.Args)

	// show help if no dir specified
	if dc.Dir == "" {
		fmt.Println(usage)
		os.Exit(0)
	}
	dc.Convert()
}
