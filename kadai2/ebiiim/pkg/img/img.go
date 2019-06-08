package img

import (
	"fmt"
	"strings"
)

// Ext is an enum type for image extensions.
type Ext string

const (
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
func ParseExt(s string) (Ext, error) {
	sl := strings.Split(s, ".")
	ext := sl[len(sl)-1]
	ext = strings.TrimSpace(ext)
	ext = strings.ToLower(ext)
	switch ext {
	case "jpg", "jpeg":
		return JPEG, nil
	case "png":
		return PNG, nil
	case "tif", "tiff":
		return TIFF, nil
	case "bmp":
		return BMP, nil
	default:
		return "", fmt.Errorf("unsupported image extension %s", s)
	}
}
