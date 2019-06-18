package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"./convert"
)

// e.g.) chimgs dir [-i=jpg/png] [-o=png/jpg]
func main() {

	var pix convert.PixConv

	pix.Src = flag.String("i", "jpeg", "string flag")
	pix.Dest = flag.String("o", "png", "string flag")
	flag.Parse()

	if (flag.NArg() == 0) && (flag.NFlag() == 0) {
		fmt.Println("Usage: chimgs DIR [-i=imgext] [-o=imgext]")
		os.Exit(1)
	}
	dirname := flag.Arg(0)

	filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			pix.Path = path
			err = convert.PixFile(data)
			return err
		})
}
