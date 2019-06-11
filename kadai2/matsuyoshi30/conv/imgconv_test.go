package conv

import (
	"os"
	"path/filepath"
	"testing"
)

const TESTDIR = "../testdata/"

func clean(filename string) {
	os.Remove(filename)
}

func TestImgconv(t *testing.T) {
	// SUCCESS CASE
	t.Run("goconv pass", func(t *testing.T) {
		testcases := []struct {
			input  string
			output string
		}{
			{"appenginegophercolor.jpg", "appenginegophercolor.png"},
			{"appenginelogo.gif", "appenginelogo.jpeg"},
			{"bumper.png", "bumper.gif"},
		}

		for _, tc := range testcases {
			testImgconv_pass(t, JPEG, PNG, filepath.Join(TESTDIR, tc.input), filepath.Join(TESTDIR, tc.output))
		}
	})

	// FAIL CASE
	t.Run("goconv fail", func(t *testing.T) {
		testImgconv_fail(t, JPEG, PNG, filepath.Join(TESTDIR, "dummy.jpeg"))
	})
}

func testImgconv_pass(t *testing.T, from, to ImageType, input, output string) {
	t.Helper()
	if err := Imgconv(from, to, input); err != nil {
		t.Fatalf("input: %v (%v -> %v): %v", input, from, to, err)
	}

	clean(output)
}

func testImgconv_fail(t *testing.T, from, to ImageType, input string) {
	t.Helper()
	if err := Imgconv(from, to, input); err == nil {
		t.Fatal("Expected error")
	}
}
