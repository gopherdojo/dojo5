package walker

import (
	"fmt"
	"os"
	"path/filepath"
)

// assertDir: 引数の文字列が有効なディレクトリか検証する
func assertDir(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		if !info.IsDir() {
			err = fmt.Errorf("directory is not valid")
		}
	}
	return err
}

// Walk :指定されたディレクトリを再帰的に操作し、見つかったファイルをチャネルで返す。
func Walk(rootPath string) (chan string, error) {

	ch := make(chan string)
	err := assertDir(rootPath)
	if err != nil {
		return ch, err
	}

	// Todo: Add error handling
	go func() {
		err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				ch <- path
			}
			return nil
		})
		defer close(ch)
	}()
	return ch, err
}
