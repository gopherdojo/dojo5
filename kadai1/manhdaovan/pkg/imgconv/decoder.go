package imgconv

import (
	pkgimg "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

var decoders = map[ImgType]decoder{
	ImgTypePNG:  pngDecoder{},
	ImgTypeJPEG: jpegDecoder{},
	ImgTypeGIF:  gifDecoder{},
}

type pngDecoder struct{}
type jpegDecoder struct{}
type gifDecoder struct{}

func (pd pngDecoder) decode(r io.Reader) (pkgimg.Image, error) {
	return png.Decode(r)
}

func (jd jpegDecoder) decode(r io.Reader) (pkgimg.Image, error) {
	return jpeg.Decode(r)
}

func (gd gifDecoder) decode(r io.Reader) (pkgimg.Image, error) {
	return gif.Decode(r)
}

func getDecoder(imgType ImgType) decoder {
	return decoders[imgType]
}
