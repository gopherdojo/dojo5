package dir

import (
	"io/ioutil"
	"path/filepath"
)

// Lookup lookup directory to find image files
func Lookup(dir string, ext string, pathList []string) ([]string, error) {

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {

		// get file info
		path := filepath.Join(dir, fileInfo.Name())

		// check if the path is about directory
		if fileInfo.IsDir() {
			pathList, err = Lookup(path, ext, pathList)
			if err != nil {
				return nil, err
			}
		}

		// check if the postfix is equal to the input format
		if filepath.Ext(path) == ext {
			pathList = append(pathList, path)
		}
	}

	return pathList, nil
}
