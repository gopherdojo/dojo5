package words_test

import (
	"testing"

	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/words"
)

func TestImport(t *testing.T) {
	cases := []struct {
		path      string
		firstWord string
		lastWord  string
		length    int
	}{
		{"../testdata/abc.txt", "a", "c", 3},
		{"../testdata/go_standard_library.txt", "archive", "unsafe", 154},
	}

	for _, c := range cases {
		c := c
		t.Run(c.path, func(t *testing.T) {
			t.Parallel()

			words, err := words.Import(c.path)
			if err != nil {
				t.Errorf("failed to Import %v: %v", c.path, err)
			}
			if words[0] != c.firstWord {
				t.Errorf("the first item of imported words is %v; it should be %v", words[0], c.firstWord)
			}
			if words[len(words)-1] != c.lastWord {
				t.Errorf("the first item of imported words is %v; it should be %v", words[0], c.lastWord)
			}
			if len(words) != c.length {
				t.Errorf("the length of imported words is %v; it should be %v", len(words), c.length)
			}
		})
	}
}
