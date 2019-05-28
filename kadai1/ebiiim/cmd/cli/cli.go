package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const usageSrcExt = `source extension (jpg, png, tiff, bmp)`
const usageTgtExt = `target extension (jpg, png, tiff, bmp)`

// Usage string
const Usage = `Usage:
  imgconv DIR [-source_ext=<ext>] [-target_ext=<ext>]
Arguments:
  -source_ext=<ext>` + "\t" + usageSrcExt + ` [default: jpg]
  -target_ext=<ext>` + "\t" + usageTgtExt + ` [default: png]`

// Cli struct
type Cli struct {
	dir    string
	srcExt string
	tgtExt string
}

// Run an imgconv command.
// 1. parse arguments
// 2. traverse dirs
// 3. convert files
func Run() {
	cli := parseArgs()
	if cli.dir == "" {
		fmt.Println(Usage)
		os.Exit(0)
	}
	//fmt.Println(cli)

	files, err := cli.traverseImageFiles()
	if err != nil {
		panic(err)
	}
	//fmt.Println(files)

	for _, oldFileName := range files {
		newFileName := oldFileName[0:len(oldFileName)-len(cli.srcExt)] + cli.tgtExt
		// TODO: convert files
		fmt.Printf("%s -> %s\n", oldFileName, newFileName)
	}
}

func (cli *Cli) traverseImageFiles() (files []string, err error) {
	err = filepath.Walk(cli.dir,
		func(path string, info os.FileInfo, err error) error {
			relPath, err := filepath.Rel(cli.dir, path)
			if !info.IsDir() && err == nil && strings.ToLower(filepath.Ext(relPath)) == cli.srcExt {
				files = append(files, relPath)
			}
			return nil
		})
	return
}

func parseArgs() (cli *Cli) {
	srcExt := flag.String("source_ext", "jpg", usageSrcExt)
	tgtExt := flag.String("target_ext", "png", usageTgtExt)
	flag.Parse()
	dir := flag.Arg(0) // get the first dir name only

	formatExt(srcExt)
	formatExt(tgtExt)

	cli = &Cli{dir, *srcExt, *tgtExt}
	return
}

func formatExt(ext *string) {
	*ext = strings.ToLower(*ext)
	if !strings.HasPrefix(*ext, ".") {
		*ext = "." + *ext
	}
}
