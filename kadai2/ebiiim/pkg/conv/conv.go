package conv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/pkg/errors"
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
	srcFile, err := os.Open(ic.SrcPath)
	if err != nil {
		return errors.Wrapf(err, "failed to open %s", ic.SrcPath)
	}
	defer func() {
		dErr := srcFile.Close()
		if dErr != nil {
			err = errors.Wrapf(err, "failed to close %v", dErr)
		}
	}()

	srcImg, err := decodeImg(srcFile, ic.SrcExt)
	if err != nil {
		return errors.Wrapf(err, "failed to decode %s", ic.SrcPath)
	}

	// write encoded image to the target file
	tgtFile, err := os.Create(ic.TgtPath)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s", ic.TgtPath)
	}
	defer func() {
		dErr := tgtFile.Close()
		if dErr != nil {
			err = errors.Wrapf(err, "failed to close %v", dErr)
		}
	}()

	err = encodeImg(tgtFile, ic.TgtExt, &srcImg, ic.Options)
	if err != nil {
		rErr := os.Remove(ic.TgtPath)
		if rErr != nil {
			err = errors.Wrapf(err, "failed to remove %v", rErr)
		}
		return errors.Wrapf(err, "failed to encode %s", ic.SrcPath)
	}

	return nil
}

func encodeImg(writer io.Writer, ext ImgExt, img *image.Image, options map[string]interface{}) (err error) {
	switch ext {
	case ImgExtJPEG:
		// TODO: apply encoder options
		err = jpeg.Encode(writer, *img, &jpeg.Options{})
	case ImgExtPNG:
		err = png.Encode(writer, *img)
	case ImgExtTIFF:
		// TODO: apply encoder options
		err = tiff.Encode(writer, *img, &tiff.Options{})
	case ImgExtBMP:
		err = bmp.Encode(writer, *img)
	default:
		err = fmt.Errorf("unsupported image extension %s", ext)
	}
	return
}

func decodeImg(reader io.Reader, ext ImgExt) (img image.Image, err error) {
	switch ext {
	case ImgExtJPEG:
		img, err = jpeg.Decode(reader)
	case ImgExtPNG:
		img, err = png.Decode(reader)
	case ImgExtTIFF:
		img, err = tiff.Decode(reader)
	case ImgExtBMP:
		img, err = bmp.Decode(reader)
	default:
		err = fmt.Errorf("unsupported image extension %s", ext)
	}
	return
}
