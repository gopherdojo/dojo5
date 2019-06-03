package imgconv

import (
	"fmt"
	"sync"
)

// ImgExt represents the extension of image type
// Eg: PNG file has "png" extension, and "jpg" for JPEG
type ImgExt string
type supportExtensions struct {
	mu   sync.Mutex
	exts map[ImgType]ImgExt
}

var extensions = supportExtensions{
	exts: map[ImgType]ImgExt{
		ImgTypeJPEG: "jpg",
		ImgTypeGIF:  "gif",
		ImgTypePNG:  "png",
	},
}

func registerNewExt(imgType ImgType, ext ImgExt) error {
	extensions.mu.Lock()
	defer extensions.mu.Unlock()
	if ex, ok := extensions.exts[imgType]; ok && ex != "" {
		return fmt.Errorf("extension for %s already registered", imgType)
	}
	extensions.exts[imgType] = ext
	return nil
}

// GetExtension returns file extension associated with given imgType
func GetExtension(imgType ImgType) ImgExt {
	extensions.mu.Lock()
	defer extensions.mu.Unlock()
	return extensions.exts[imgType]
}
