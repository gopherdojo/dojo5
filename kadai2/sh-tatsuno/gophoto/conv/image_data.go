package conv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// ImageData image struct
type ImageData struct {
	Path string
	Data image.Image
}

// NewImageData generate ImageData struct
func NewImageData(path string) (*ImageData, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return &ImageData{
		Path: path,
		Data: image,
	}, nil
}

// Convert convert image
func (i *ImageData) Convert(ext string) error {

	// check path
	if filepath.Ext(i.Path) == ext {
		return nil
	}

	// save file
	newPath := i.Path[:len(i.Path)-len(filepath.Ext(i.Path))] + ext

	// if file exist, do nothing
	if _, err := os.Stat(newPath); !os.IsNotExist(err) {
		return nil
	}

	i.Path = newPath
	if err := func(ext string) error {
		for _, suffix := range []string{".jpeg", ".jpg", ".gif", ".png"} {
			if ext == suffix {
				return nil
			}
		}
		return xerrors.New("invalid extension")
	}(ext); err != nil {
		return err
	}

	// if file exist, do nothing
	if _, err := os.Stat(i.Path); !os.IsNotExist(err) {
		return nil
	}

	dst, err := os.Create(i.Path)
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
	if err != nil {
		return err
	}
	return nil

	//// remove old file -> helper
	// if err := os.Remove(i.Path); err != nil {
	// 	return err
	// }

	return nil
}
