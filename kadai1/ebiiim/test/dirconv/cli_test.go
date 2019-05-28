package dirconv

import (
	"testing"

	"../../cmd/dirconv"
)

func check(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}

func TestNewCli1(t *testing.T) {
	// normal: no options
	args := []string{"imgconv", "../../testdata"}
	cli := dirconv.NewCli(args)
	check(t, *cli, dirconv.Cli{Dir: "../../testdata", SrcExt: ".jpg", TgtExt: ".png"})
}

func TestNewCli2(t *testing.T) {
	// normal: with options
	args := []string{"imgconv", "-source_ext=PNG", "-target_ext=.tiff", "../../testdata"}
	cli := dirconv.NewCli(args)
	check(t, *cli, dirconv.Cli{Dir: "../../testdata", SrcExt: ".png", TgtExt: ".tiff"})
}

//func TestCli_DirConv(t *testing.T) {
//	args := []string{"imgconv", "../../testdata"}
//	dirconv.NewCli(args).DirConv() // FIXME: nil pointer dereference
//}
