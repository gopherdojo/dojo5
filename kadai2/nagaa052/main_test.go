package main

import (
	"fmt"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/nagaa052/internal/gocon"
)

func TestParseOptions(t *testing.T) {

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

	errMessage := `
	Expected and actual are different.
		expected: %v
		actual: %v
	`

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			imageFormat = fmt.Sprintf("%s:%s", c.expectedFrom, c.expectedTo)
			opt, err := parseOptions()

			if c.isError {
				if err == nil {
					t.Errorf("No error occurred. Test : %s", c.name)
				}
			} else {
				if opt.FromFormat != gocon.ImgFormat(c.expectedFrom) {
					t.Error(fmt.Sprintf(errMessage, gocon.ImgFormat(c.expectedFrom), opt.FromFormat))
				}
				if opt.ToFormat != gocon.ImgFormat(c.expectedTo) {
					t.Error(fmt.Sprintf(errMessage, gocon.ImgFormat(c.expectedTo), opt.ToFormat))
				}
			}
		})
	}
}
