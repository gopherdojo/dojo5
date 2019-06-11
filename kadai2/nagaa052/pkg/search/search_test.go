package search_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/nagaa052/pkg/search"
)

func TestWalk(t *testing.T) {
	t.Parallel()

	errMessage := `
	The quantity found is different.
		expected: %d
		actual: %d
	`
	cases := []struct {
		name     string
		ext      []string
		expected int
	}{
		{name: "search jpg", ext: []string{".jpg"}, expected: 4},
		{name: "search png", ext: []string{".png"}, expected: 2},
		{name: "search jpg and png", ext: []string{".jpg", ".png"}, expected: 6},
	}

	path := filepath.Join("testdata", t.Name())
	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			actual := 0
			search.WalkWithExtHandle(path, c.ext, func(path string, info os.FileInfo, err error) {
				actual++
			})
			if c.expected != actual {
				t.Error(fmt.Sprintf(errMessage, c.expected, actual))
			}
		})
	}
}
