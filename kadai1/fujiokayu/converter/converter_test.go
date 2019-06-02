package converter

import (
	"os"
	"testing"
)

const decodeFile string = "./cat.jpg"
const encodeFile string = "./cat.png"

func Test_Convert(t *testing.T) {
	err := Convert(decodeFile, "jpg", "png")
	if err != nil {
		t.Fatal("failed test: Convert error")
	}

	// 変換後のファイルが存在するかチェック
	info, err := os.Stat(encodeFile)
	if err != nil {
		t.Fatal("failed test: File not generated")
	}
	//　変換後のファイルサイズが0バイトではないかチェック
	if info.Size() <= 0 {
		t.Fatal("failed test: Encoded file is invalid")
	}
	// テストで生成したファイルを削除する。
	err = os.Remove(encodeFile)
	if err != nil {
		t.Fatal("failed to remove the file")
	}
}
