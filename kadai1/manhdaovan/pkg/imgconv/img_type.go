package imgconv

import (
	"fmt"
	"sync"
)

// ImgType represents image type option
// that this converter support
type ImgType string

const (
	// ImgTypeJPEG is constant for JPEG image
	ImgTypeJPEG ImgType = "jpeg"
	// ImgTypePNG is constant for PNG image
	ImgTypePNG ImgType = "png"
	// ImgTypeGIF is constant for GIF image
	ImgTypeGIF ImgType = "gif"
)

// SupportSrcImgTypes is a list of ImgType
type SupportSrcImgTypes []ImgType

// IsSupport returns imgType is supported by this tool or not
func (ssits SupportSrcImgTypes) IsSupport(imgType ImgType) bool {
	for _, it := range ssits {
		if it == imgType {
			return true
		}
	}

	return false
}

// GetSupportSrcImgTypes returns all supported types
// of input and output image
func GetSupportSrcImgTypes() SupportSrcImgTypes {
	imgTypes.mu.Lock()
	defer imgTypes.mu.Unlock()

	return imgTypes.types
}

type supportTypes struct {
	mu    sync.Mutex
	types SupportSrcImgTypes
}

var imgTypes = supportTypes{
	types: SupportSrcImgTypes{
		ImgTypeJPEG,
		ImgTypePNG,
		ImgTypeGIF,
	},
}

func registerImgType(its ...ImgType) error {
	imgTypes.mu.Lock()
	defer imgTypes.mu.Unlock()

	for _, it := range its {
		if imgTypes.types.IsSupport(it) {
			return fmt.Errorf("image type %s already registered", it)
		}
		imgTypes.types = append(imgTypes.types, it)
	}

	return nil
}
