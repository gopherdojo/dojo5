package main

import (
	"flag"
	"fmt"

	"imgconv"
	"option"
)

var (
	s = flag.String("s", "jpeg", "source image format. support:[jpeg, png, gif]")
	d = flag.String("d", "png", "destination image format. support:[jpeg, png, gif]")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("set the directory path in the argument.")
		return
	}

	err := option.Valid(*s, *d)
	if err != nil {
		fmt.Println(err)
		return
	}

	conv := imgconv.Converter{SrcExt: *s, DstExt: *d}
	err = conv.Convert(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("complete!")
}
