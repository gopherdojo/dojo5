package main

import (
	"os"

	"github.com/gopherdojo/dojo5/kadai1/ebiiim/pkg/dirconv"
)

func main() {
	dirconv.NewCli(os.Args).DirConv()
}
