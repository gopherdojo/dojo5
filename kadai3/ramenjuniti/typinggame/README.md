# Typing Game

## Description

タイピングゲームです。

一行ずつ判定します。

`Score` は、どれぐらい出題単語を正しくタイプできているかどうかで加算されます。

例えば、`abcdef`と出題された時に`abcd`とタイプした場合、Score は 4 ポイント加算されます。

`type/sec`は、1 秒間に正しく入力したタイプ数です。

## Usage

```bash
make build
./typinggame -t {制限時間}
```

## Build

```bash
make build
```

## Test

```bash
make test
```
