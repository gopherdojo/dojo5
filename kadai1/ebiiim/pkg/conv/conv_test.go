package conv_test

import (
	"os"
	"testing"

	"github.com/gopherdojo/dojo5/kadai1/ebiiim/pkg/conv"
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
	ic = conv.ImgConv{SrcPath: f1, TgtPath: f2, Options: nil}
	err = ic.Convert()
	if err != nil {
		t.Errorf(e1)
	}

	// normal: png -> bmp
	ic = conv.ImgConv{SrcPath: f2, TgtPath: f3, Options: nil}
	err = ic.Convert()
	if err != nil {
		t.Errorf(e1)
	}

	// normal: bmp -> tiff
	ic = conv.ImgConv{SrcPath: f3, TgtPath: f4, Options: nil}
	err = ic.Convert()
	if err != nil {
		t.Errorf(e1)
	}

	// non-normal: jpg -> png
	ic = conv.ImgConv{SrcPath: n1, TgtPath: n2, Options: nil}
	err = ic.Convert()
	if err == nil {
		t.Errorf(e1)
	}
}
