package game

import (
	"fmt"
	"io"
)

type colorCode int

const (
	red    colorCode = 31
	green  colorCode = 32
	yellow colorCode = 33
	blue   colorCode = 34
)

func cFPrint(code colorCode, w io.Writer, text string) {
	fmt.Fprintf(w, "\x1b[%dm%s\x1b[0m\n", code, text)
}
