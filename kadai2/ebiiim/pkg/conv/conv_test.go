package conv_test

import (
	"os"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/conv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

func removeFile(t *testing.T, path string) {
	t.Helper()
	err := os.Remove(path)
	if err != nil {
		t.Fatalf("failed to remove file %v %v", path, err)
	}
}

func TestImgConv_Convert(t *testing.T) {
	const (
		f1 = "../testdata/gopherA.jpg"
		f2 = "../testdata/gopherA.png"
		f3 = "../testdata/gopherA.bmp"
		f4 = "../testdata/gopherA.tif"

		n1 = "../testdata/dummy.jpg"
		n2 = "../testdata/dummy.png"
		n3 = "../testdata/dummy.tif"
	)
	defer func() {
		removeFile(t, f2)
		removeFile(t, f3)
		removeFile(t, f4)
	}()
	cases := []struct {
		name   string
		src    string
		srcExt img.Ext
		tgt    string
		tgtExt img.Ext
		opt    map[string]interface{}
		isErr  bool
	}{
		{name: "jpg2png", src: f1, srcExt: img.JPEG, tgt: f2, tgtExt: img.PNG, opt: nil, isErr: false},
		{name: "png2bmp", src: f2, srcExt: img.PNG, tgt: f3, tgtExt: img.BMP, opt: nil, isErr: false},
		{name: "bmp2tif", src: f3, srcExt: img.BMP, tgt: f4, tgtExt: img.TIFF, opt: nil, isErr: false},
		{name: "F_invalid_jpg2tiff", src: n1, srcExt: img.JPEG, tgt: n3, tgtExt: img.TIFF, opt: nil, isErr: true},
		{name: "F_invalid_png2tiff", src: n2, srcExt: img.PNG, tgt: n3, tgtExt: img.TIFF, opt: nil, isErr: true},
		{name: "F_not_found", src: "../testdata/notFound.jpg", srcExt: img.JPEG, tgt: "../testdata/notFound.tif", tgtExt: img.TIFF, opt: nil, isErr: true},
		{name: "F_src_is_dir", src: "../testdata", srcExt: img.JPEG, tgt: "../testdata.tif", tgtExt: img.TIFF, opt: nil, isErr: true},
		{name: "F_tgt_is_dir", src: f1, srcExt: img.JPEG, tgt: "../testdata", tgtExt: img.TIFF, opt: nil, isErr: true},
	}
	// NOTE: this test cannot be run in parallel
	for _, c := range cases {
		converter := &conv.ImgConv{SrcPath: c.src, SrcExt: c.srcExt, TgtPath: c.tgt, TgtExt: c.tgtExt, Options: c.opt}
		err := converter.Convert()
		if !((err != nil) == c.isErr) {
			t.Errorf("ImgConv %v, want %v(isErr), got %v", converter, c.isErr, err)
		}
		// verify file
		stat, err := os.Stat(c.tgt)
		if c.isErr {
			// WHEN c.isErr THEN (file does not exist) or (file is a dir)
			if !(err != nil || stat.IsDir()) {
				t.Errorf("failed delete %s", c.tgt)
			}
		} else {
			// WHEN !c.isErr THEN (file exists) and (file is not a dir)
			if !(err == nil && !stat.IsDir()) {
				t.Errorf("file does not found %s", c.tgt)
			}
		}
	}
}
