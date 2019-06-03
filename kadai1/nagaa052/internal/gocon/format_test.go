package gocon_test

import (
	"fmt"
	"testing"

	"github.com/gopherdojo/dojo5/kadai1/nagaa052/internal/gocon"
)

func TestGetExtentions(t *testing.T) {
	errMessage := `
	Incorrect extension.
		expected: %s
		actual: %s
	`

	cases := []struct {
		name     string
		format   string
		expected []string
		isError  bool
	}{
		{
			name:     "success test",
			format:   "jpeg",
			expected: []string{".jpeg", ".jpg"},
			isError:  false,
		},
		{
			name:     "format does not exist",
			format:   "hogehoge",
			expected: []string{".hogehoge"},
			isError:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			i := gocon.ImgFormat(c.format)
			actual, err := i.GetExtentions()
			if c.isError {
				if err == nil {
					t.Error(err)
				}
			} else {
				if err != nil {
					t.Error(err)
				}

				if len(actual) <= 0 {
					t.Error(fmt.Sprintf(errMessage, c.expected, actual))
				}

				if c.expected[0] != actual[0] {
					t.Error(fmt.Sprintf(errMessage, c.expected, actual))
				}
			}
		})
	}
}
