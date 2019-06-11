package conv

import (
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func Test_Img(t *testing.T) {
	t.Run("OK: .png -> .jpeg", func(t *testing.T) {

		// ### Given ###
		img, err := NewImg("./lena.jpeg")
		if err != nil {
			t.Fatalf("Cannot load file. err: %v", err)
		}

		// ### When ###
		path := "./lena-png.jpeg"
		if err = img.Save(path); err != nil {
			t.Fatalf("Cannot save file. err: %v", err)
		}

		// ### Then ###
		file, err := os.Open(path)
		if err != nil {
			t.Fatalf("Cannot open saved file. err: %v", err)
		}
		if _, err = jpeg.Decode(file); err != nil {
			t.Fatalf("Cannot decode saved file. err: %v", err)
		}
		if err := os.Remove(path); err != nil {
			t.Fatalf("Cannot remove saved file. err: %v", err)
		}

	})

	t.Run("OK: .jpeg -> .png", func(t *testing.T) {

		// ### Given ###
		img, err := NewImg("./lena.png")
		if err != nil {
			t.Fatalf("Cannot load file. err: %v", err)
		}

		// ### When ###
		path := "./lena-jpeg.png"
		if err = img.Save(path); err != nil {
			t.Fatalf("Cannot save file. err: %v", err)
		}

		// ### Then ###
		file, err := os.Open(path)
		if err != nil {
			t.Fatalf("Cannot open saved file. err: %v", err)
		}
		if _, err = png.Decode(file); err != nil {
			t.Fatalf("Cannot decode saved file. err: %v", err)
		}
		if err := os.Remove(path); err != nil {
			t.Fatalf("Cannot remove saved file. err: %v", err)
		}

	})

	t.Run("NG: cannot openfile", func(t *testing.T) {

		// ### Given ###
		_, err := NewImg("./noexist.png")
		if err == nil {
			t.Fatal("Load no exist file.")
		}
	})

	t.Run("NG: invalid extension", func(t *testing.T) {

		// ### Given ###
		_, err := NewImg("./img.go")
		if err == nil {
			t.Fatal("Load invalid file.")
		}
	})

}
