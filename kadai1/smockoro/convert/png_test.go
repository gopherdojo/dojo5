package convert

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestNewPngConverter(t *testing.T) {
	c := NewPngConverter()
	expected := reflect.TypeOf(&pngConverter{})
	actual := reflect.TypeOf(c)
	if expected != actual {
		t.Errorf("Expected type is %v but actual type is %v", expected, actual)
	}
	if _, ok := c.(Converter); !ok {
		t.Errorf("Not Implement Converter")
	}
}

func TestPngImageConvert(t *testing.T) {
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

	c := NewPngConverter()
	for _, fpath := range filepaths {
		t.Run(fpath, func(t *testing.T) {
			if err := c.ImageConvert(fpath); err != nil {
				t.Errorf("Expected err is nil but actual err is not nil: %v\n", err)
			}
			newfpath := strings.TrimSuffix(fpath, filepath.Ext(fpath)) + ".png"
			if _, err := os.Stat(newfpath); os.IsNotExist(err) {
				t.Errorf("Expected %s is exested but actual not exitst", newfpath)
			}
		})
	}

	TestFileCleanup(t)
}
