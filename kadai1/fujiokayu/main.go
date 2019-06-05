package main

import (
	"log"
	"myConverter/args"
	"myConverter/converter"
	"myConverter/walker"
	"path/filepath"
	"strings"
)

// フォルダ探索によって見つかったファイルの拡張子が、decodeType と一致した場合に converter を起動する。
func execute(filePath string, decodeType string, encodeType string) {
	// filepath.Ext が抽出する拡張子は "." を含むため、オプションで指定された拡張子と揃えるために "." を除去する。
	ext := strings.ToLower(strings.Trim(filepath.Ext(filePath), "."))

	if ext == strings.ToLower(decodeType) {
		err := converter.Convert(filePath, ext, encodeType)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	args, err := args.ParseArgs()
	if err != nil {
		log.Fatal(err)
	}

	folder := strings.Join(args.RootFolderName, " ")

	ch, err := walker.Walk(folder)
	if err != nil {
		log.Fatal(err)
	}
	for filePath := range ch {
		execute(filePath, args.DecodeType, args.EncodeType)
	}
}
