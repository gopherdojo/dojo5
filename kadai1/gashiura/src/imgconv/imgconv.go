// Package imgconv is a lib to convert image formats.
package imgconv

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Converter struct {
	// image extension before conversion
	SrcExt string
	// image extension after conversion
	DstExt string
}

const (
	extJpeg = "jpeg"
	extPng = "png"
	extGif = "gif"
)

/*
Convert converts the image format file specified as the conversion source under the specified directory
into the image format of the conversion destination
*/
func (conv Converter) Convert(rootDir string) error {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "." + conv.SrcExt {
			file, err := os.Open(path)
			if err != nil { return err }
			defer file.Close()

			img, err := conv.decodeImage(file)
			if err != nil { return err }

			outPath := strings.Replace(path, conv.SrcExt, conv.DstExt, -1)

			out, err := os.Create(outPath)
			if err != nil { return err }
			defer out.Close()

			err = conv.encodeImage(out, img)
			if err != nil { return err }
		}
		return nil
	})

	return err
}

func (conv Converter) decodeImage(file io.Reader) (img image.Image, err error) {
	switch conv.SrcExt {
	case extJpeg:
		img, err = jpeg.Decode(file)
	case extPng:
		img, err = png.Decode(file)
	case extGif:
		img, err = gif.Decode(file)
	default:
		img, err = image.Image(nil), errors.New("extension is incorrect.")
	}
	return
}

func (conv Converter) encodeImage(file io.Writer, image image.Image) (err error) {
	switch conv.DstExt {
	case extJpeg:
		opt := &jpeg.Options{Quality: 100}
		err = jpeg.Encode(file, image, opt)
	case extPng:
		err = png.Encode(file, image)
	case extGif:
		err = gif.Encode(file, image, nil)
	default:
		err = errors.New("extension is incorrect.")
	}
	return
}
