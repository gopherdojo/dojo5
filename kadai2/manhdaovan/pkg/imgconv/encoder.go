package imgconv

import (
	pkgimg "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

// PNGEncoder is default encoder for PNG
type PNGEncoder struct{}

// JPEGEncoder is default encoder for JPEG
type JPEGEncoder struct{}

// GIFEncoder is default encoder for GIF
type GIFEncoder struct{}

// Encode is a wrapper of png.Encode method
func (PNGEncoder) Encode(w io.Writer, m pkgimg.Image) error {
	return png.Encode(w, m)
}

// Encode is a wrapper of jpeg.Encode method
func (JPEGEncoder) Encode(w io.Writer, m pkgimg.Image) error {
	return jpeg.Encode(w, m, nil)
}

// Encode is a wrapper of gif.Encode method
func (GIFEncoder) Encode(w io.Writer, m pkgimg.Image) error {
	return gif.Encode(w, m, nil)
}
