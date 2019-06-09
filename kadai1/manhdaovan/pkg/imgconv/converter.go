// Package imgconv provides image converting ability
// by calling ConvertDir() method
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

// RegisterDestImgType is used when you want to encode to a new custom destination image type,
// that not in default support of this package.
// Ref to README.md file for more details about how to do it.
func RegisterDestImgType(destImgType ImgType, enc Encoder, destImgExt ImgExt) error {
	if err := registerNewEncoder(destImgType, enc); err != nil {
		return err
	}
	if err := registerNewExt(destImgType, destImgExt); err != nil {
		return err
	}

	return nil
}

// RegisterSrcImgType is used when you want to decode to a new custom destination image type,
// that not in default support of this package.
// Ref to README.md file for more details about how to do it.
func RegisterSrcImgType(srcImgType ImgType, dec Decoder) error {
	if err := registerImgType(srcImgType); err != nil {
		return err
	}
	if err := registerNewDecoder(srcImgType, dec); err != nil {
		return err
	}
	return nil
}

// Decoder decodes a io.Reader to pkgimg.Image
type Decoder interface {
	Decode(r io.Reader) (pkgimg.Image, error)
}

// Encoder encodes a io.Writer to pkgimg.Image
type Encoder interface {
	Encode(w io.Writer, m pkgimg.Image) error
}

// ConvertDir converts all images having srcImgType type
// to destImgType type in dirPath path recursively
func ConvertDir(dirPath string, srcImgType, destImgType ImgType, skipErr bool) error {
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
	conv.decoder = GetDecoder(conv.srcImgType)
	conv.encoder = GetEncoder(conv.destImgType)
	conv.destImgExt = GetExtension(conv.destImgType)
	conv.errOnConvImg = errBuilder(conv.skipErr)

	if err := conv.validate(); err != nil {
		return err
	}
	if err := conv.convert(); err != nil {
		return err
	}

	return nil
}

type converter struct {
	dirPath     string
	srcImgType  ImgType
	destImgType ImgType
	destImgExt  ImgExt
	decoder     Decoder
	encoder     Encoder
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

func (conv converter) validate() error {
	if conv.dirPath == "" {
		return fmt.Errorf("dir path is not set")
	}
	if conv.srcImgType == ImgType("") {
		return fmt.Errorf("src img type is not set")
	}
	if conv.destImgType == ImgType("") {
		return fmt.Errorf("dest img type is not set")
	}
	if conv.destImgExt == ImgExt("") {
		return fmt.Errorf("extension is not found for %s", conv.destImgExt)
	}
	if conv.decoder == nil {
		return fmt.Errorf("decoder not found for img type: %s", conv.srcImgType)
	}
	if conv.encoder == nil {
		return fmt.Errorf("encoder not found for img type: %s", conv.destImgType)
	}
	if conv.destImgExt == ImgExt("") {
		return fmt.Errorf("destination image extension not found for img type: %s", conv.destImgType)
	}

	return nil
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

	decodedImg, err := conv.decoder.Decode(srcImg)
	if err != nil {
		return fmt.Errorf("error on decode %+v, err: %+v", img, err)
	}

	destPath := img.toDestImgPath(conv.destImgType, conv.destImgExt)
	destImg, err := os.Create(destPath)
	defer destImg.Close()
	if err != nil {
		return fmt.Errorf("error on create dest file at path: %s, err: %+v", destPath, err)
	}

	if err := conv.encoder.Encode(destImg, decodedImg); err != nil {
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

func (img image) toDestImgPath(destImgType ImgType, destImgExt ImgExt) string {
	ext := filepath.Ext(img.path)
	destPath := img.path[0:len(img.path)-len(ext)] + "." + string(destImgExt)
	return destPath
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
