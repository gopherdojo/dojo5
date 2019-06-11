package imgconv

import (
	"os"
	"strings"
	"testing"
)

func assertEq(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
}

func assertFileExists(t *testing.T, filePath string, expected bool) {
	t.Helper()

	_, err := os.Stat(filePath)
	actual := err == nil

	if actual != expected {
		switch expected {
		case true:
			t.Errorf("file %v should exist but does not", filePath)
		case false:
			t.Errorf("file %v should not exist but does", filePath)
		}
	}
}

func TestConvert(t *testing.T) {
	cases := []struct {
		from              string
		to                string
		expectedSuccess   bool
		expectedFileNames map[string]bool
	}{
		{"jpeg", "png", true, map[string]bool{
			"not_image.png":         false,
			"not_image2.png":        false,
			"sample1.png":           true,
			"sample2.png":           true,
			"sample4.png":           false,
			"child_dir/sample3.png": true,
			"child_dir/sample4.png": false,
			"child_dir/sample5.png": false,
		}},
		{"png", "gif", true, map[string]bool{
			"not_image.gif":         false,
			"not_image2.gif":        false,
			"sample1.gif":           false,
			"sample2.gif":           false,
			"sample4.gif":           true,
			"child_dir/sample3.gif": false,
			"child_dir/sample4.gif": true,
			"child_dir/sample5.gif": false,
		}},
		{"gif", "jpeg", true, map[string]bool{
			"not_image.jpg":         false,
			"not_image2.jpg":        false,
			"sample1.jpg":           false,
			"sample2.jpg":           false,
			"sample4.jpg":           false,
			"child_dir/sample3.jpg": false,
			"child_dir/sample4.jpg": false,
			"child_dir/sample5.jpg": true,
		}},
		{"jpeg", "jpeg", false, nil},
		{"rb", "go", false, nil},
	}

	outputDir := "./output/testdata"
	for _, c := range cases {
		c := c
		t.Run(strings.Join([]string{c.from, c.to}, "->"), func(t *testing.T) {
			defer os.RemoveAll("./output")

			err := Convert("../testdata", c.from, c.to)
			if err != nil && c.expectedSuccess == true {
				t.Errorf("function Convert is expected to succeed, but actually failed")
			}
			if err == nil && c.expectedSuccess == false {
				t.Errorf("function Convert is expected to fail, but actually succeeded")
			}
			for f, b := range c.expectedFileNames {
				filePath := strings.Join([]string{outputDir, f}, "/")
				assertFileExists(t, filePath, b)
			}
		})
	}
}

func TestConverter_GenerateOutputPath(t *testing.T) {
	cases := []struct {
		path     string
		fmtTo    ImgFmt
		expected string
	}{
		{
			"path/to/hoge.jpg",
			ImgFmt("png"),
			"./output/path/to/hoge.png",
		},
		{
			"./path/to/fuga.PNG",
			ImgFmt("jpeg"),
			"./output/path/to/fuga.jpg",
		},
		{
			"piyo.png",
			ImgFmt("gif"),
			"./output/piyo.gif",
		},
		{
			"../../path/to/foobar.gif",
			ImgFmt("jpeg"),
			"./output/path/to/foobar.jpg",
		},
	}

	cv := &Converter{}
	cv.imgFmtExts.Init()
	for _, c := range cases {
		c := c
		t.Run(c.path, func(t *testing.T) {
			cv.fmtTo = c.fmtTo
			assertEq(t, cv.generateOutputPath(c.path), c.expected)
		})
	}
}

func TestImgFmt_Detect(t *testing.T) {
	cases := []struct {
		extStr   string
		expected ImgFmt
	}{
		{"png", ImgFmt("png")},
		{"jpg", ImgFmt("jpeg")},
		{"JPEG", ImgFmt("jpeg")},
		{"GIF", ImgFmt("gif")},
	}

	cv := &Converter{}
	cv.imgFmtExts.Init()
	for _, c := range cases {
		c := c
		t.Run(c.extStr, func(t *testing.T) {
			var imgFmt ImgFmt
			imgFmt.Detect(cv, c.extStr)
			assertEq(t, imgFmt, c.expected)
		})
	}
}

func TestImgFmt_Match(t *testing.T) {
	cases := []struct {
		fileName string
		imgFmt   ImgFmt
		expected bool
	}{
		{"hoge.jpg", ImgFmt("jpeg"), true},
		{"fuga.png", ImgFmt("gif"), false},
		{"piyo.png", ImgFmt("png"), true},
		{"foo.js", ImgFmt("png"), false},
		{".JPEG", ImgFmt("jpeg"), true},
		{"jpeg", ImgFmt("jpeg"), false},
		{"foopng", ImgFmt("png"), false},
	}

	cv := &Converter{}
	cv.imgFmtExts.Init()
	for _, c := range cases {
		c := c
		t.Run(c.fileName, func(t *testing.T) {
			assertEq(t, c.imgFmt.Match(cv, c.fileName), c.expected)
		})
	}
}
