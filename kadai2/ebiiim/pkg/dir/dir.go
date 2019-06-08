package dir

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
)

func TraverseImageFiles(targetDir string, ext img.Ext) ([]string, error) {
	var (
		files []string
		err   error
	)

	// check the dir exists
	fileInfo, err := os.Stat(targetDir)
	if err != nil {
		return files, err // if the dir does not exist, return an empty slice
	}
	if !fileInfo.IsDir() {
		return files, fmt.Errorf("%s is not a directory", targetDir)
	}

	// traverse the dir and return a list of image files has the given file extension
	err = filepath.Walk(targetDir,
		func(path string, info os.FileInfo, err error) error {
			relPath, err := filepath.Rel(targetDir, path)
			if !info.IsDir() && err == nil && img.ParseExt(relPath) == ext {
				files = append(files, relPath)
			}
			return nil
		})
	return files, err
}
