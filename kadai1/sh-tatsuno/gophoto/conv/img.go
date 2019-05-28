package conv

import (
	"image"
	"image/gif"  // import gif
	"image/jpeg" // import jpeg
	"image/png"  // import png
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// Img image struct
type Img struct {
	Path string
	Data image.Image
}

// NewImg generate Img struct
func NewImg(path string) (*Img, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return &Img{
		Path: path,
		Data: img,
	}, nil
}

// Save save img for input format if format satisfies ".jpeg", ".jpg", ".gif", ".png"
func (i *Img) Save(path string) error {
	ext := filepath.Ext(path)
	var err error
	err = func(ext string) error {
		for _, postfix := range []string{".jpeg", ".jpg", ".gif", ".png"} {
			if ext == postfix {
				return nil
			}
		}
		return xerrors.New("invalid extension")
	}(ext)
	if err != nil {
		return err
	}

	// if file exist, do nothing
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	switch ext {
	case ".jpeg", ".jpg":
		err = jpeg.Encode(dst, i.Data, nil)
	case ".gif":
		err = gif.Encode(dst, i.Data, nil)
	case ".png":
		err = png.Encode(dst, i.Data)
	default:
		err = xerrors.New("error in main method")
	}
	return err
}

// AddExt add extension
func (i *Img) AddExt(ext string) string {
	return i.Path + ext
}

// Replace replace image
func (i *Img) Replace(ext string) error {

	// check path
	prevPath := i.Path
	if filepath.Ext(prevPath) == ext {
		return nil
	}

	// save file
	newPath := prevPath[:len(prevPath)-len(filepath.Ext(prevPath))] + ext

	// if file exist, do nothing
	if _, err := os.Stat(newPath); !os.IsNotExist(err) {
		return nil
	}

	if err := i.Save(newPath); err != nil {
		return err
	}

	// remove old file
	if err := os.Remove(prevPath); err != nil {
		return err
	}
	i.Path = newPath
	return nil
}
