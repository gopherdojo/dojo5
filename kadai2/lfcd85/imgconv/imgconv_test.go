package imgconv

import (
	"os"
	"testing"
)

func assertEq(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
}

func assertNil(t *testing.T, obj interface{}) {
	if obj != nil {
		t.Errorf("actual: not nil, expected: nil")
	}
}

func assertNotNil(t *testing.T, obj interface{}) {
	if obj == nil {
		t.Errorf("actual: nil, expected: not nil")
	}
}

func TestConvert(t *testing.T) {
	cases := []struct {
		from     string
		to       string
		expected bool
	}{
		{"jpeg", "png", true},
		{"png", "gif", true},
		{"gif", "jpeg", true},
		{"jpeg", "jpeg", false},
		{"rb", "go", false},
	}

	for _, c := range cases {
		defer os.RemoveAll("./output")

		err := Convert("../testdata", c.from, c.to)
		if c.expected == true {
			assertNil(t, err)
		} else {
			assertNotNil(t, err)
		}
	}
}

func TestGenerateOutputPath(t *testing.T) {
	cases := []struct {
		path     string
		fmtTo    ImgFmt
		expected string
	}{
		{
			"path/to/hoge.jpg",
			ImgFmt("png"),
			"path/to/hoge.png",
		},
		{
			"./path/to/fuga.PNG",
			ImgFmt("jpeg"),
			"./path/to/fuga.jpg",
		},
		{
			"piyo.png",
			ImgFmt("gif"),
			"piyo.gif",
		},
	}

	imgFmtExts.Init()
	for _, c := range cases {
		fmtTo = c.fmtTo
		assertEq(t, generateOutputPath(c.path), c.expected)
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

	imgFmtExts.Init()
	var imgFmt ImgFmt
	for _, c := range cases {
		imgFmt.Detect(c.extStr)
		assertEq(t, imgFmt, c.expected)
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

	imgFmtExts.Init()
	for _, c := range cases {
		assertEq(t, c.imgFmt.Match(c.fileName), c.expected)
	}
}
