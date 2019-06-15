package word

import (
	"bufio"
	"errors"
	"os"
)

func GenerateSource(file string) ([]string, error) {
	var words []string

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if len(words) == 0 {
		return nil, errors.New("no word for typing")
	}
	return words, nil
}
