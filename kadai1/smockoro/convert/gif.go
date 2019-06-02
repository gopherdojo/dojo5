package convert

import (
	"image"
	"image/gif"
	"os"
	"path/filepath"
	"strings"
)

type gifConverter struct{}

func NewGifConverter() Converter {
	return &gifConverter{}
}

func (gc *gifConverter) ImageConvert(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newpath := strings.TrimSuffix(path, filepath.Ext(path)) + ".gif"
	outfile, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	opts := &gif.Options{NumColors: 256}

	err = gif.Encode(outfile, img, opts)
	if err != nil {
		return err
	}

	return nil
}
