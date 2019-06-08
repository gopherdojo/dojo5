package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/conv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dirconv"
)

const usageSrcExt = `source extension (jpg, png, tiff, bmp)`
const usageTgtExt = `target extension (jpg, png, tiff, bmp)`
const usage = `Usage:
  imgconv [-source_ext=<ext>] [-target_ext=<ext>] DIR
Arguments:
  -source_ext=<ext>` + "\t" + usageSrcExt + ` [default: jpg]
  -target_ext=<ext>` + "\t" + usageTgtExt + ` [default: png]`

// NoArgsException is used when the program do not have args.
type NoArgsException struct{ s string }

// Error returns a NoArgsException error message.
func (e *NoArgsException) Error() string { return e.s }

// NoArgs is used by IsNoArgs to check the interface.
func (e *NoArgsException) NoArgs() bool { return true }

// IsNoArgs checks the given error implements NoArgsException or do not.
func IsNoArgs(err error) bool {
	b, ok := errors.Cause(err).(interface{ NoArgs() bool })
	return ok && b.NoArgs()
}

// ParseArgs initializes a DirConv struct with given command line arguments.
func ParseArgs(args []string) (*dirconv.DirConv, error) {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		srcExt = flags.String("source_ext", "jpg", usageSrcExt)
		tgtExt = flags.String("target_ext", "png", usageTgtExt)
	)
	err := flags.Parse(args[1:])
	if err != nil {
		return &dirconv.DirConv{}, err
	}

	dir := flags.Arg(0) // get the first dir name only

	// if no args, return an NoArgsException indicates the program should show an usage message
	if len(dir) == 0 {
		return &dirconv.DirConv{}, &NoArgsException{"please show usage"}
	}

	return &dirconv.DirConv{Dir: dir, SrcExt: conv.ParseImgExt(*srcExt), TgtExt: conv.ParseImgExt(*tgtExt)}, nil
}

func main() {
	dc, err := ParseArgs(os.Args)
	if err != nil {
		if IsNoArgs(err) {
			// show usage if no dir specified
			fmt.Fprintf(os.Stdout, "%s\n", usage)
		} else {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		os.Exit(0)
	}
	_, err = dc.Convert()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
