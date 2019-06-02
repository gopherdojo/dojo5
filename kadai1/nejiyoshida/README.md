# １.画像変換コマンドをつくろう

## 使い方
imgconv {-s source-extension} {-t target-extension} search-directory save-directory

## 説明
指定したディレクトリ配下を再帰的に探索し、指定形式の画像を変換します。  
保存先を指定する必要があり、指定の保存先が存在しない場合、新たにディレクトリが作られます。  
対応形式はjpg, png, bmp, gif の4種類です。