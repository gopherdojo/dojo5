package converter

import (
	"os"
	"path/filepath"
	"testing"
)

var defaultFile = []string{
	"testdata",
	"testdata/empty",
	"testdata/gif-image.gif",
	"testdata/gif-image1.gif",
	"testdata/jpeg-image.jpeg",
	"testdata/jpeg-image1.jpeg",
	"testdata/jpg-image.jpg",
	"testdata/jpg-image1.jpg",
	"testdata/png-image.png",
	"testdata/png-image1.png",
}

func TestConvertImg(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		cases := []struct {
			from            string
			to              string
			path            string
			expectedFiles   []string
			expectedOutputs []string
		}{
			{
				from:            ".jpeg",
				to:              ".png",
				path:            "./testdata",
				expectedFiles:   []string{"testdata/jpeg-image.jpeg", "testdata/jpeg-image1.jpeg"},
				expectedOutputs: []string{"testdata/jpeg-image.png", "testdata/jpeg-image1.png"},
			},
			{
				from:            ".jpeg",
				to:              ".gif",
				path:            "./testdata",
				expectedFiles:   []string{"testdata/jpeg-image.jpeg", "testdata/jpeg-image1.jpeg"},
				expectedOutputs: []string{"testdata/jpeg-image.gif", "testdata/jpeg-image1.gif"},
			},
			{
				from:            ".png",
				to:              ".jpeg",
				path:            "./testdata",
				expectedFiles:   []string{"testdata/png-image.png", "testdata/png-image1.png"},
				expectedOutputs: []string{"testdata/png-image.jpeg", "testdata/png-image1.jpeg"},
			},
			{
				from:            ".gif",
				to:              ".jpeg",
				path:            "./testdata",
				expectedFiles:   []string{"testdata/gif-image.gif", "testdata/gif-image1.gif"},
				expectedOutputs: []string{"testdata/gif-image.jpeg", "testdata/gif-image1.jpeg"},
			},
		}

		for i, tc := range cases {
			files, err := Convert(tc.from, tc.to, tc.path)

			if err != nil {
				t.Fatalf("#%d converter.Convert returned enexpected error: %s", i, err)
			}

			if got, want := len(files), len(tc.expectedFiles); got != want {
				t.Errorf("#%d converter.Convert returnede unexpected number of files: want: %d, got: %d", i, want, got)
			}

			for _, file := range files {

				ok := false
				for _, expected := range tc.expectedFiles {
					if file == expected {
						ok = true
					}
				}

				if !ok {
					t.Errorf("#%d converter.Convert returned unexpected file: %s", i, file)
				}
			}

			for _, expected := range tc.expectedOutputs {
				if _, err := os.Stat(expected); err != nil {
					t.Errorf("#%d converter.Convert did not output expected file: %s", i, expected)
				}
			}

			remove(t)
		}
	})

	t.Run("abnormal", func(t *testing.T) {
		cases := []struct {
			from    string
			to      string
			path    string
			managed bool
		}{
			{from: ".jpeg", to: ".png", path: "./unknown", managed: true},
			{from: ".jpeg", to: ".png", path: "./testdata/empty", managed: true},
			{from: ".svg", to: ".png", path: "./testdata", managed: true},
			{from: ".jpeg", to: ".png", path: "./unknown.jpeg", managed: true},
		}

		for i, ca := range cases {
			_, err := Convert(ca.from, ca.to, ca.path)

			if err == nil {
				t.Fatalf("#%d error is not supposed to be nil", i)
			}

			if _, ok := err.(Managed); ok != ca.managed {
				t.Fatalf("#%d managed is suppoed to be %v. error: %s", i, ca.managed, err)
			}
		}

	})
}

func remove(t *testing.T) {
	t.Helper()

	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		deletable := true
		for _, file := range defaultFile {
			if path == file {
				deletable = false
			}
		}

		if deletable {
			return os.Remove(path)
		}

		return nil
	})

	if err != nil {
		t.Errorf("failed to cleanup testdata")
	}
}
