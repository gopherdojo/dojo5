package convert

/*
画像変換コマンド

次の仕様を満たすコマンドを作って下さい

    ディレクトリを指定する
    指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
    ディレクトリ以下は再帰的に処理する
    変換前と変換後の画像形式を指定できる（オプション）

以下を満たすように開発してください

    mainパッケージと分離する
    自作パッケージと標準パッケージと準標準パッケージのみ使う
    準標準パッケージ：golang.org/x以下のパッケージ
    ユーザ定義型を作ってみる
    GoDocを生成してみる
*/

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// picture file 変換関数
func PixFile(org string, src string, dest string) error {

	// open file
	file, err := os.Open(org)
	if err != nil {
		return err
	}
	defer file.Close()

	// image reading.
	img, format, err := image.Decode(file)
	if err != nil {
		// not image
		return err
	}

	// 元ファイルが指定外ならスキップ
	if format != src {
		return nil
	}

	// 出力先ファイル
	savefile, err := os.Create(org + "." + dest)
	if err != nil {
		return err
	}
	defer savefile.Close()

	switch dest {
	case "jpg", "jpeg":
		opts := &jpeg.Options{}
		jpeg.Encode(savefile, img, opts)
	case "png":
		png.Encode(savefile, img)
	case "gif":
		opts := &gif.Options{}
		gif.Encode(savefile, img, opts)
	}

	return nil
}
