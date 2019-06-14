package dir_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dir"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

func TestTraverseImageFiles(t *testing.T) {
	cases := []struct {
		name  string
		dir   string
		ext   img.Ext
		files []string
		isErr bool
	}{
		// TODO: more cases
		{name: "jpg", dir: "../testdata", ext: img.JPEG, files: []string{"dummy.jpg", "gopherA.jpg"}, isErr: false},
	}
	for _, c := range cases {
		c := c
		dirs, err := dir.TraverseImageFiles(c.dir, c.ext)
		if !cmp.Equal(dirs, c.files) || !((err != nil) == c.isErr) {
			t.Errorf("input %s, want %v %v(isErr), got %v %v", c.dir, c.files, c.isErr, dirs, err)
		}
	}
}
