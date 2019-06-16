// Package imgconv provides multiple or single image converting ability
// by calling ConvertDir() or ConvertImg() method,
// or it could be judged automatically by calling Convert() method.
//
// All images in given directory and its sub directories with given image type
// will be converted recursively.
// Each image and sub directory will be processed concurrently in a goroutine.
package imgconv

import (
	"fmt"
	pkgimg "image"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// ImgExt represents image extension
type ImgExt string

// ImgType represents image type
type ImgType string

// Decoder decodes a io.Reader to pkgimg.Image
type Decoder interface {
	Decode(r io.Reader) (pkgimg.Image, error)
}

// Encoder encodes a io.Writer to pkgimg.Image
type Encoder interface {
	Encode(w io.Writer, m pkgimg.Image) error
}

// ImgPicker picks all image paths inside `path`
type ImgPicker interface {
	Pick(path string, imgType ImgType) (imgPaths []string, err error)
}

// Converter can convert image to other image type
// with given Decoder, Encoder
type Converter struct {
	DestImgExt ImgExt
	Dec        Decoder
	Enc        Encoder
	Picker     ImgPicker
	// SkipErr indicates converter will keep continuing to next file
	// or stop and return error when convert single image one by one.
	//
	// Eg: converter converts fileA.png, fileB.png and fileC.png,
	// fileA.png is converted successfully, fileB.png is error on converting, then:
	// if SkipErr is true, converter returns error and stop converting immediately,
	// if SkipErr is false, converter ignores this error and keeps converting to next files,
	// that means fileC.png will be converted later.
	SkipErr bool
	// KeepSrcImg indicates the original image is kept after converted.
	KeepSrcImg bool
	// errorOnConv returns error of image when converting
	// error value is based on SkipErr
	errOnConvImg func(err error) error
	path         string
	srcImgType   ImgType
}

// Convert is a wapper of ConvertDir and ConvertImg method,
// that calls to ConvertDir or ConvertImg method based on `path` is a directory or a file
func (conv *Converter) Convert(path string, srcImgType ImgType) error {
	fInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("error on stat %s: %+v", path, err)
	}

	if fInfo.IsDir() {
		return conv.ConvertDir(path, srcImgType)
	}

	return conv.ConvertImg(path)
}

// ConvertDir converts all image files in `dirPath`
// from `srcType` to `destType` image type recursively.
func (conv *Converter) ConvertDir(dirPath string, srcImgType ImgType) error {
	conv.path = dirPath
	if conv.Picker == nil {
		conv.Picker = DefaultImgPicker{}
	}
	conv.srcImgType = srcImgType
	conv.errOnConvImg = errBuilder(conv.SkipErr)
	if err := conv.validate(true); err != nil {
		return errors.Wrap(err, "invalid Converter")
	}
	return conv.convertDir()
}

// ConvertImg coverts single image
// from `srcType` to `destType` image type.
func (conv *Converter) ConvertImg(imgPath string) error {
	conv.errOnConvImg = errBuilder(conv.SkipErr)
	conv.path = imgPath
	if err := conv.validate(false); err != nil {
		return errors.Wrap(err, "invalid Converter")
	}
	return conv.errOnConvImg(conv.convertImg(imgPath))
}

func (conv *Converter) validate(isConvertDir bool) error {
	fInfo, err := os.Stat(conv.path)
	if err != nil {
		return errors.Wrapf(err, "cannot stat path: %s", conv.path)
	}
	if fInfo.IsDir() && !isConvertDir {
		return fmt.Errorf("%s is dir, use ConvertDir method instead", conv.path)
	}
	if !fInfo.IsDir() && isConvertDir {
		return fmt.Errorf("%s is file, use ConvertImg method instead", conv.path)
	}
	if conv.Dec == nil {
		return fmt.Errorf("decoder is not set")
	}
	if conv.Enc == nil {
		return fmt.Errorf("encoder is not set")
	}
	if conv.Picker == nil {
		return fmt.Errorf("image file picker is not set")
	}
	if conv.DestImgExt == ImgExt("") {
		return fmt.Errorf("destination extension is not set")
	}
	return nil
}

func (conv *Converter) convertDir() error {
	imgPaths, err := conv.Picker.Pick(conv.path, conv.srcImgType)
	if err != nil {
		return errors.Wrapf(err, "error on pick images in path: %s", conv.path)
	}

	var eg errgroup.Group
	// convert all image in current directory recursively and concurrency
	eg.Go(func() error {
		var innerEg errgroup.Group
		for _, ip := range imgPaths {
			imgPath := ip
			innerEg.Go(func() error {
				return conv.errOnConvImg(conv.convertImg(imgPath))
			})
		}

		return innerEg.Wait()
	})

	return eg.Wait()
}

func (conv *Converter) convertImg(imgPath string) error {
	img := image{path: imgPath}
	srcImg, err := os.Open(img.path)
	defer srcImg.Close()
	if err != nil {
		return fmt.Errorf("error on open file to convert. file: %s, err: %+v", img.path, err)
	}

	decodedImg, err := conv.Dec.Decode(srcImg)
	if err != nil {
		return fmt.Errorf("error on decode %+v, err: %+v", img, err)
	}

	destPath := img.toDestImgPath(conv.DestImgExt)
	destImg, err := os.Create(destPath)
	defer destImg.Close()
	if err != nil {
		return fmt.Errorf("error on create dest file at path: %s, err: %+v", destPath, err)
	}

	if err := conv.Enc.Encode(destImg, decodedImg); err != nil {
		return fmt.Errorf("error on encode img: %+v, error: %+v", img, err)
	}

	if !conv.KeepSrcImg {
		if err := img.remove(); err != nil {
			return fmt.Errorf("error on remove old img: %+v", err)
		}
	}

	fmt.Printf("converted %+v to %s\n", img, conv.DestImgExt)
	return nil
}

type image struct {
	path string
}

func (img image) toDestImgPath(destImgExt ImgExt) string {
	ext := filepath.Ext(img.path)
	destPath := img.path[0:len(img.path)-len(ext)] + "." + string(destImgExt)
	return destPath
}

func (img image) remove() error {
	return os.Remove(img.path)
}

// errBuilder returns function that
// print errMsg to stdout and returns nil in case of skipErr
// or print to stderr and return new error otherwise
func errBuilder(skipErr bool) func(err error) error {
	if skipErr {
		return func(err error) error {
			if err != nil {
				fmt.Println(err)
			}
			return nil
		}
	}

	return func(err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return fmt.Errorf("%+v", err)
		}

		return nil
	}
}
