// Package imgconv provides a recursive conversion of images in the directory.
package imgconv

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// MapImgFmtExts is a map of the image formats and its extensions.
type MapImgFmtExts map[ImgFmt]Exts

// Exts is a slice of image extensions.
type Exts []Ext

// Ext is a image extension.
type Ext string

// ImgFmt is a image format.
type ImgFmt string

// Converter is a struct which contains info about image formats and extensions.
type Converter struct {
	fmtFrom    ImgFmt
	fmtTo      ImgFmt
	imgFmtExts MapImgFmtExts
}

// Convert recursively seeks a given directory and converts images from and to given formats.
func Convert(dir string, from string, to string) error {
	if dir == "" {
		return errors.New("directory name is not provided")
	}

	cv := &Converter{}
	cv.imgFmtExts.Init()
	cv.fmtFrom.Detect(cv, from)
	cv.fmtTo.Detect(cv, to)
	if cv.fmtFrom == "" || cv.fmtTo == "" {
		return errors.New("given image format is not supported")
	}
	if cv.fmtFrom == cv.fmtTo {
		return errors.New("image formats before and after conversion are the same")
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		err = cv.convSingleFile(path, info)
		return err
	})
	return err
}

func (cv *Converter) convSingleFile(path string, info os.FileInfo) error {
	if info.IsDir() {
		// FIXME: create output directories
		outputPath := "./output/" + strings.TrimLeft(path, "./")
		return os.MkdirAll(outputPath, 0777)
	}
	if !cv.fmtFrom.Match(cv, info.Name()) {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := cv.decodeImg(f)
	if err != nil {
		return nil
	}

	return cv.writeOutputFile(img, path)
}

func (cv *Converter) writeOutputFile(img image.Image, path string) error {
	// FIXME: temporarily separate input and output directories
	outputPath := "./output/" + strings.TrimLeft(cv.generateOutputPath(path), "./")

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = cv.encodeImg(f, img)
	return err
}

func (cv *Converter) decodeImg(r io.Reader) (image.Image, error) {
	img, fmtStr, err := image.Decode(r)
	if ImgFmt(fmtStr) != cv.fmtFrom {
		err = errors.New("image format does not match")
	}
	return img, err
}

func (cv *Converter) encodeImg(w io.Writer, img image.Image) error {
	switch cv.fmtTo {
	case "jpeg":
		if err := jpeg.Encode(w, img, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(w, img); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(w, img, nil); err != nil {
			return err
		}
	}
	return nil
}

func (cv *Converter) generateOutputPath(path string) string {
	dirAndBase := strings.TrimRight(path, filepath.Ext(path))
	ext := cv.imgFmtExts.ConvToExt(cv.fmtTo)
	return strings.Join([]string{dirAndBase, string(ext)}, ".")
}

// Detect specifies image format from file extension string.
func (imgFmt *ImgFmt) Detect(cv *Converter, extStr string) {
	ext := Ext(strings.ToLower(extStr))
	*imgFmt = cv.imgFmtExts.ConvToImgFmt(ext)
}

// Match checks whether the file has an extension of the image format.
func (imgFmt ImgFmt) Match(cv *Converter, fileName string) bool {
	fileExtStr := strings.TrimPrefix(filepath.Ext(fileName), ".")
	fileExt := Ext(strings.ToLower(fileExtStr))
	fileImgFmt := cv.imgFmtExts.ConvToImgFmt(fileExt)
	return fileImgFmt == imgFmt
}

// Init creates the map of image formats and its extensions available.
func (m *MapImgFmtExts) Init() {
	*m = MapImgFmtExts{
		"jpeg": Exts{"jpg", "jpeg"},
		"png":  Exts{"png"},
		"gif":  Exts{"gif"},
	}
}

// ConvToImgFmt converts image extension to its format.
func (m MapImgFmtExts) ConvToImgFmt(ext Ext) ImgFmt {
	for imgFmt, fmtExts := range m {
		for _, fmtExt := range fmtExts {
			if ext == fmtExt {
				return imgFmt
			}
		}
	}
	return ""
}

// ConvToExt converts image format to its extension.
func (m MapImgFmtExts) ConvToExt(imgFmt ImgFmt) Ext {
	for keyImgFmt, fmtExts := range m {
		if imgFmt == keyImgFmt {
			return fmtExts[0]
		}
	}
	return ""
}
