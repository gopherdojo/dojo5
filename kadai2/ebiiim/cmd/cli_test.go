package main

import (
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dirconv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

func check(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}

func TestParseArgs(t *testing.T) {
	// normal: no options
	args := []string{"imgconv", "../../testdata"}
	dc, _ := ParseArgs(args)
	check(t, *dc, dirconv.DirConv{Dir: "../../testdata", SrcExt: img.JPEG, TgtExt: img.PNG})
}

func TestParseArgs2(t *testing.T) {
	// normal: with options
	args := []string{"imgconv", "-source_ext=PNG", "-target_ext=.tiff", "../../testdata"}
	dc, _ := ParseArgs(args)
	check(t, *dc, dirconv.DirConv{Dir: "../../testdata", SrcExt: img.PNG, TgtExt: img.TIFF})
}
