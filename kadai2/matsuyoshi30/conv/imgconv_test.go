package conv_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matsuyoshi30/dojo5/kadai2/matsuyoshi30/conv"
)

const TESTDIR = "../testdata/"

var (
	normal = filepath.Join(TESTDIR, "normal")
	gr     = filepath.Join(TESTDIR, "gr")
	dummy  = filepath.Join(TESTDIR, "dummy")
)

func clean(filename string) {
	os.Remove(filename)
}

func TestImgconv(t *testing.T) {
	testcases := []struct {
		name       string
		testtype   string
		inputtype  conv.ImageType
		outputtype conv.ImageType
		inputdir   string
		output     string
	}{
		{"test1", "SUCCESS", conv.JPEG, conv.PNG, normal, "appenginegophercolor.png"},
		{"test2", "SUCCESS", conv.GIF, conv.JPEG, normal, "appenginelogo.jpeg"},
		{"test3", "SUCCESS", conv.PNG, conv.GIF, normal, "bumper.gif"},
		{"test4", "FAIL", conv.JPEG, conv.PNG, dummy, "dummy.png"},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.testtype {
			case "SUCCESS":
				testImgconv_pass(t, tt.inputtype, tt.outputtype,
					tt.inputdir, filepath.Join(tt.inputdir, tt.output))
			case "FAIL":
				testImgconv_fail(t, tt.inputtype, tt.outputtype, tt.inputdir)
			}
		})
	}
}

func testImgconv_pass(t *testing.T, from conv.ImageType, to conv.ImageType, dir string, output string) {
	t.Helper()
	if err := conv.Imgconv(from, to, dir); err != nil {
		t.Fatalf("DIR: %v (%v -> %v): %v", dir, from, to, err)
	}

	clean(output)
}

func testImgconv_fail(t *testing.T, from conv.ImageType, to conv.ImageType, dir string) {
	t.Helper()
	if err := conv.Imgconv(from, to, dir); err == nil {
		t.Fatal("Expected error")
	}
}
