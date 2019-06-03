package conv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// ImageType 画像の拡張子
type ImageType string

const (
	JPEG ImageType = "jpeg"
	JPG  ImageType = "jpg"
	PNG  ImageType = "png"
	GIF  ImageType = "gif"
)

// Imgconv dirpath 配下にある、bf に指定されたフォーマットの画像を、af に指定されたフォーマットの画像に変換する
func Imgconv(bf, af ImageType, dirpath string) error {
	filelist := make([]string, 0)
	err := filepath.Walk(dirpath, func(fp string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			ext := ImageType(filepath.Ext(fp))
			ext = ext[1:]

			if ext == bf {
				filelist = append(filelist, fp)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return imgconv(bf, af, filelist)
}

// imgconv filelist内の、bf に指定されたフォーマットの画像を、af に指定されたフォーマットの画像に変換する
func imgconv(bf, af ImageType, filelist []string) error {
	for _, f := range filelist {
		fmt.Printf("INPUT: %s", filepath.Base(f))
		img, err := decoder(f, bf)
		if err != nil {
			return err
		}

		dir, fn := filepath.Split(f)
		of := filepath.Base(fn[:len(fn)-len(filepath.Ext(fn))])
		outfile := dir + of

		err = encoder(img, outfile, af)
		if err != nil {
			return err
		}
		fmt.Printf(" => OUTPUT: %s.%s\n", of, af)
	}

	return nil
}

// decoder filename というファイル（フォーマットが format ）をデコードして image.Image を返す
func decoder(filename string, format ImageType) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// TODO: format とファイルの内容が一致しているか
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// encoder img という image.Image を、filename というファイル（フォーマットが format ）にエンコードして出力する
func encoder(img image.Image, filename string, format ImageType) error {
	out, err := os.Create(fmt.Sprintf("%s.%s", filename, format))
	if err != nil {
		return err
	}
	defer out.Close()

	switch format {
	case JPEG:
	case JPG:
		convertToJpeg(out, img)
	case PNG:
		convertToPng(out, img)
	case GIF:
		convertToGif(out, img)
	default:
		fmt.Println("Unknown format")
	}

	return nil
}

func convertToPng(out *os.File, img image.Image) {
	png.Encode(out, img)
}

func convertToJpeg(out *os.File, img image.Image) {
	jpeg.Encode(out, img, nil)
}

func convertToGif(out *os.File, img image.Image) {
	gif.Encode(out, img, nil)
}

func logError(err error, stop bool) {
	if err != nil {
		log.Fatal(err)
		if stop {
			os.Exit(1)
		}
	}
}
