package conv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// ImgConv struct holds two paths and converter options.
type ImgConv struct {
	// source path e.g. `path/to/A.jpg`
	SrcPath string
	// target path e.g. `path/to/A.png`
	TgtPath string
	// converter options (currently, this value is not interpreted by encodeImg())
	Options map[string]interface{}
}

// Convert an image file.
//   1. opens the source file and decodes it
//   2. encodes the image with the target format and writes it to the target file
func (ic *ImgConv) Convert() error {

	// load the source image
	srcFile, err := os.OpenFile(ic.SrcPath, os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open: %s", ic.SrcPath)
	}
	defer func() {
		dErr := srcFile.Close()
		if dErr != nil {
			err = fmt.Errorf("failed to close: %v (%v)", dErr, err)
		}
	}()

	srcImg, err := decodeImg(srcFile)
	if err != nil {
		return fmt.Errorf("failed to decode: %s", ic.SrcPath)
	}

	// write encoded image to the target file
	tgtFile, err := os.Create(ic.TgtPath)
	if err != nil {
		return fmt.Errorf("failed to create : %s", ic.TgtPath)
	}
	defer func() {
		dErr := tgtFile.Close()
		if dErr != nil {
			err = fmt.Errorf("failed to close: %v (%v)", dErr, err)
		}
	}()

	err = encodeImg(tgtFile, &srcImg, ic.Options)
	if err != nil {
		rErr := os.Remove(ic.TgtPath)
		if rErr != nil {
			err = fmt.Errorf("failed to remove: %v (%v)", rErr, err)
		}
		return fmt.Errorf("failed to encode: %s", ic.SrcPath)
	}

	return nil
}

func encodeImg(file *os.File, img *image.Image, options map[string]interface{}) (err error) {
	fileN := file.Name()
	fileE := strings.ToLower(filepath.Ext(fileN))
	switch fileE {
	case ".jpg", ".jpeg":
		// TODO: apply encoder options
		err = jpeg.Encode(file, *img, &jpeg.Options{})
	case ".png":
		err = png.Encode(file, *img)
	case ".tiff":
		// TODO: apply encoder options
		err = tiff.Encode(file, *img, &tiff.Options{})
	case ".bmp":
		err = bmp.Encode(file, *img)
	default:
		err = fmt.Errorf("unsupported image extension: %s", fileE)
	}
	return
}

func decodeImg(file *os.File) (img image.Image, err error) {
	fileN := file.Name()
	fileE := strings.ToLower(filepath.Ext(fileN))
	switch fileE {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".tiff":
		img, err = tiff.Decode(file)
	case ".bmp":
		img, err = bmp.Decode(file)
	default:
		err = fmt.Errorf("unsupported image extension: %s", fileE)
	}
	return
}
