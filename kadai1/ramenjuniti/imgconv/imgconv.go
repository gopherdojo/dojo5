package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// ImgData デコードされた画像データが入る
type ImgData struct {
	Path string
	Data image.Image
}

// Decode pathの画像をデコードし、ImgDataに入れる
func Decode(path string) (*ImgData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return &ImgData{Path: path, Data: img}, nil
}

// Encode iのDataをout形式のファイルにエンコードし、ファイル作成
// jpg, png, gif, tif, bmpに対応
func (i *ImgData) Encode(out string) error {
	file, err := os.Create(rmExt(i.Path) + "." + out)
	if err != nil {
		return err
	}

	switch out {
	case "png":
		err = png.Encode(file, i.Data)
	case "jpg", "jpeg":
		err = jpeg.Encode(file, i.Data, nil)
	case "gif":
		err = gif.Encode(file, i.Data, nil)
	case "tif", "tiff":
		err = tiff.Encode(file, i.Data, nil)
	case "bmp":
		err = bmp.Encode(file, i.Data)
	}

	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

// rmExt pathの拡張子を除いた文字列を返す
func rmExt(path string) string {
	return path[:len(path)-len(filepath.Ext(path))]
}
