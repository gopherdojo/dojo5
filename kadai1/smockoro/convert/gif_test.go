package convert

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestFileCleanup(t *testing.T) {
	t.Helper()
	if _, err := os.Stat("./testdata/"); !os.IsNotExist(err) {
		err := os.RemoveAll("./testdata/")
		if err != nil {
			t.Fatal("can't clean test data")
		}
	}
}

func TestNewGifConverter(t *testing.T) {
	c := NewGifConverter()
	expected := reflect.TypeOf(&gifConverter{})
	actual := reflect.TypeOf(c)
	if expected != actual {
		t.Errorf("Expected type is %v but actual type is %v", expected, actual)
	}
	if _, ok := c.(Converter); !ok {
		t.Errorf("Not Implement Converter")
	}
}

func TestGifImageConvert(t *testing.T) {
	TestFileCleanup(t)
	/*
	   testdata
	   |- png
	   |   |- gopher.png
	   |- jpg
	   |   |- gopher.jpg
	   |- jpeg
	   |   |- gopher.jpeg
	   |- gif
	       |- gopher.gif
	*/
	dirpaths := []string{
		"./testdata/png",
		"./testdata/jpg",
		"./testdata/jpeg",
		"./testdata/gif",
	}
	dirmaker := NewDirmaker(dirpaths)
	dirmaker.Make()
	if dirmaker.err != nil {
		t.Fatal(dirmaker.err.Error())
	}
	filepaths := []string{
		"./testdata/png/gopher.png",
		"./testdata/jpg/gopher.jpg",
		"./testdata/jpeg/gopher.jpeg",
		"./testdata/gif/gopher.gif",
	}
	filemaker := NewFilemaker(filepaths)
	filemaker.Make()
	if filemaker.err != nil {
		t.Fatal(filemaker.err.Error())
	}

	c := NewGifConverter()
	for _, fpath := range filepaths {
		t.Run(fpath, func(t *testing.T) {
			if err := c.ImageConvert(fpath); err != nil {
				t.Errorf("Expected err is nil but actual err is not nil: %v\n", err)
			}
			newfpath := strings.TrimSuffix(fpath, filepath.Ext(fpath)) + ".gif"
			if _, err := os.Stat(newfpath); os.IsNotExist(err) {
				t.Errorf("Expected %s is exested but actual not exitst", newfpath)
			}
		})
	}

	filepaths = []string{"./testdata/gopher.txt"}
	filemaker.Make()
	filepaths = append(filepaths, "./aaaaaaa/gopher.png")
	if filemaker.err != nil {
		t.Fatal(filemaker.err.Error())
	}
	for _, fpath := range filepaths {
		t.Run(fpath, func(t *testing.T) {
			if err := c.ImageConvert(fpath); err == nil {
				t.Errorf("Expected err is not nil but actual err is nil")
			}
			newfpath := strings.TrimSuffix(fpath, filepath.Ext(fpath)) + ".gif"
			if _, err := os.Stat(newfpath); !os.IsNotExist(err) {
				t.Errorf("Expected %s is not exested but actual exitst", newfpath)
			}
		})
	}

	TestFileCleanup(t)
}

type Filemaker struct {
	filepaths []string
	err       error
}

func NewFilemaker(filepaths []string) *Filemaker {
	return &Filemaker{
		filepaths: filepaths,
		err:       nil,
	}
}

func (f *Filemaker) Make() bool {
	for _, fpath := range f.filepaths {
		img := image.NewRGBA(image.Rect(0, 0, 100, 100))
		/*
			□□□□
			□■■□
			□■■□
			□□□□
		*/
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				switch {
				case x > 25 && x < 75 && y > 25 && y < 75:
					img.Set(x, y, color.Black)
				default:
					img.Set(x, y, color.White)
				}
			}
		}
		file, err := os.Create(fpath)
		if err != nil {
			f.err = err
			return false
		}
		defer file.Close()
		switch filepath.Ext(fpath) {
		case ".jpg":
			jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
		case ".jpeg":
			jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
		case ".png":
			png.Encode(file, img)
		case ".gif":
			gif.Encode(file, img, &gif.Options{NumColors: 256})
		default:
		}
	}
	return false
}

type DirMaker struct {
	dirpaths []string
	err      error
}

func NewDirmaker(dirpaths []string) *DirMaker {
	return &DirMaker{
		dirpaths: dirpaths,
		err:      nil,
	}
}

func (d *DirMaker) Make() bool {
	for _, dirpath := range d.dirpaths {
		err := os.MkdirAll(dirpath, 0777)
		if err != nil {
			d.err = err
			return false
		}
	}
	return false
}
