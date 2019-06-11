package converter

import (
	"os"
	"testing"
)

func fileAssertionHelper(t *testing.T, testCase string, file string) {
	t.Helper()
	// 変換後のファイルが存在するかチェック
	info, err := os.Stat(file)
	if err != nil {
		t.Errorf("Case (%s) Failed test: File not generated",
			testCase)
	}
	//　変換後のファイルサイズが0バイトではないかチェック
	if info.Size() <= 0 {
		t.Errorf("Case (%s) Failed test: Encoded file is invalid",
			testCase)
	}

}

func Test_Convert(t *testing.T) {
	cases := []struct {
		testCase   string
		decodeFile string
		encodeFile string
		decodeType string
		encodeType string
	}{
		{testCase: "J2P", decodeFile: "../testdata/cat.jpg", encodeFile: "../testdata/cat.png", decodeType: "jpg", encodeType: "png"},
		{testCase: "J2G", decodeFile: "../testdata/cat.jpg", encodeFile: "../testdata/cat.gif", decodeType: "jpg", encodeType: "gif"},
		{testCase: "P2J", decodeFile: "../testdata/cat.png", encodeFile: "../testdata/cat.jpg", decodeType: "png", encodeType: "jpg"},
		{testCase: "P2G", decodeFile: "../testdata/cat.png", encodeFile: "../testdata/cat.gif", decodeType: "png", encodeType: "gif"},
		{testCase: "G2J", decodeFile: "../testdata/cat.gif", encodeFile: "../testdata/cat.jpg", decodeType: "gif", encodeType: "jpg"},
		{testCase: "G2P", decodeFile: "../testdata/cat.gif", encodeFile: "../testdata/cat.png", decodeType: "gif", encodeType: "png"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.testCase, func(t *testing.T) {
			err := Convert(c.decodeFile, c.decodeType, c.encodeType)
			if err != nil {
				t.Errorf(
					"Case (%s) Failed test: Convert error.",
					c.testCase)
			}
			fileAssertionHelper(t, c.testCase, c.encodeFile)
		})
	}
}
