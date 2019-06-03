package conv

import (
	"strings"
)

// ImgExt is an enum type for image extensions.
type ImgExt int

const (
	// ImgExtUNDEF indicates an undefined image extension
	ImgExtUNDEF ImgExt = iota
	// ImgExtJPEG indicates the JPEG image extension
	ImgExtJPEG
	// ImgExtPNG indicates the PNG image extension
	ImgExtPNG
	// ImgExtTIFF indicates the TIFF image extension
	ImgExtTIFF
	// ImgExtBMP indicates the BMP image extension
	ImgExtBMP
)

func (e ImgExt) String() string {
	switch e {
	case ImgExtJPEG:
		return ".jpg"
	case ImgExtPNG:
		return ".png"
	case ImgExtTIFF:
		return ".tiff"
	case ImgExtBMP:
		return ".bmp"
	default:
		return "undefined"
	}
}

// ParseImgExt returns ImgExt from a file name
func ParseImgExt(s string) ImgExt {
	sl := strings.Split(s, ".")
	ext := sl[len(sl)-1]
	ext = strings.TrimSpace(ext)
	ext = strings.ToLower(ext)
	switch ext {
	case "jpg", "jpeg":
		return ImgExtJPEG
	case "png":
		return ImgExtPNG
	case "tiff":
		return ImgExtTIFF
	case "bmp":
		return ImgExtBMP
	default:
		return ImgExtUNDEF
	}
}
