package dirconv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/conv"
	"github.com/pkg/errors"
)

// DirConv struct
type DirConv struct {
	// directory name to traverse
	Dir string
	// source extension
	SrcExt conv.ImgExt
	// target extension
	TgtExt conv.ImgExt
}

// Result struct
type Result struct {
	// this value is usually not continuous because Convert uses goroutine
	Index int
	// relative path from the dir passed to args
	RelPath string
	// if err == nil then true
	IsOk bool
}

// Convert runs an imgconv command (parsed by ParseArgs()).
//   1. traverses dirs
//   2. converts files
//   3. shows logs and returns results
// Returns a list of results likes (with an error):
//   [{0 dummy.jpg false} {2 dirA/figB.jpg true} {1 figA.jpg true} ...], nil
func (dc *DirConv) Convert() ([]Result, error) {
	var results []Result

	// get file paths to convert
	files, err := dc.traverseImageFiles()
	if err != nil {
		return []Result{}, errors.Wrapf(err, "failed to traverse %s", dc.Dir)
	}

	// convert files (goroutined)
	var wg sync.WaitGroup
	for i, v := range files {
		wg.Add(1)
		go func(idx int, val string) {
			defer wg.Done()

			// make file paths
			oldFileName := filepath.Join(dc.Dir, val)
			newFileName := fmt.Sprintf("%s.%s", strings.TrimSuffix(oldFileName, filepath.Ext(oldFileName)), dc.TgtExt)
			log := fmt.Sprintf("%s -> %s", oldFileName, newFileName)

			// make a new ImgConv with file paths and file extensions
			ic := conv.ImgConv{SrcPath: oldFileName, SrcExt: dc.SrcExt, TgtPath: newFileName, TgtExt: dc.TgtExt, Options: nil}

			// do convert and check the result
			err := ic.Convert()
			ok := true
			if err != nil {
				ok = false
				_, _ = fmt.Fprintln(os.Stderr, err)
				log = fmt.Sprintf("[Failed] %s", log)
			} else {
				log = fmt.Sprintf("[OK] %s", log)
			}

			// make a new Result and append it to the results list
			results = append(results, Result{Index: idx, RelPath: val, IsOk: ok})
			fmt.Println(log)
		}(i, v)
	}
	wg.Wait()

	return results, nil
}

func (dc *DirConv) traverseImageFiles() ([]string, error) {
	var (
		files []string
		err   error
	)

	// check the dir exists
	fileInfo, err := os.Stat(dc.Dir)
	if err != nil {
		return files, err // if the dir does not exist, return an empty slice
	}
	if !fileInfo.IsDir() {
		return files, fmt.Errorf("%s is not a directory", dc.Dir)
	}

	// traverse the dir and return a list of image files has the given file extension
	err = filepath.Walk(dc.Dir,
		func(path string, info os.FileInfo, err error) error {
			relPath, err := filepath.Rel(dc.Dir, path)
			if !info.IsDir() && err == nil && conv.ParseImgExt(relPath) == dc.SrcExt {
				files = append(files, relPath)
			}
			return nil
		})
	return files, err
}
