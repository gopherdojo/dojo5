package imgconv

import (
	"os"
	"testing"
)

var tests = []struct {
	name    string
	path    string
	out     string
	outFile string
}{
	{
		name:    "jpg to png",
		path:    "../testdata/imgconv/sample.jpg",
		out:     "png",
		outFile: "../testdata/imgconv/sample.png",
	},
	{
		name:    "jpg to gif",
		path:    "../testdata/imgconv/sample.jpg",
		out:     "gif",
		outFile: "../testdata/imgconv/sample.gif",
	},
	{
		name:    "jpg to tif",
		path:    "../testdata/imgconv/sample.jpg",
		out:     "tif",
		outFile: "../testdata/imgconv/sample.tif",
	},
	{
		name:    "jpg to bmp",
		path:    "../testdata/imgconv/sample.jpg",
		out:     "bmp",
		outFile: "../testdata/imgconv/sample.bmp",
	},
}

func TestConvert(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			img, err := Decode(test.path)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}

			err = img.Encode(test.out)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}

			_, err = os.Stat(test.outFile)
			if err != nil {
				t.Fatalf("did not create %v", test.outFile)
			}

			err = os.Remove(test.outFile)
			if err != nil {
				t.Fatalf("cannot remove %v", test.outFile)
			}
		})
	}
}
