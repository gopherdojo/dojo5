# myConverter

## how to build
```
make
```

## usage

```
# convert
./myConverter [-from ext] [-to ext] directory

# show help
./myConverter -h

# example 
./myConverter -from png -to jpg testdir
```

## Specification
- 引数で指定されたディレクトリを再帰的に走査し、-from オプションで指定された画像形式のファイルを -to オプションで指定された画像形式のファイルに変換します。
    - オプションで指定できるフォーマットは gif、jpg(jpeg)、png のみです。
    - オプションを指定しなかった場合は jpg -> png に変換します。
- 変換後のファイルは変換元のファイルと同じディレクトリに出力されます。
- 変換元のファイルは削除されず、そのまま残ります。
- 変換中にディスクサイズが足りなくなった場合の挙動は（恐らく）処理系定義になります。

## learning memo
- ディレクトリの再帰的な探索は [path/filepath.Walk](https://golang.org/pkg/path/filepath/#Walk) を使うのが一番近道だと思った。  
    - しかし、filepath.Walk 内で呼び出せる関数が [WalkFunc](https://golang.org/pkg/path/filepath/#WalkFunc) に限定されているようで、かつ WalkFunc の型が決まっていたので少し使い辛かった。
    - 一方で、これは[先月の Software Design](https://gihyo.jp/magazine/SD/archive/2019/201905) で見た Generator Pattern を試すと Go らしくなるのでは考えた。
    - [同じことを考える人](https://gist.github.com/sethamclean/9475737)が既に居たので、参考にした。
- flag パッケージに少し使い辛さを感じる。
    - 特に、フラグ無し引数が一つ入ると以降のフラグも全て Parse できなくなるのが少し使いづらい。
