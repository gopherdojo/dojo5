package gocon_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/nagaa052/internal/gocon"
)

func TestRun(t *testing.T) {
	t.Parallel()

	errMessage := `
	It was not converted properly.
		Error: %v
	`
	streamErrMessage := `
	Convert Error.
		%v
	`

	srcDir, err := filepath.Abs(filepath.Join("testdata", t.Name()))
	if err != nil {
		t.Error(err)
	}
	destDir, err := filepath.Abs("out")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		name     string
		options  gocon.Options
		expected int
		isError  bool
	}{
		{
			name: "Success Test",
			options: gocon.Options{
				FromFormat: gocon.JPEG,
				ToFormat:   gocon.PNG,
				DestDir:    destDir,
			},
			expected: gocon.ExitOK,
			isError:  false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
			g, err := gocon.New(srcDir, c.options, outStream, errStream)
			if !c.isError && err != nil {
				t.Error(fmt.Sprintf("gocon make error : %+v", err))
			}

			actual := g.Run()

			if _, err := os.Stat(destDir); err == nil {
				if err := os.RemoveAll(destDir); err != nil {
					fmt.Println(err)
				}
			}

			if c.expected != actual {
				t.Error(fmt.Sprintf(errMessage, c.expected, actual))
			}

			if errStream.String() != "" {
				t.Error(fmt.Sprintf(streamErrMessage, errStream.String()))
			}

			messageFmt := strings.Replace(gocon.SuccessConvertFileMessageFmt, "\n", "", -1)
			expectedMessage := fmt.Sprintf(messageFmt, destDir)
			if !strings.Contains(outStream.String(), expectedMessage) {
				t.Error("Conversion failed")
			}
		})
	}
}
