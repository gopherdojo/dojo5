/*
Package 'search' provides file search function
*/
package search

import (
	"os"
	"path/filepath"
)

// WalkWithExtHandle searches the target directory recursively for files with matching extensions.
// If a matching file is found, the handle function is called.
func WalkWithExtHandle(dir string, ext []string, handle func(string, os.FileInfo, error)) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if contains(ext, filepath.Ext(path)) {
			handle(path, info, err)
		}
		return nil
	})
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
