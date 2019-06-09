package imgconv

import (
	pkgimg "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

// PNGDecoder is default decoder for PNG
type PNGDecoder struct{}

// JPEGDecoder is default decoder for JPEG
type JPEGDecoder struct{}

// GIFDecoder is default decoder for GIF
type GIFDecoder struct{}

// Decode is a wrapper of png.Decode method
func (pd PNGDecoder) Decode(r io.Reader) (pkgimg.Image, error) {
	return png.Decode(r)
}

// Decode is a wrapper of jpeg.Decode method
func (jd JPEGDecoder) Decode(r io.Reader) (pkgimg.Image, error) {
	return jpeg.Decode(r)
}

// Decode is a wrapper of gif.Decode method
func (gd GIFDecoder) Decode(r io.Reader) (pkgimg.Image, error) {
	return gif.Decode(r)
}
