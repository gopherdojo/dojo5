package convert

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

type jpgConverter struct{}

func NewJpgConverter() Converter {
	return &jpgConverter{}
}

func (jc *jpgConverter) ImageConvert(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newpath := strings.TrimSuffix(path, filepath.Ext(path)) + ".jpg"
	outfile, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	opts := &jpeg.Options{Quality: 100}
	err = jpeg.Encode(outfile, img, opts)
	if err != nil {
		return err
	}

	return nil
}
