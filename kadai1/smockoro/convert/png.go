package convert

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type pngConverter struct{}

func NewPngConverter() Converter {
	return &pngConverter{}
}

func (pc *pngConverter) ImageConvert(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newpath := strings.TrimSuffix(path, filepath.Ext(path)) + ".png"
	outfile, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	err = png.Encode(outfile, img)
	if err != nil {
		return err
	}

	return nil
}
