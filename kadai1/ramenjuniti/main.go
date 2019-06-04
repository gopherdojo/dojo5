package main

import (
	"flag"
	"log"

	"github.com/gopherdojo/dojo5/kadai1/ramenjuniti/imgconv"
	"github.com/gopherdojo/dojo5/kadai1/ramenjuniti/target"
)

func main() {
	in := flag.String("in", "jpg", "file type before conversion")
	out := flag.String("out", "png", "file type to convert")
	formats := []string{"jpg", "png", "gif", "tif", "bmp"}

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("directory not specified")
	}

	if *in == "jpeg" {
		*in = "jpg"
	}
	if *in == "tiff" {
		*in = "tif"
	}
	if *out == "jpeg" {
		*out = "jpg"
	}
	if *out == "tiff" {
		*out = "tif"
	}

	if !canConv(*in, formats) {
		log.Fatalf("cannot convert %v file", *in)
	}
	if !canConv(*out, formats) {
		log.Fatalf("cannot convert %v file", *out)
	}

	for _, arg := range args {
		targets, err := target.Get(arg, *in)
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range targets {
			img, err := imgconv.Decode(t)
			if err != nil {
				log.Fatal(err)
			}

			err = img.Encode(*out)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// canConv tがconvに含まれるかどうかをboolで返す
func canConv(f string, formats []string) bool {
	for _, v := range formats {
		if f == v {
			return true
		}
	}
	return false
}
