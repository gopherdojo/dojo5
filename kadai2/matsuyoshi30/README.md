## 課題1 and 課題2

### Usage

```bash
$ goconv [-from before image format] [-to after image format] path/to/dir...
```

```
NAME:
   goconv - image format converter written in Go

USAGE:
   goconv [-from before image format] [-to after image format] path/to/dir

VERSION:
   0.1.0

GLOBAL OPTIONS:
   -from           specify format before converted (jpg, png, gif) [DEFAULT: jpg]
   -to             specify format after converted  (png, jpg, gif) [DEFAULT: png]
   --help, -h      show help
```

---

## 課題2

### kadai2-2: io.Readerとio.Writerについて調べる

#### 標準パッケージでどのように使われているか

- Read, Write メソッドの実装
  - os.File, net.IPConn, strings.Reader, strings.Builder を見た
  - 自パッケージ内で定義した型（ Reader インタフェースを満たす）を返す関数 NewReader を定義
- 他のパッケージで、上の NewReader の戻り値を引数に取るようなメソッドで使われるパターン
  - ```dec := json.NewDecoder(strings. NewReader(`{"Name": "matsuyoshi30", "Text": "kadai2"}`))```
  - 参考⇒ https://play.golang.org/p/HY19J7lg3aC
  - Writer についても同様

#### io.Readerとio.Writerがあることで、どういう利点があるのか具体例を挙げて考えてみる

- io.Copy
  - ```func Copy(dst Writer, src Reader) (written int64, err error)```
  - 引数がインタフェースを満たした具象型でさえいれば io.Copy できる
  - ファイルでもネットワーク情報でも文字列でも同じように io.Copy が書ける
