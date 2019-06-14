package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo5/kadai2/ramenjuniti/imgconv"
	"github.com/gopherdojo/dojo5/kadai2/ramenjuniti/target"
)

func main() {
	in := flag.String("in", "jpg", "file type before conversion")
	out := flag.String("out", "png", "file type to convert")
	formats := []string{"jpg", "png", "gif", "tif", "bmp"}

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "directory not specified")
		os.Exit(1)
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
		fmt.Fprintf(os.Stderr, "cannot convert %v file", *in)
		os.Exit(1)
	}
	if !canConv(*out, formats) {
		fmt.Fprintf(os.Stderr, "cannot convert %v file", *out)
		os.Exit(1)
	}

	for _, arg := range args {
		targets, err := target.Get(arg, *in)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		for _, t := range targets {
			img, err := imgconv.Decode(t)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			if err = img.Encode(*out); err != nil {
				fmt.Fprintln(os.Stderr, err)
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
