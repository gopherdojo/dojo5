package main

import (
	"fmt"
	"os"

	godl "github.com/matsuyoshi30/dojo5/kadai3-2/matsuyoshi30"
)

func main() {
	g := godl.NewGodl()
	if err := g.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
	}
}
