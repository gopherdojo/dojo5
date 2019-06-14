package target

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name    string
	root    string
	in      string
	targets []string
}{
	{
		name: "get jpg files",
		root: "../testdata/target",
		in:   "jpg",
		targets: []string{
			"../testdata/target/jpg/sample1.jpg",
			"../testdata/target/jpg/sample2.jpg",
			"../testdata/target/jpg/sample3.jpg",
		},
	},
	{
		name: "get png files",
		root: "../testdata/target",
		in:   "png",
		targets: []string{
			"../testdata/target/png/sample1.png",
			"../testdata/target/png/sample2.png",
			"../testdata/target/png/sample3.png",
		},
	},
	{
		name: "get gif files",
		root: "../testdata/target",
		in:   "gif",
		targets: []string{
			"../testdata/target/gif/sample1.gif",
			"../testdata/target/gif/sample2.gif",
			"../testdata/target/gif/sample3.gif",
		},
	},
	{
		name: "get tif files",
		root: "../testdata/target",
		in:   "tif",
		targets: []string{
			"../testdata/target/tif/sample1.tif",
			"../testdata/target/tif/sample2.tif",
			"../testdata/target/tif/sample3.tif",
		},
	},
	{
		name: "get bmp files",
		root: "../testdata/target",
		in:   "bmp",
		targets: []string{
			"../testdata/target/bmp/sample1.bmp",
			"../testdata/target/bmp/sample2.bmp",
			"../testdata/target/bmp/sample3.bmp",
		},
	},
	{
		name:    "no file",
		root:    "../testdata/target",
		in:      "hoge",
		targets: []string{},
	},
}

func TestGet(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Get(test.root, test.in)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if !reflect.DeepEqual(got, test.targets) {
				t.Fatalf("got targets: %v, want targets: %v", got, test.targets)
			}
		})
	}
}
