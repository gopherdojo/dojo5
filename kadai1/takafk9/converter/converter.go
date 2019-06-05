package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Managed interface represents the error is properly managed and expected
type Managed interface {
	Managed() bool
}

type ManagedError struct {
	Message string
}

func (m *ManagedError) Error() string {
	return m.Message
}

func (m *ManagedError) Managed()bool{
	return true
}

func new()Managed{
	return &ManagedError{
		"hoge",
	}
}

func Convert(from, to, path string) ([]string, error) {
	filePaths, err := findFilePaths(from, path)
	if err != nil {
		return nil, err
	}

	for _, filePath := range filePaths {
		if err := convertImg(to, filePath); err != nil {
			return nil, err
		}
	}

	return filePaths, nil
}

func findFilePaths(from, path string) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if _, ok := err.(*os.PathError); ok {
			return &ManagedError{err.Error()}
		}

		if err != nil {
			return err
		}

		if filepath.Ext(path) == from {
			filePaths = append(filePaths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(filePaths) == 0 {
		return nil, &ManagedError{fmt.Sprintf("could not find files with the specified extension. path: %s, extension: %s", path, from)}
	}

	return filePaths, nil

}

func convertImg(to, filePath string) error {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)

	if err != nil {
		return err
	}

	out, err := os.Create(makeFilePath(to, filePath))

	if err != nil {
		return err
	}

	switch to {
	case ".gif":
		return gif.Encode(out, img, nil)
	case ".jpeg", ".jpg":
		return jpeg.Encode(out, img, nil)
	case ".png":
		return png.Encode(out, img)
	default:
		return fmt.Errorf("unsupported extension is specified: %s", to)
	}
}

func makeFilePath(to, filePath string) string {
	ext := filepath.Ext(filePath)
	return filePath[:len(filePath)-len(ext)] + to
}
