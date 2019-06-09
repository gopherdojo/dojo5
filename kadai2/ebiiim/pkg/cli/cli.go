// Package cli provides a command line parser to generate new DirConv instances.
package cli

import (
	"flag"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dirconv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
	"github.com/pkg/errors"
)

const usageSrcExt = `source extension (jpg, png, tiff, bmp)`
const usageTgtExt = `target extension (jpg, png, tiff, bmp)`

// Usage string
const Usage = `Usage:
  imgconv [-source_ext=<ext>] [-target_ext=<ext>] DIR
Arguments:
  -source_ext=<ext>` + "\t" + usageSrcExt + ` [default: jpg]
  -target_ext=<ext>` + "\t" + usageTgtExt + ` [default: png]`

// InvalidArgsException is used when the program do not have args.
type InvalidArgsException struct{ s string }

// Error returns a InvalidArgsException error message.
func (e *InvalidArgsException) Error() string { return e.s }

// NoArgs is used by IsInvalidArgs to check the interface.
func (e *InvalidArgsException) NoArgs() bool { return true }

// IsInvalidArgs checks the given error implements InvalidArgsException or do not.
func IsInvalidArgs(err error) bool {
	b, ok := errors.Cause(err).(interface{ NoArgs() bool })
	return ok && b.NoArgs()
}

// ParseArgs initializes a DirConv struct with given command line arguments.
func ParseArgs(args []string) (*dirconv.DirConv, error) {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		argSrcExt = flags.String("source_ext", "jpg", usageSrcExt)
		argTgtExt = flags.String("target_ext", "png", usageTgtExt)
	)
	err := flags.Parse(args[1:])
	if err != nil {
		return &dirconv.DirConv{}, err
	}

	dir := flags.Arg(0) // get the first dir name only
	if len(dir) == 0 {
		return &dirconv.DirConv{}, &InvalidArgsException{"no directory specified"}
	}

	srcExt, err := img.ParseExt(*argSrcExt)
	if err != nil {
		return &dirconv.DirConv{}, &InvalidArgsException{"invalid source extension"}
	}
	tgtExt, err := img.ParseExt(*argTgtExt)
	if err != nil {
		return &dirconv.DirConv{}, &InvalidArgsException{"invalid target extension"}
	}

	return &dirconv.DirConv{Dir: dir, SrcExt: srcExt, TgtExt: tgtExt}, nil
}
