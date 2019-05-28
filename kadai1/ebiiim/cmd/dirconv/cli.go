package dirconv

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"../conv"
)

const usageSrcExt = `source extension (jpg, png, tiff, bmp)`
const usageTgtExt = `target extension (jpg, png, tiff, bmp)`

// Usage string
const Usage = `Usage:
  imgconv [-source_ext=<ext>] [-target_ext=<ext>] DIR
Arguments:
  -source_ext=<ext>` + "\t" + usageSrcExt + ` [default: jpg]
  -target_ext=<ext>` + "\t" + usageTgtExt + ` [default: png]`

// Cli struct
type Cli struct {
	// directory name to traverse
	Dir string
	// source extension
	SrcExt string
	// target extension
	TgtExt string
}

// Result struct
type Result struct {
	Index   int
	RelPath string
	IsOk    bool
}

// DirConv run an imgconv command.
// 1. traverses dirs
// 2. converts files
// 3. shows logs and returns results
func (cli Cli) DirConv() []Result {
	var results []Result

	// show help if no dir specified
	if cli.Dir == "" {
		fmt.Println(Usage)
		os.Exit(0)
	}
	fmt.Println(cli)
	// get file paths to convert
	files, err := cli.traverseImageFiles()
	if err != nil {
		panic(err)
	}

	// TODO: goroutine
	for i, v := range files {
		oldFileName := fmt.Sprintf("%s/%s", cli.Dir, v)
		newFileName := oldFileName[0:len(oldFileName)-len(cli.SrcExt)] + cli.TgtExt
		log := fmt.Sprintf("%s -> %s", oldFileName, newFileName)

		ic := conv.ImgConv{SrcPath: oldFileName, TgtPath: newFileName}
		err := ic.Convert()

		ok := true
		if err != nil {
			ok = false
			_, _ = fmt.Fprintln(os.Stderr, err)
			log = fmt.Sprintf("[Failed] %s", log)
		} else {
			log = fmt.Sprintf("[OK] %s", log)
		}

		results = append(results, Result{Index: i, RelPath: v, IsOk: ok})
		fmt.Println(log)
	}

	return results
}

func (cli *Cli) traverseImageFiles() (files []string, err error) {
	err = filepath.Walk(cli.Dir,
		func(path string, info os.FileInfo, err error) error {
			relPath, err := filepath.Rel(cli.Dir, path)
			if !info.IsDir() && err == nil && strings.ToLower(filepath.Ext(relPath)) == cli.SrcExt {
				files = append(files, relPath)
			}
			return nil
		})
	return
}

// NewCli initializes imgconv dirconv.
func NewCli() (cli *Cli) {
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
