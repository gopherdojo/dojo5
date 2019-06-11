/*
Package 'convert' provides image conversion
*/
package convert

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Convert is a structure required for conversion
type Convert struct {
	SrcPath  string
	FileName string
	DestDir  string
}

// New generates and returns a new Convert.
// An error is returned if the original directory for image conversion does not exist.
func New(srcPath string, destDir string) (*Convert, error) {
	stat, err := os.Stat(srcPath)
	if err != nil {
		return nil, err
	}

	return &Convert{
		SrcPath:  srcPath,
		FileName: stat.Name(),
		DestDir:  destDir,
	}, nil
}

// JpegOptions is an option to convert Jpeg
type JpegOptions jpeg.Options

// ToJpeg does Jpeg conversion. You can specify conversion options with JpegOptions.
func (c *Convert) ToJpeg(opt *JpegOptions) (string, error) {

	image, outFilePath, err := c.convertPreProcessing(".jpeg")
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(outFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		return "", err
	}
	defer f.Close()

	jo := jpeg.Options(*opt)
	return outFilePath, jpeg.Encode(f, *image, &jo)
}

// ToPng does Png conversion
func (c *Convert) ToPng() (string, error) {

	image, outFilePath, err := c.convertPreProcessing(".png")
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(outFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		return "", err
	}

	return outFilePath, png.Encode(f, *image)
}

func (c *Convert) convertPreProcessing(ext string) (*image.Image, string, error) {
	file, err := os.Open(c.SrcPath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	_, err = os.Stat(c.DestDir)
	if err != nil {
		if err := os.MkdirAll(c.DestDir, 0777); err != nil {
			return nil, "", err
		}
	}

	outFilePath := filepath.Join(c.DestDir, replaceExt(c.FileName, ext))

	return &image, outFilePath, nil
}

func replaceExt(filePath, to string) string {
	ext := filepath.Ext(filePath)
	return strings.Replace(filePath, ext, to, -1)
}
