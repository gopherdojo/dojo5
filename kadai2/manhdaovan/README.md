## 概要
- このツールについて
- io.Readerとio.Writerについて

## このツールについて
### 使い方
```
imgconv -- convert image to other image type
usage: $./bin/imgconv [-s] [-d] [-k] dir
options:
  -s string
	Source image type. jpeg, png and gif are supported
  -d string
  	Destination image type. jpeg, png and gif are supported
  -k
	Skip error and continue while converting
```

### Build & test
```bash
$cd /path/this/repo
# for install deps
$make install
# for build
$make build
# for test 
$make test
# for test with coverage
$make coverage
```

## io.Readerとio.Writerについて
### 標準パッケージでどのように使われているか
- IO系の抽象化。
- IO系の操作の箇所に使われてる。
- 抽象化なので同じデータを複数のIOの対象で出力・入力できる。
  - Network
  - Memory
  - Disk
- 抽象化なのでテストがしやすくなってる

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
ソースコードがテストしやすいって利点について挙げて考えてみる。

例えば、下記のソースはキーボードからテキストを入力して、画面に出力するって機能です：
```go
// GetTextAndDisplay reads and display data to io
func GetTextAndDisplay(r io.Reader, w io.Writer) error {
	buf := make([]byte, 10<<20) // 10MB
	readBytes, err := r.Read(buf)
	if err != nil {
		return err
	}

	_, err = w.Write(buf[0:readBytes])
	if err != nil {
		return err
	}

	return nil
}
// 利用する
func main(){
  if err := GetTextAndDisplay(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
```

- もしそもそも`GetTextAndDisplay`は`GetTextAndDisplay(r os.Stdin, w os.Stdout)`になったら、テストに`os.Stdinとos.Stdout`を直接に利用すると、実際のテキストの読み込みと書き出しが正しいかどうかを難しい。
- なのでio.Readerとio.Writerを利用すると、テストがしやすくなってます。
  - テスト用のio.Readerのmock:
```go
  type mockReaderOk struct {
    bytes []byte
  }

  func (mro *mockReaderOk) Read(bytes []byte) (int, error) {
    src := []byte("reader ok")
    for i, b := range src {
      bytes[i] = b
    }
    return len(src), nil
  }

  type mockReaderErr struct{}

  func (mre *mockReaderErr) Read(b []byte) (int, error) {
    return 0, fmt.Errorf("error on Reader")
  }
```

  - テスト用のio.Writerのmock
```go
  type mockWriterOk struct {
    bytes []byte
  }

  func (mwo *mockWriterOk) Write(b []byte) (int, error) {
    mwo.bytes = b
    return len(b), nil
  }

  type mockWriterErr struct {
    bytes []byte
  }

  func (mwe *mockWriterErr) Write(b []byte) (int, error) {
    return 0, fmt.Errorf("error on Writer")
  }
```

  - 一つのテストの例
```go
t.Run("reader ok, writer ok", func(t *testing.T) {
  r := &mockReaderOk{}
  w := &mockWriterOk{}
  var wantErr error
  wantStr := "reader ok"

  if err := GetTextAndDisplay(r, w); !reflect.DeepEqual(wantErr, err) {
    t.Errorf("GetTextAndDisplay() error = %v, wantErr %v", err, wantErr)
    return
  }
  if gotStr := string(w.bytes); gotStr != wantStr {
    t.Errorf("GetTextAndDisplay() = %v, want %v", gotStr, wantStr)
  }
})
```

全てのテストのソースは[こちら](https://gist.github.com/manhdaovan/6dbd06964eb2e6eaca0c80afe4d678b6)です