// Command bin/convert ...
package main

import (
	"flag"
	"fmt"

	"github.com/lfcd85/dojo5/kadai1/lfcd85/imgconv"
)

func main() {
	from := flag.String("f", "jpeg", "Image format before conversion (default: jpeg)")
	to := flag.String("t", "png", "Image format after conversion (default: png)")
	flag.Parse()
	dirName := flag.Arg(0)

	err := imgconv.Convert(dirName, *from, *to)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
