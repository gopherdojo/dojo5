# Split Downloader (my pget)

A CLI for split downloading, kadai3-2 of Gopherdojo #5.

Gopher道場 #5 課題3-2 `分割ダウンローダを作ろう` の実装です。

## Installation

```bash
$ make build
```

## Usage

URL to download is required as argument.

ダウンロードするURLを引数として渡してください。


```bash
$ bin/mypget https://domain.name/path/to/file
```

With `-n` option, you can specify the number of split ranges. Default number is `8` . The number must be less than the length of file to download.

`-n` オプションで分割の数を指定できます。デフォルトの数は `8` です。分割数はダウンロード対象のファイルサイズよりも小さい必要があります。

```bash
$ bin/mypget -n 16 https://domain.name/path/to/file
```
