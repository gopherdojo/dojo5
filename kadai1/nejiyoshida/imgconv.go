package main

import (
	"flag"
	"fmt"

	"github.com/gopherdojo/dojo5/kadai1/nejiyoshida/dircrawler"
	"github.com/gopherdojo/dojo5/kadai1/nejiyoshida/imgconverter"
)

func main() {

	var (
		s string
		t string
	)

	flag.StringVar(&s, "s", ".jpg", "format of image before conversion")
	flag.StringVar(&t, "t", ".png", "format of image after conversion")
	flag.Parse()

	args := flag.Args()
	searchDir, saveDir := args[0], args[1]
	files := dircrawler.SearchSpecificFormatFiles(searchDir, s)
	c := imgconverter.New(t, saveDir, files)
	c.Convert()
	fmt.Println("finished")
}
