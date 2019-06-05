package ifconv

import (
	"image"
	"image/jpeg"
	"os"
)

func Execute(dir string, before_f string, after_f string) error {
	sf, err := os.Open("test1.png")
	if err != nil {
		return err
	}
	defer sf.Close()

	img, format, err := image.Decode(sf)
	if err != nil {
		return err
	}

	if format == "hoge" {
		return err
	}

	file, err := os.Create("test1.jpeg")
	if err != nil {
		return err
	}
	defer file.Close()

	jpeg.Encode(file, img, &jpeg.Options{})
	return nil
}
