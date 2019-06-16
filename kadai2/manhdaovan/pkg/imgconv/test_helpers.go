package imgconv

import (
	"fmt"
	pkgimg "image"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// this file contains variables, methods that would be shared between tests

var (
	testdataDir = "testdata"
	// orgFilesDir contains:
	// - textFile.txt
	// - textFileRenameToPNG.png
	// - validGIF.gif
	// - validJPEG.jpg
	// - validPNG.png
	orgFilesDir = testdataDir + "/orgFiles"
	orgFiles    = paths{
		orgFilesDir + "/textFile.txt",
		orgFilesDir + "/textFileRenameToPNG.png",
		orgFilesDir + "/validGIF.gif",
		orgFilesDir + "/validJPEG.jpg",
		orgFilesDir + "/validPNG.png",
	}
	rootForTestDir = testdataDir + "/rootForTest"
	subDir         = rootForTestDir + "/subdir"
)

type paths []string

// verifyImgs verifies that images are existing or not,
// and corresponding to it type
func verifyImgs(t *testing.T, imgs paths, imgType ImgType) {
	t.Helper()

	for _, img := range imgs {
		ok, err := isImgWithType(img, imgType)
		if err != nil {
			t.Error(err)
		}
		if !ok {
			t.Errorf("converted image type is wrong, img: %s, expect type: %s", img, imgType)
		}
	}
}

// verifyFiles verifies files are existing or not
func verifyFiles(t *testing.T, files paths, checkExisting bool) {
	t.Helper()

	for _, f := range files {
		file, err := os.Open(f)
		defer file.Close()
		switch checkExisting {
		case true:
			if err == nil {
				continue
			}
			t.Errorf("need file %v existing, got err: %+v", f, err)
		case false:
			if os.IsNotExist(err) {
				continue
			}
			t.Errorf("need file %v removed, got err: %v", f, err)
		}
	}
}

func copyTestFilesToDir(t *testing.T, files paths, dirs paths) {
	for _, dir := range dirs {
		for _, f := range files {
			srcFile, err := ioutil.ReadFile(f)
			if err != nil {
				t.Error(err)
			}

			fileNameIdx := strings.LastIndex(f, "/")
			destPath := dir + string(f[fileNameIdx:])
			err = ioutil.WriteFile(destPath, srcFile, 0644)
			if err != nil {
				t.Error(err)
			}
		}
	}
}

func deleteAllFiles(t *testing.T, dirs paths) {
	for _, dir := range dirs {
		outerErr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() { // file
				return os.Remove(path)
			}

			return nil
		})

		if outerErr != nil {
			t.Error(outerErr)
		}
	}
}

// isSameFiles returns `files1` is same with `files2` in unordered,
// when `files1` and `files2` do not contain duplicated files.
func isSameFiles(files1, files2 paths) bool {
	if files1 == nil && files2 == nil {
		return true
	}
	if len(files1) != len(files2) {
		return false
	}

	checkedFiles := make(map[string]bool)
	numChecked := 0
	numFiles := len(files1)

	for _, f1 := range files1 {
		checkedFiles[f1] = false
	}

	for _, f2 := range files2 {
		if checked, found := checkedFiles[f2]; checked || !found {
			return false // file already checked
		}
		checkedFiles[f2] = true
		numChecked++
	}

	if numChecked != numFiles {
		return false
	}

	return true
}

type errDecoder struct{}

func (errDecoder) Decode(r io.Reader) (pkgimg.Image, error) {
	return nil, fmt.Errorf("error on errDecoder.Decode")
}
