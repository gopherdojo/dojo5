// Package dirconv provides a per directory image converter.
// Using package dir to traverse directories
// and using package conv to convert images.
package dirconv

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/conv"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/dir"
	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
	"github.com/pkg/errors"
)

// Logger outputs logs one after another during Convert.
var Logger *log.Logger

// DirConv struct
type DirConv struct {
	// directory name to traverse
	Dir string
	// source extension
	SrcExt img.Ext
	// target extension
	TgtExt img.Ext
}

// Result struct
type Result struct {
	// this value is usually not continuous because Convert uses goroutine
	Index int
	// relative path from the dir passed to args
	RelPath string
	// if err == nil then true
	Err error
}

// Convert runs an imgconv command (parsed by ParseArgs()).
// 1. traverses dirs
// 2. converts files
// 3. shows logs and returns results
// Returns a list of results likes (with an error):
//   [{0 dummy.jpg someErrorString} {2 dirA/figB.jpg <nil>} {1 figA.jpg <nil>} ...], <nil>
func (dc *DirConv) Convert() ([]*Result, error) {
	var results []*Result

	// get file paths to convert
	files, err := dir.TraverseImageFiles(dc.Dir, dc.SrcExt)
	if err != nil {
		return []*Result{}, errors.Wrapf(err, "failed to traverse %s", dc.Dir)
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
			logStr := fmt.Sprintf("%s -> %s", oldFileName, newFileName)

			// make a new ImgConv with file paths and file extensions
			ic := &conv.ImgConv{SrcPath: oldFileName, SrcExt: dc.SrcExt, TgtPath: newFileName, TgtExt: dc.TgtExt, Options: nil}

			// do convert and check the result
			err := ic.Convert()
			if err != nil {
				logStr = fmt.Sprintf("[Failed] %s", logStr)
			} else {
				logStr = fmt.Sprintf("[OK] %s", logStr)
			}

			// make a new Result and append it to the results list
			results = append(results, &Result{Index: idx, RelPath: val, Err: err})
			if Logger != nil {
				Logger.Printf("%s\n", logStr)
			}
		}(i, v)
	}
	wg.Wait()

	return results, nil
}
