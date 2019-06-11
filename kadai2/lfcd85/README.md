# imgconv

An implementation for the image conversion command, kadai-1 of Gopherdojo #5.

Gopher道場 #5 課題1 `画像変換コマンド` の実装です。

## Installation

```bash
$ make bin/convert
```

## Usage

Specify the target directory as an argument. The given directory is recursively processed. Converted files are outputted under `./output/` directory.

コマンド引数に対象ディレクトリを指定してください。ディレクトリ以下は再帰的に処理されます。変換後のファイルは `./output/` ディレクトリ以下に出力されます。

```bash
$ bin/convert test/
```

Input and output image formats can be set by `-f` (from) and `-t` (to) options. Default formats are from JPEG to PNG. JPEG, PNG, GIF are available.

画像形式は `-f` オプション（変換前）・ `-t` オプション（変換後）で指定できます。デフォルトは JPEG → PNG です。JPEG, PNG, GIF 形式が利用可能です。

```bash
$ bin/convert -f png -t jpeg test/
```
