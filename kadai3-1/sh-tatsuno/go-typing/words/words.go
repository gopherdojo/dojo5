package words

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Import(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in os.Open: %v", err)
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in ioutil.ReadAll: %v", err)
		return nil, err
	}

	s := strings.Split(string(b), "\n")
	return s, nil
}
