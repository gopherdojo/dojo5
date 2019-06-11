package convert_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/nagaa052/pkg/convert"
)

func TestToJpeg(t *testing.T) {
	t.Parallel()

	errMessage := `
	Expected and actual are different.
		expected: %v
		actual: %v
	`

	srcBase, err := filepath.Abs(filepath.Join("testdata", t.Name()))
	if err != nil {
		t.Error(err)
	}

	destDir := filepath.Join(srcBase, "out")

	cases := []struct {
		name     string
		srcPath  string
		expected string
		isError  bool
	}{
		{
			name:     "Success Test",
			srcPath:  filepath.Join(srcBase, "1.png"),
			expected: filepath.Join(destDir, "1.jpeg"),
			isError:  false,
		},
		{
			name:     "Not an image file",
			srcPath:  filepath.Join(srcBase, "2.png"),
			expected: "",
			isError:  true,
		},
		{
			name:     "Error if DestDir already exists",
			srcPath:  filepath.Join(srcBase, "1.png"),
			expected: "",
			isError:  true,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			con, err := convert.New(c.srcPath, destDir)
			if err != nil {
				t.Error(err)
			}

			actual, err := con.ToJpeg(&convert.JpegOptions{})
			if c.isError {
				if err == nil {
					t.Errorf("No error occurred. Test : %s", c.name)
				}
			} else if actual != c.expected {
				t.Error(fmt.Sprintf(errMessage, c.expected, actual))
			}
		})
	}

	if _, err := os.Stat(destDir); err == nil {
		if err := os.RemoveAll(destDir); err != nil {
			fmt.Println(err)
		}
	}
}

func TestToPng(t *testing.T) {
	t.Parallel()

	errMessage := `
	Expected and actual are different.
		expected: %v
		actual: %v
	`

	sourceBase, err := filepath.Abs(filepath.Join("testdata", t.Name()))
	if err != nil {
		t.Error(err)
	}

	outDir := filepath.Join(sourceBase, "out")

	cases := []struct {
		name       string
		sourcePath string
		expected   string
		isError    bool
	}{
		{
			name:       "success test",
			sourcePath: filepath.Join(sourceBase, "1.jpg"),
			expected:   filepath.Join(outDir, "1.png"),
			isError:    false,
		},
		{
			name:       "failed test",
			sourcePath: filepath.Join(sourceBase, "1.jpg"),
			expected:   "",
			isError:    true,
		},
	}
	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			con, err := convert.New(c.sourcePath, outDir)
			if err != nil {
				t.Error(err)
			}

			actual, err := con.ToPng()
			if c.isError && err == nil {
				t.Error(err)
			}

			if actual != c.expected {
				t.Error(fmt.Sprintf(errMessage, c.expected, actual))
			}
		})
	}

	if _, err := os.Stat(outDir); err == nil {
		if err := os.RemoveAll(outDir); err != nil {
			fmt.Println(err)
		}
	}
}
