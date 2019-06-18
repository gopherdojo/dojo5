# Typing Game

An implementation for the typing game, kadai3-1 of Gopherdojo #5. The standard library packages for Golang are the words for typing.

Gopher道場 #5 課題3-1 `タイピングゲームを作ろう` の実装です。Go言語の標準ライブラリのパッケージ名がタイピング対象の単語になっています。

## Installation

```bash
$ make build
```

## Usage

The game starts after executing the command below. Let's type the shown package name. Your score will be shown after reaching to the time limit and finishing the game.

下記のコマンドを実行するとゲームが開始します。表示されたパッケージ名をタイプしてください。制限時間に達するとゲームが終了し、スコアが表示されます。

```bash
$ bin/typinggame
Let's type the standard package names! ( Time limit: 30s )
> hash/fnv
hash/fnv
hash/fnv ... OK! current score: 1

> debug
d 
d ... NG: try again.
> debug

30s has passed: you correctly typed 1 package(s)!
```

The time limit can be set by `-t` option. Default value is 30 sec.

制限時間は `-t` オプションで変更可能です。デフォルトは30秒です。

```bash
$ bin/typinggame -t 10
Let's type the standard package names! ( Time limit: 10s )
```
