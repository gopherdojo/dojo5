package typinggame

import (
	"fmt"
	"time"
)

type Words []string

type Game struct {
	Words     Words
	TimeLimit time.Duration
}

func Execute(g Game) error {
	fmt.Println(g.Words)
	fmt.Println(g.TimeLimit)
	return nil
}
