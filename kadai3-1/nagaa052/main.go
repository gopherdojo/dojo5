package main

import (
	"flag"
	"log"
	"os"

	"github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/game"
)

func main() {
	var timeout int
	var color bool
	flag.IntVar(&timeout, "t", game.DefaultOptions.TimeUpSecond, "Timeout Seconds")
	flag.BoolVar(&color, "c", game.DefaultOptions.IsColor, "Print Color")
	flag.Parse()

	g, err := game.New(game.Options{
		TimeUpSecond: timeout,
		IsColor:      color,
	}, os.Stdin, os.Stdout, os.Stderr)

	if err != nil {
		log.Fatal("Failed to start the game")
	}

	os.Exit(g.Start())
}
