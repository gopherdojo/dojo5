package conv_test

import (
	"os"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/conv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

func TestImgConv_Convert(t *testing.T) {
	const (
		f1 = "../testdata/gopherA.jpg"
		f2 = "../testdata/gopherA.png"
		f3 = "../testdata/gopherA.bmp"
		f4 = "../testdata/gopherA.tiff"

		n1 = "../testdata/dummy.jpg"
		n2 = "../testdata/dummy.png"

		e1 = "unexpected behavior"
	)

	defer func() {
		os.Remove(f2)
		os.Remove(f3)
		os.Remove(f4)
	}()

	var (
		ic  conv.ImgConv
		err error
	)

	// normal: jpg -> png
	ic = conv.ImgConv{SrcPath: f1, SrcExt: img.JPEG, TgtPath: f2, TgtExt: img.PNG, Options: nil}
	err = ic.Convert()
	if err != nil {
		t.Errorf(e1)
	}

	// normal: png -> bmp
	ic = conv.ImgConv{SrcPath: f2, SrcExt: img.PNG, TgtPath: f3, TgtExt: img.BMP, Options: nil}
	err = ic.Convert()
	if err != nil {
		t.Errorf(e1)
	}

	// normal: bmp -> tiff
	ic = conv.ImgConv{SrcPath: f3, SrcExt: img.BMP, TgtPath: f4, TgtExt: img.TIFF, Options: nil}
	err = ic.Convert()
	if err != nil {
		t.Errorf(e1)
	}

	// non-normal: jpg -> png
	ic = conv.ImgConv{SrcPath: n1, SrcExt: img.JPEG, TgtPath: n2, TgtExt: img.PNG, Options: nil}
	err = ic.Convert()
	if err == nil {
		t.Errorf(e1)
	}
}
