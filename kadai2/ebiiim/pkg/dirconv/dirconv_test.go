package dirconv_test

import (
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dirconv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

func TestCli_DirConv(t *testing.T) {
	dc := dirconv.DirConv{Dir: "../testdata", SrcExt: img.JPEG, TgtExt: img.PNG}
	got := dc.Convert()
	want := []dirconv.Result{
		{Index: 0, RelPath: "dummy.jpg", IsOk: false},
		{Index: 1, RelPath: "gopherA.jpg", IsOk: true},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want %v", got, want)
	}
}
