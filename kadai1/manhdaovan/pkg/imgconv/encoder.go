package imgconv

import (
	"fmt"
	pkgimg "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"sync"
)

type supportEncoders struct {
	mu   sync.Mutex
	encs map[ImgType]Encoder
}

var encoders = supportEncoders{
	encs: map[ImgType]Encoder{
		ImgTypePNG:  pngEncoder{},
		ImgTypeJPEG: jpegEncoder{},
		ImgTypeGIF:  gifEncoder{},
	},
}

type pngEncoder struct{}
type jpegEncoder struct{}
type gifEncoder struct{}

// Encode is a wrapper of png.Encode method
func (pe pngEncoder) Encode(w io.Writer, m pkgimg.Image) error {
	return png.Encode(w, m)
}

// Encode is a wrapper of jpeg.Encode method
func (je jpegEncoder) Encode(w io.Writer, m pkgimg.Image) error {
	return jpeg.Encode(w, m, nil)
}

// Encode is a wrapper of gif.Encode method
func (gifEncoder) Encode(w io.Writer, m pkgimg.Image) error {
	return gif.Encode(w, m, nil)
}

// GetEncoder returns Encoder associated with given imgType
func GetEncoder(imgType ImgType) Encoder {
	extensions.mu.Lock()
	defer extensions.mu.Unlock()

	return encoders.encs[imgType]
}

func registerNewEncoder(imgType ImgType, enc Encoder) error {
	encoders.mu.Lock()
	defer encoders.mu.Unlock()
	if enc, ok := encoders.encs[imgType]; ok && enc != nil {
		return fmt.Errorf("encoder for %s already registered", imgType)
	}
	encoders.encs[imgType] = enc
	return nil
}
