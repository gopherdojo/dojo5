// Package imgconv provides image converting ability
// by calling Convert() method
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

	"golang.org/x/sync/errgroup"
)

// ImgType represents image type option
// that this converter support
type ImgType string

const (
	// ImgTypeJPEG is constant for JPEG image
	ImgTypeJPEG ImgType = "jpeg"
	// ImgTypePNG is constant for PNG image
	ImgTypePNG ImgType = "png"
	// ImgTypeGIF is constant for GIF image
	ImgTypeGIF ImgType = "gif"
)

var extensions = map[ImgType]string{
	ImgTypeJPEG: "jpg",
	ImgTypePNG:  "png",
	ImgTypeGIF:  "gif",
}

// SupportSrcImgTypes is a list of ImgType
type SupportSrcImgTypes []ImgType

// IsSupport returns imgType is supported by this tool or not
func (ssits SupportSrcImgTypes) IsSupport(imgType string) bool {
	for _, it := range ssits {
		if string(it) == imgType {
			return true
		}
	}

	return false
}

// GetSupportSrcImgTypes returns all supported types
// of input and output image
func GetSupportSrcImgTypes() SupportSrcImgTypes {
	return SupportSrcImgTypes{ImgTypeJPEG, ImgTypePNG, ImgTypeGIF}
}

// decoder decodes a io.Reader to pkgimg.Image
type decoder interface {
	decode(r io.Reader) (pkgimg.Image, error)
}

// encoder encodes a io.Writer to pkgimg.Image
type encoder interface {
	encode(w io.Writer, m pkgimg.Image) error
}

// Convert converts all images having srcImgType type
// to destImgType type in dirPath path recursively
func Convert(dirPath string, srcImgType, destImgType ImgType, skipErr bool) error {
	if srcImgType == destImgType {
		// same image format, then nothing todo
		return nil
	}

	conv := converter{
		dirPath:     dirPath,
		srcImgType:  srcImgType,
		destImgType: destImgType,
		skipErr:     skipErr,
	}

	decoder := getDecoder(conv.srcImgType)
	if decoder == nil {
		return fmt.Errorf("decoder not found for img type: %s", conv.srcImgType)
	}
	conv.dec = decoder

	encoder := getEncoder(conv.destImgType)
	if encoder == nil {
		return fmt.Errorf("encoder not found for img type: %s", conv.destImgType)
	}
	conv.enc = encoder

	conv.errOnConvImg = errBuilder(conv.skipErr)

	if err := conv.convert(); err != nil {
		return err
	}

	return nil
}

type converter struct {
	dirPath     string
	srcImgType  ImgType
	destImgType ImgType
	dec         decoder
	enc         encoder
	// skipErr indicates converter will keep continuing to next file
	// or stop and return error when convert single image one by one.
	//
	// Eg: converter converts fileA.png, fileB.png and fileC.png,
	// fileA.png is converted successfully, fileB.png is error on converting, then:
	// if skipErr is true, converter returns error and stop converting immediately,
	// if skipErr is false, converter ignores this error and keeps converting to next files,
	// that means fileC.png will be converted later.
	skipErr bool
	// errorOnConv returns error of image when converting
	// error value is based on skipErr
	errOnConvImg func(err error) error
}

func (conv converter) convert() error {
	imgs, err := conv.pickSrcImgs()
	if err != nil {
		return err
	}

	var eg errgroup.Group

	// convert all image in current directory
	eg.Go(func() error {
		var innerEg errgroup.Group
		for _, i := range imgs {
			img := i
			innerEg.Go(func() error {
				return conv.errOnConvImg(conv.convertImg(img))
			})
		}

		return innerEg.Wait()
	})

	return eg.Wait()
}

func (conv converter) convertImg(img image) error {
	srcImg, err := os.Open(img.path)
	defer srcImg.Close()
	if err != nil {
		return fmt.Errorf("error on open file to convert. file: %s, err: %+v", img.path, err)
	}

	decodedImg, err := conv.dec.decode(srcImg)
	if err != nil {
		return fmt.Errorf("error on decode %+v, err: %+v", img, err)
	}

	destPath := img.toDestImgPath(conv.destImgType)
	destImg, err := os.Create(destPath)
	defer destImg.Close()
	if err != nil {
		return fmt.Errorf("error on create dest file at path: %s, err: %+v", destPath, err)
	}

	if err := conv.enc.encode(destImg, decodedImg); err != nil {
		return fmt.Errorf("error on encode img: %+v, error: %+v", img, err)
	}

	if err := img.remove(); err != nil {
		return fmt.Errorf("error on remove old img: %+v", err)
	}

	fmt.Printf("converted %+v to %s\n", img, conv.destImgType)
	return nil
}

func (conv converter) pickSrcImgs() (imgs []image, outerErr error) {
	outerErr = filepath.Walk(conv.dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() { // file
			if ok, err := isImgWithType(path, conv.srcImgType); err != nil {
				return err
			} else if ok {
				imgs = append(imgs, image{path: path})
			}
		}

		return nil
	})

	return
}

type image struct {
	path string
}

func (img image) toDestImgPath(destImgType ImgType) string {
	ext := filepath.Ext(img.path)
	newExt := extensions[destImgType]
	return img.path[0:len(img.path)-len(ext)] + "." + newExt
}

func (img image) remove() error {
	return os.Remove(img.path)
}

func isImgWithType(path string, imgType ImgType) (bool, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return false, err
	}

	_, format, err := pkgimg.DecodeConfig(file)
	if err != nil {
		return false, nil // skip unknown format file
	}

	return format == string(imgType), nil
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
