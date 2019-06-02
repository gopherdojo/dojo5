package main

import (
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/dojo5/kadai1/smockoro/convert"
	"github.com/dojo5/kadai1/smockoro/walk"
)

var (
	fromExt = flag.String("from", "jpg", "Extension before conversion")
	toExt   = flag.String("to", "png", "Extension after conversion")
)

func main() {
	flag.Parse()
	fmt.Println(*fromExt, " ", *toExt)
	dirPath := flag.Arg(0)
	if dirPath == "" {
		log.Fatalln("Please input dir path")
	}

	// Walk
	w := walk.NewWalker()
	paths, err := w.Find(dirPath, *fromExt)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(paths)

	// Convert
	var c convert.Converter
	switch *toExt {
	case "png":
		c = convert.NewPngConverter()
	case "jpg":
		c = convert.NewJpgConverter()
	case "gif":
		c = convert.NewGifConverter()
	}

	for _, path := range paths {
		c.ImageConvert(path)
	}

}
