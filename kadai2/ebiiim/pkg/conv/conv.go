package conv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/gopherdojo/dojo5/kadai2/ebiiim/pkg/img"
	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// ImgConv struct holds two paths and converter options.
type ImgConv struct {
	// source path e.g. `path/to/A.jpg`
	SrcPath string
	// source image extension
	SrcExt img.Ext
	// target path e.g. `path/to/A.png`
	TgtPath string
	// target image extension
	TgtExt img.Ext
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
	srcImg, err := decodeImg(srcFile, ic.SrcExt)
	if err != nil {
		return errors.Wrapf(err, "failed to decode %s", ic.SrcPath)
	}
	err = srcFile.Close()
	if err != nil {
		return errors.Wrapf(err, "failed to close %v", err)
	}

	// write encoded image to the target file
	tgtFile, err := os.Create(ic.TgtPath)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s", ic.TgtPath)
	}
	err = encodeImg(tgtFile, ic.TgtExt, &srcImg, ic.Options)
	cErr := tgtFile.Close() // close the file before it may be deleted.
	if cErr != nil {
		return errors.Wrapf(cErr, "failed to close %v", cErr)
	}
	if err != nil {
		rErr := os.Remove(ic.TgtPath) // if failed to encode, delete the unnecessary zero-byte file
		if rErr != nil {
			err = errors.Wrapf(err, "failed to remove %v", rErr)
		}
		return errors.Wrapf(err, "failed to encode %s", ic.SrcPath)
	}

	return nil
}

func encodeImg(writer io.Writer, ext img.Ext, image *image.Image, options map[string]interface{}) (err error) {
	switch ext {
	case img.JPEG:
		// TODO: apply encoder options
		err = jpeg.Encode(writer, *image, &jpeg.Options{})
	case img.PNG:
		err = png.Encode(writer, *image)
	case img.TIFF:
		// TODO: apply encoder options
		err = tiff.Encode(writer, *image, &tiff.Options{})
	case img.BMP:
		err = bmp.Encode(writer, *image)
	default:
		err = fmt.Errorf("unsupported image extension %s", ext)
	}
	return
}

func decodeImg(reader io.Reader, ext img.Ext) (image image.Image, err error) {
	switch ext {
	case img.JPEG:
		image, err = jpeg.Decode(reader)
	case img.PNG:
		image, err = png.Decode(reader)
	case img.TIFF:
		image, err = tiff.Decode(reader)
	case img.BMP:
		image, err = bmp.Decode(reader)
	default:
		err = fmt.Errorf("unsupported image extension %s", ext)
	}
	return
}
