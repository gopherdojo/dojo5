package converter

import (
	"log"
	"os"
	"testing"
)

const decodeFile string = "./cat.jpg"
const encodeFile string = "./cat.png"

func Test_Convert(t *testing.T) {
	Convert(decodeFile, "jpg", "png")

	// 変換後のファイルが存在するかチェック
	info, err := os.Stat(encodeFile)
	if err != nil {
		log.Fatal(err)
	}
	//　変換後のファイルサイズが0バイトではないかチェック
	if info.Size() <= 0 {
		log.Fatal("encoded file is not valid")
	}
	// テストで生成したファイルを削除する。ファイルが生成されていない場合はテストに失敗したと見做す。
	err = os.Remove(encodeFile)
	if err != nil {
		t.Fatal("failed test")
	}
}
