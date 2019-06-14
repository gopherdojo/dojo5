package img_test

import (
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

func TestParseExt(t *testing.T) {
	cases := []struct {
		name  string
		input string
		ext   img.Ext
		isErr bool
	}{
		{name: "jpg", input: "abc.jpg", ext: img.JPEG, isErr: false},
		{name: "jpeg", input: "abc.jpeg", ext: img.JPEG, isErr: false},
		{name: "png", input: "abc.png", ext: img.PNG, isErr: false},
		{name: "tif", input: "abc.tif", ext: img.TIFF, isErr: false},
		{name: "tiff", input: "abc.tiff", ext: img.TIFF, isErr: false},
		{name: "bmp", input: "abc.bmp", ext: img.BMP, isErr: false},
		{name: "long_name", input: "a/b/c/d/e.png.jpg", ext: img.JPEG, isErr: false},
		{name: "Windows", input: `C:\Users\hoge\a.jpg`, ext: img.JPEG, isErr: false},
		{name: "weird_name", input: "abc..........jpg", ext: img.JPEG, isErr: false},
		{name: "F_undef_gif", input: "abc.gif", ext: "", isErr: true},
		{name: "F_undef_txt", input: "abc.txt", ext: "", isErr: true},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			ext, err := img.ParseExt(c.input)
			if ext != c.ext || !((err != nil) == c.isErr) {
				t.Errorf("input %s, want %v %v(isErr), got %v %v", c.input, c.ext, c.isErr, ext, err)
			}
		})
	}
}
