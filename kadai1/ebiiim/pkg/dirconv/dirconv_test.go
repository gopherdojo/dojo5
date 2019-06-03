package dirconv_test

import (
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo5/kadai1/ebiiim/pkg/conv"
	"github.com/gopherdojo/dojo5/kadai1/ebiiim/pkg/dirconv"
)

func TestCli_DirConv(t *testing.T) {
	dc := dirconv.DirConv{Dir: "../testdata", SrcExt: conv.ImgExtJPEG, TgtExt: conv.ImgExtPNG}
	got := dc.Convert()
	want := []dirconv.Result{
		{Index: 0, RelPath: "dummy.jpg", IsOk: false},
		{Index: 1, RelPath: "gopherA.jpg", IsOk: true},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want %v", got, want)
	}
}
