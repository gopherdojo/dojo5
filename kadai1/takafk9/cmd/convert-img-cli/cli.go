package main

import (
	"flag"
	"fmt"
	"github.com/takafk9/dojo5/kadai1/takafk9/converter"
	"io"
)

const (
	ExitCodeOK = iota
	ExitCodeExpectedError
	ExitCodeUnexpectedError
	ExitCodeInvalidArgsError
	ExitCodeParseFlagsError
	ExitCodeInvalidFlagError
)

var usage = `Usage: convert-img-cli [options...] PATH

converter-cli is a command line tool to convert image extension

OPTIONS:
  --from value, -f value  specifies a image extension converted from (default: .jpg)
  --to value, -t value    specifies a image extension converted to (default: .png)
  --help, -h              prints out help
`

const Name = "convert-img-cli"

var exts = [4]string{".gif", ".jpeg", ".jpg", ".png"}

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var from string
	var to string

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(cli.outStream, usage)
	}

	flags.StringVar(&from, "from", ".jpg", "")
	flags.StringVar(&from, "f", ".jpg", "")

	flags.StringVar(&to, "to", ".png", "")
	flags.StringVar(&to, "t", ".png", "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if !validateExt(from) {
		fmt.Fprintf(cli.errStream, "Failed to set up convert-img-cli: invalid extension `%s` is given for --from flag\n"+
			"Please choose an extension from one of those: %v\n\n", from, exts)
		return ExitCodeInvalidFlagError
	}

	if !validateExt(to) {
		fmt.Fprintf(cli.errStream, "Failed to set up convert-img-cli: invalid extension `%s` is given for --to flag\n"+
			"Please choose an extension from one of those: %v\n\n", to, exts)
		return ExitCodeInvalidFlagError
	}

	path := flags.Args()
	if len(path) != 1 {
		fmt.Fprintf(cli.errStream, "Failed to set up convert-img-cli: invalid argument\n"+
			"Please specify the exact one path to a directly or a file\n\n")
		return ExitCodeInvalidArgsError
	}

	files, err := converter.Convert(from, to, path[0])
	if err != nil {
		if _, ok := err.(converter.Managed); ok {
			fmt.Fprintf(cli.errStream, "Failed to execute convert-img-cli\n"+
				"%s\n\n", err)
			return ExitCodeExpectedError
		}
		fmt.Fprintf(cli.errStream, `conver-img-cli failed because of the following error.

%s

You might encounter a bug with converter-cli, so please report it to https://github.com/xxx/xxxx

`, err)

		return ExitCodeUnexpectedError
	}

	fmt.Fprintf(cli.outStream, "convert-img-cli successfully converted following files to `%s`.\n", to)
	fmt.Fprintf(cli.outStream, "%s\n\n", files)
	return ExitCodeOK

}

func validateExt(ext string) bool {
	for _, e := range exts {
		if ext == e {
			return true
		}

	}
	return false
}
