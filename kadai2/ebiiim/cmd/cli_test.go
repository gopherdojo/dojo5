package main_test

import (
	"testing"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/cmd"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/conv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dirconv"
)

func check(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}

func TestParseArgs(t *testing.T) {
	// normal: no options
	args := []string{"imgconv", "../../testdata"}
	dc := main.ParseArgs(args)
	check(t, *dc, dirconv.DirConv{Dir: "../../testdata", SrcExt: conv.ImgExtJPEG, TgtExt: conv.ImgExtPNG})
}

func TestParseArgs2(t *testing.T) {
	// normal: with options
	args := []string{"imgconv", "-source_ext=PNG", "-target_ext=.tiff", "../../testdata"}
	dc := main.ParseArgs(args)
	check(t, *dc, dirconv.DirConv{Dir: "../../testdata", SrcExt: conv.ImgExtPNG, TgtExt: conv.ImgExtTIFF})
}
