/*
Package converter

converter は特定の画像形式のファイルを変換するためのパッケージです。

How to use

string 型の3つの引数を指定して Convert 関数を呼び出してください。
 Convert(filePath string, decodeType string, encodeType string)

利用できるファイル形式は gif、jpg(jpeg)、png だけです。
*/
package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

//dec: Decode filePath file
func dec(filePath string, decodeType string) (image.Image, error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	switch decodeType {
	case "gif":
		return gif.Decode(reader)
	case "jpeg", "jpg":
		return jpeg.Decode(reader)
	case "png":
		return png.Decode(reader)
	}
	return nil, err
}

//enc: Encode m to encodeType
func enc(filePath string, encodeType string, m image.Image) error {
	writer, err := os.Create(strings.TrimSuffix(filePath, path.Ext(filePath)) + "." + encodeType)
	if err != nil {
		return err
	}
	defer writer.Close()

	switch encodeType {
	case "gif":
		return gif.Encode(writer, m, nil)
	case "jpeg", "jpg":
		return jpeg.Encode(writer, m, nil)
	case "png":
		return png.Encode(writer, m)
	}
	return nil
}

//Convert : convert decodeType file to encodeType file
func Convert(filePath string, decodeType string, encodeType string) error {
	fmt.Println("Converting: ", filePath)

	m, err := dec(filePath, decodeType)
	if err != nil {
		return err
	}

	err = enc(filePath, encodeType, m)
	if err != nil {
		return err
	}

	fmt.Println("Converted : ", strings.TrimSuffix(filePath, path.Ext(filePath)) + "." + encodeType)
	return nil
}
