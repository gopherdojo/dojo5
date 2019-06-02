// imgconverter パッケージは指定の画像ファイルを指定の形式に変換する機能を提供します
package imgconverter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
)

type Client struct {

	// tgtExt：変換先のファイル形式
	tgtExt string
	// savePath：保存先ディレクトリ、存在しない場合新たに作成されます
	savePath string
	// 変換対象の画像リスト
	images []string
}

func New(tgtExt, savePath string, images []string) *Client {
	return &Client{

		tgtExt:   tgtExt,
		savePath: savePath,
		images:   images,
	}
}

func (c *Client) Convert() {

	if _, err := os.Stat(c.savePath); os.IsNotExist(err) {
		os.MkdirAll(c.savePath, 0755)
	}

	for _, image := range c.images {
		err := convertImage(image, c.tgtExt, c.savePath)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}

//一枚の画像を特定のフォーマットに変換し指定のディレクトリに保存します
func convertImage(imgPath, tgtExt, savePath string) error {

	if tgtExt == filepath.Ext(imgPath) {
		err := fmt.Errorf("file \"%s\"'s extension is already %s, conversion skipped", imgPath, tgtExt)
		return err
	}

	src, err := os.Open(imgPath)
	if err != nil {

		return err
	}
	defer src.Close()

	decoded, err := decode(src)
	if err != nil {
		return err
	}

	name := imgPath[0 : len(imgPath)-len(filepath.Ext(imgPath))]
	name = filepath.Base(name)

	tgt, err := os.Create(savePath + "/" + name + tgtExt)
	if err != nil {
		return err
	}
	defer tgt.Close()

	err = encode(tgt, decoded)

	return err
}

//画像のデコードを行います
func decode(srcImg *os.File) (decodedImg image.Image, err error) {
	name := srcImg.Name()
	ext := filepath.Ext(name)
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		decodedImg, err = jpeg.Decode(srcImg)
	case ".png":
		decodedImg, err = png.Decode(srcImg)
	case ".gif":
		decodedImg, err = gif.Decode(srcImg)
	case ".bmp":
		decodedImg, err = bmp.Decode(srcImg)
	}

	return
}

//特定の形式へのエンコードを行います
func encode(tgt *os.File, decodedImg image.Image) (err error) {
	name := tgt.Name()
	ext := filepath.Ext(name)
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		opt := &jpeg.Options{Quality: 100}
		err = jpeg.Encode(tgt, decodedImg, opt)
	case ".png":
		err = png.Encode(tgt, decodedImg)
	case ".gif":
		opt := &gif.Options{}
		err = gif.Encode(tgt, decodedImg, opt)
	case ".bmp":
		err = bmp.Encode(tgt, decodedImg)
	}

	return
}
