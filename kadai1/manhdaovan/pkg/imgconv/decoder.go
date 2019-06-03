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

type supportDecoders struct {
	mu   sync.Mutex
	decs map[ImgType]Decoder
}

var decoders = supportDecoders{
	decs: map[ImgType]Decoder{
		ImgTypePNG:  pngDecoder{},
		ImgTypeJPEG: jpegDecoder{},
		ImgTypeGIF:  gifDecoder{},
	},
}

type pngDecoder struct{}
type jpegDecoder struct{}
type gifDecoder struct{}

// Decode is a wrapper of png.Decode method
func (pd pngDecoder) Decode(r io.Reader) (pkgimg.Image, error) {
	return png.Decode(r)
}

// Decode is a wrapper of jpeg.Decode method
func (jd jpegDecoder) Decode(r io.Reader) (pkgimg.Image, error) {
	return jpeg.Decode(r)
}

// Decode is a wrapper of gif.Decode method
func (gd gifDecoder) Decode(r io.Reader) (pkgimg.Image, error) {
	return gif.Decode(r)
}

// GetDecoder returns Decoder associated with given imgType
func GetDecoder(imgType ImgType) Decoder {
	decoders.mu.Lock()
	defer decoders.mu.Unlock()
	return decoders.decs[imgType]
}

func registerNewDecoder(imgType ImgType, dec Decoder) error {
	decoders.mu.Lock()
	defer decoders.mu.Unlock()

	if d, ok := decoders.decs[imgType]; ok && d != nil {
		return fmt.Errorf("decoder for %s already registered", imgType)
	}

	decoders.decs[imgType] = dec
	return nil
}
