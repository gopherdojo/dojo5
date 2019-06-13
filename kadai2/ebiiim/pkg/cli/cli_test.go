package cli_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/cli"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dirconv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

// TODO: table driven test
//func TestIsInvalidArgs(t *testing.T) {
//	e1 := cli.invalidArgsException
//	e2 := errors.New("")
//	if !cli.IsInvalidArgs(e1) || cli.IsInvalidArgs(e2) {
//		t.Errorf("error %v %v", e1, e2)
//	}
//}

func TestParseArgs(t *testing.T) {
	cases := []struct {
		name          string
		args          []string
		parsed        *dirconv.DirConv
		isErr         bool
		isInvalidArgs bool
	}{
		{name: "no_options",
			args:   []string{"imgconv", "../testdata"},
			parsed: &dirconv.DirConv{Dir: "../testdata", SrcExt: img.JPEG, TgtExt: img.PNG},
			isErr:  false, isInvalidArgs: false},
		{name: "with_options",
			args:   []string{"imgconv", "-source_ext=PNG", "-target_ext=.tiff", "../testdata"},
			parsed: &dirconv.DirConv{Dir: "../testdata", SrcExt: img.PNG, TgtExt: img.TIFF},
			isErr:  false, isInvalidArgs: false},
		{name: "with_options_src",
			args:   []string{"imgconv", "-source_ext=tif", "../testdata"},
			parsed: &dirconv.DirConv{Dir: "../testdata", SrcExt: img.TIFF, TgtExt: img.PNG},
			isErr:  false, isInvalidArgs: false},
		{name: "with_options_tgt",
			args:   []string{"imgconv", "-target_ext=bmp", "../testdata"},
			parsed: &dirconv.DirConv{Dir: "../testdata", SrcExt: img.JPEG, TgtExt: img.BMP},
			isErr:  false, isInvalidArgs: false},
		{name: "usage",
			args:   []string{"imgconv"},
			parsed: nil,
			isErr:  true, isInvalidArgs: true},
		{name: "no_dir",
			args:   []string{"imgconv", "-source_ext=PNG", "-target_ext=.tiff"},
			parsed: nil,
			isErr:  true, isInvalidArgs: true},
		{name: "invalid_options",
			args:   []string{"imgconv", "-source_ext=txt", "../testdata"},
			parsed: nil,
			isErr:  true, isInvalidArgs: true},
	}
	for _, c := range cases {
		c := c
		dc, err := cli.ParseArgs(c.args)
		// verify errors
		if !((err != nil) == c.isErr) || cli.IsInvalidArgs(err) != c.isInvalidArgs {
			t.Errorf("input %s, want %v(isErr) %v(isInvalidArgs), got %v", c.args, c.isErr, c.isInvalidArgs, err)
		}
		// verify DirConv
		if !cmp.Equal(dc, c.parsed) {
			t.Errorf("input %s, want %v, got %v", c.args, c.parsed, dc)
		}
	}
}
