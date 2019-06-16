# 概要
* 画像変換ツールgophotoの改良
* io.Readerとio.Writerについて(下記)

# 画像変換ツールgophoto
使い方としては下記のようにdirectoryを-d, -iで入力形式, -oで出力形式を指定する  
-nで最大変換数を指定する. 現在の対応は.jpeg, .jpg, .png, .gif
```shell
$ gophoto -d dir -i .png -o .jpeg -n 10
```


# io.Readerとio.Writerの調査
* 標準パッケージでどのように使われているか
* io.Readerとio.Writerがあることでどういう利点があるのか

## 概要
* io.Writer: 出力の抽象化
* io.Reader: 入力の抽象化

```golang
type Writer interface {
	Write(p []byte)(n int, err error)
}
```

```golang
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

ReadとWriteをそれぞれ実装したものが全てReader, Writerのインターフェースとして扱われる

## 標準パッケージ
### Write/Readが実装されているもの
* syscall
	* システムコールに利用する
* file
	* ファイルの読み書き
* os.Stdout
	* 標準出力
* bytes.Buffer
	* バイトへの読み書き
* strings.Builder
	* Stringとのやりとり
* net.IPConn
	* ネットワークへの読み書き

## io.Writer
io.Writerという入出力に共通する処理を仕様として満たすインターフェース型を定義することでバイト列bを書き込んでバイト数n及びエラーeを出力する処理を通常のファイルに限定せずに汎用的に行うことができる

例えば、
io.MultiWriterのように複数のio.Writerを受け取り書き込み内容を同時に書き込む関数を使うことで、様々な形式の書き込みに対して抽象的に処理できる

## io.Reader
読み込むデータ型がRead関数を実装していれば予め用意した[]byte形式のバッファに読み込むことができる。ただし、バッファ管理をしつつ、毎回格納するバッファのサイズを用意してそこに突っ込むのは面倒なため、ioutil.RealAllのような補助関数を利用することでより利便性高く利用できる

## その他
io.WriteString()のような関数がある。これらはbyteのsliceを使わずに直接stringに書き込むため標準のWriteメソッドではなくこちらを使った方が効率がいい

## 類似するioインターフェース
* io.Closer
	* Closeメソッドを持ち、使用したファイルを閉じる
* io.Seeker
	* 読み書きの位置を移動する
* io.ReadAt
	* 対象オブジェクトがランダムアクセスできるときに特定の位置に自由にアクセスできる

## 参考
Goならわかるシステムプログラミング, 渋川よしき, LambdaNote, 2017
[How to use the io.Writer interface · YourBasic Go](https://yourbasic.org/golang/io-writer-interface-explained/)

## テストカバレッジ
- conv: 71.4%
- dir: 91.7%