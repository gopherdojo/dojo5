package img

import (
	"strings"
)

// Ext is an enum type for image extensions.
type Ext string

const (
	// UNDEF indicates an undefined image extension
	UNDEF Ext = "undefined"
	// JPEG indicates the JPEG image extension
	JPEG Ext = "jpg"
	// PNG indicates the PNG image extension
	PNG Ext = "png"
	// TIFF indicates the TIFF image extension
	TIFF Ext = "tif"
	// BMP indicates the BMP image extension
	BMP Ext = "bmp"
)

// ParseExt returns Ext from a file name
func ParseExt(s string) Ext {
	sl := strings.Split(s, ".")
	ext := sl[len(sl)-1]
	ext = strings.TrimSpace(ext)
	ext = strings.ToLower(ext)
	switch ext {
	case "jpg", "jpeg":
		return JPEG
	case "png":
		return PNG
	case "tif", "tiff":
		return TIFF
	case "bmp":
		return BMP
	default:
		return UNDEF
	}
}
