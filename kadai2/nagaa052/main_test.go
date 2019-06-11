package main

import (
	"fmt"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/nagaa052/internal/gocon"
)

func TestParseOptions(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		expectedFrom string
		expectedTo   string
		isError      bool
	}{
		{
			name:         "Parse success",
			expectedFrom: "png",
			expectedTo:   "jpeg",
			isError:      false,
		},
		{
			name:         "The same format specification is an error",
			expectedFrom: "jpeg",
			expectedTo:   "jpeg",
			isError:      true,
		},
		{
			name:         "Format that does not exist is an error",
			expectedFrom: "png",
			expectedTo:   "jpeggg",
			isError:      true,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			imageFormat = fmt.Sprintf("%s:%s", c.expectedFrom, c.expectedTo)
			opt, err := parseOptions()

			if c.isError {
				if err == nil {
					t.Errorf("No error occurred. Test : %s", c.name)
				}
			} else {
				if err != nil {
					t.Errorf("No error occurred. Test : %s", c.name)
					return
				}

				testEqFormat(t, opt.FromFormat, gocon.ImgFormat(c.expectedFrom))
				testEqFormat(t, opt.ToFormat, gocon.ImgFormat(c.expectedTo))
			}
		})
	}
}

func testEqFormat(t *testing.T, actual, expected gocon.ImgFormat) {
	t.Helper()

	if actual != expected {
		t.Fatal(fmt.Sprintf(`
		Expected and actual are different.
			expected: %v
			actual: %v
		`, expected, actual))
	}
}
