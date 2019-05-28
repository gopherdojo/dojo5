package main

import (
	"os"

	"../cmd/dirconv"
)

func main() {
	dirconv.NewCli(os.Args).DirConv()
}
