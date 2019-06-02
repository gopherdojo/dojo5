// dircrawler パッケージは指定のディレクトリ以下にある指定の形式のファイルのリストを再帰的に取得する機能を提供します
package dircrawler

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// 指定ディレクトリ配下の特定形式のファイルパスリストを再帰的に取得します
func SearchSpecificFormatFiles(rootDir, format string) []string {

	paths := searchFilePaths(rootDir)
	files := selectFormat(paths, format)

	return files

}

// 指定ディレクトリのファイルパスリストを再帰的に取得します
func searchFilePaths(rootDir string) []string {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	var paths []string

	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, searchFilePaths(filepath.Join(rootDir, file.Name()))...)
			continue
		}

		paths = append(paths, filepath.Join(rootDir, file.Name()))
	}

	return paths
}

//ファイルパスリストから特定のフォーマットのもののみ抜き出します
func selectFormat(paths []string, format string) []string {
	var files []string
	for _, path := range paths {
		if isSpecifiedFormat(path, format) {
			files = append(files, path)
		}
	}
	return files
}

// ファイルパスが特定のフォーマットか確認します。
func isSpecifiedFormat(path, format string) bool {
	ext := filepath.Ext(path)
	return ext == format
}
