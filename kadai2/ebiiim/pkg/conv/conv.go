package conv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// ImgConv struct holds two paths and converter options.
type ImgConv struct {
	// source path e.g. `path/to/A.jpg`
	SrcPath string
	// source image extension
	SrcExt ImgExt
	// target path e.g. `path/to/A.png`
	TgtPath string
	// target image extension
	TgtExt ImgExt
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

	srcImg, err := decodeImg(srcFile, ic.SrcExt)
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

	err = encodeImg(tgtFile, ic.TgtExt, &srcImg, ic.Options)
	if err != nil {
		rErr := os.Remove(ic.TgtPath)
		if rErr != nil {
			err = fmt.Errorf("failed to remove: %v (%v)", rErr, err)
		}
		return fmt.Errorf("failed to encode: %s", ic.SrcPath)
	}

	return nil
}

func encodeImg(file *os.File, ext ImgExt, img *image.Image, options map[string]interface{}) (err error) {
	switch ext {
	case ImgExtJPEG:
		// TODO: apply encoder options
		err = jpeg.Encode(file, *img, &jpeg.Options{})
	case ImgExtPNG:
		err = png.Encode(file, *img)
	case ImgExtTIFF:
		// TODO: apply encoder options
		err = tiff.Encode(file, *img, &tiff.Options{})
	case ImgExtBMP:
		err = bmp.Encode(file, *img)
	default:
		err = fmt.Errorf("unsupported image extension: %s", ext)
	}
	return
}

func decodeImg(file *os.File, ext ImgExt) (img image.Image, err error) {
	switch ext {
	case ImgExtJPEG:
		img, err = jpeg.Decode(file)
	case ImgExtPNG:
		img, err = png.Decode(file)
	case ImgExtTIFF:
		img, err = tiff.Decode(file)
	case ImgExtBMP:
		img, err = bmp.Decode(file)
	default:
		err = fmt.Errorf("unsupported image extension: %s", ext)
	}
	return
}
