package imgconv

import (
	pkgimg "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

var encoders = map[ImgType]encoder{
	ImgTypePNG:  pngEncoder{},
	ImgTypeJPEG: jpegEncoder{},
	ImgTypeGIF:  gifEncoder{},
}

type pngEncoder struct{}
type jpegEncoder struct{}
type gifEncoder struct{}

func (pe pngEncoder) encode(w io.Writer, m pkgimg.Image) error {
	return png.Encode(w, m)
}

func (je jpegEncoder) encode(w io.Writer, m pkgimg.Image) error {
	return jpeg.Encode(w, m, nil)
}

func (gifEncoder) encode(w io.Writer, m pkgimg.Image) error {
	return gif.Encode(w, m, nil)
}

func getEncoder(imgType ImgType) encoder {
	return encoders[imgType]
}
