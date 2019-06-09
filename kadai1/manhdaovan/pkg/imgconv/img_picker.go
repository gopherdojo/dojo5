package imgconv

import (
	pkgimg "image"
	"os"
	"path/filepath"
)

// DefaultImgPicker is default image picker,
// that implements ImgPicker interface
type DefaultImgPicker struct{}

// Pick picks image by given `imgType` in `path`.
// If given `path` is a directory, all files with `imgType` will be picked recursively,
// or `path` will be returned if `path` is a satisfied image.
func (dp DefaultImgPicker) Pick(path string, imgType ImgType) (imgPaths []string, outerErr error) {
	outerErr = filepath.Walk(path, func(fPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() { // file
			if ok, err := isImgWithType(fPath, imgType); err != nil {
				return err
			} else if ok {
				imgPaths = append(imgPaths, fPath)
			}
		}

		return nil
	})

	return
}

func isImgWithType(path string, imgType ImgType) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, format, err := pkgimg.DecodeConfig(file)
	if err != nil {
		return false, nil // skip unknown format file
	}

	return format == string(imgType), nil
}
