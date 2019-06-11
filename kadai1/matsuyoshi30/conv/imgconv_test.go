package conv

import (
	"os"
	"testing"
)

const (
	testfile1 = "../test/test1/appenginegophercolor.jpg"
	testfile2 = "../test/test1/appenginelogo.gif"
	testfile3 = "../test/test1/bumper.png"
	testfile4 = "../test/test2/dummy.jpeg"

	outfile1 = "../test/test1/appenginegophercolor.png" // jpg -> png
	outfile2 = "../test/test1/appenginelogo.jpeg"       // gif -> jpg
	outfile3 = "../test/test1/bumper.gif"               // png -> gif
)

func clean() {
	os.Remove(outfile1)
	os.Remove(outfile2)
	os.Remove(outfile3)
}

func TestImgconv(t *testing.T) {
	testImgconv_pass(t)
	testImgconv_fail(t)
}

func testImgconv_pass(t *testing.T) {
	if err := Imgconv(JPEG, PNG, testfile1); err != nil {
		t.Errorf("jpeg -> png: %v", err)
	}
	if err := Imgconv(GIF, JPEG, testfile2); err != nil {
		t.Errorf("gif -> jpeg: %v", err)
	}
	if err := Imgconv(PNG, GIF, testfile3); err != nil {
		t.Errorf("png -> gif: %v", err)
	}

	clean()
}

func testImgconv_fail(t *testing.T) {
	if err := Imgconv(JPEG, PNG, testfile4); err == nil {
		t.Errorf("Expected error")
	}
}
