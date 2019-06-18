package words

import (
	"bufio"
	"os"

	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/typinggame"
	"github.com/hashicorp/go-multierror"
)

// Import reads the text file and returns the words for the typing game.
func Import(path string) (typinggame.Words, error) {
	var words typinggame.Words
	var result error

	f, err := os.Open(path)
	if err != nil {
		result = multierror.Append(result, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			result = multierror.Append(result, err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
	if err := s.Err(); err != nil {
		result = multierror.Append(result, err)
	}

	return words, result
}
