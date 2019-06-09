package dirconv_test

import (
	"fmt"
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dirconv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

func TestDirConv_Convert(t *testing.T) {
	cases := []struct {
		name    string
		dir     string
		srcExt  img.Ext
		tgtExt  img.Ext
		results []*dirconv.Result
		isErr   bool
	}{
		// TODO: more cases
		{name: "jpg2png", dir: "../testdata", srcExt: img.JPEG, tgtExt: img.PNG,
			results: []*dirconv.Result{
				{Index: 0, RelPath: "dummy.jpg", Err: fmt.Errorf("")},
				{Index: 1, RelPath: "gopherA.jpg", Err: nil},
			}, isErr: false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			dc := &dirconv.DirConv{Dir: c.dir, SrcExt: c.srcExt, TgtExt: c.tgtExt}
			results, err := dc.Convert()
			// verify err
			if !((err != nil) == c.isErr) {
				t.Errorf("input %s, want %v(isErr), got %v", c.dir, c.isErr, err)
			}
			// verify results
			// TODO: sort results by Result.Index
			for i, r := range results {
				if r.Index != c.results[i].Index ||
					r.RelPath != c.results[i].RelPath ||
					(r.Err == nil) != (c.results[i].Err == nil) {
					t.Errorf("input %s, want %v, got %v", c.dir, *c.results[i], *r)
				}
			}
		})
	}
}
