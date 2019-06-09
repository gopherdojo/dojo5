package imgconv

const (
	// ImgTypeJPEG is constant for JPEG image
	ImgTypeJPEG ImgType = "jpeg"
	// ImgTypePNG is constant for PNG image
	ImgTypePNG ImgType = "png"
	// ImgTypeGIF is constant for GIF image
	ImgTypeGIF ImgType = "gif"

	// ImgExtJPEG is constant for extension of JPEG image
	ImgExtJPEG ImgExt = "jpg"
	// ImgExtPNG is constant for extension of PNG image
	ImgExtPNG ImgExt = "png"
	// ImgExtGIF is constant for extension of GIF image
	ImgExtGIF ImgExt = "gif"
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

// GetSupportSrcImgTypes returns the image types that supported by default
func GetSupportSrcImgTypes() SupportSrcImgTypes {
	return SupportSrcImgTypes{
		ImgTypeJPEG,
		ImgTypePNG,
		ImgTypeGIF,
	}
}
