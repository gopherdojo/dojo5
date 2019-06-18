# Typing Game

Gopher道場 #5 課題3-1 `タイピングゲームを作ろう` の実装です。

```
$ go run main.go -t 60 -f animals
```
などとすることでタイピングゲームを行うことができます。

## 困っている点

- goroutineを使う際のテストがよく分かりませんでした
- errorのチェックをする場合に、xerrors.Newで作ったものは内部的にframeという値を持ってしまうため、厳密にDeepReflectできず、どのように比較をしたらよろしいでしょうか