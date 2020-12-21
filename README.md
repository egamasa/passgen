# passgen

パスワード生成ツール


## 使い方

- `-a`
  - パスワードに使用する文字に任意の文字を追加
  - ```powershell
    PS > passgen.exe -a "()[]"
    [0] 1cK%x(IDCzf7
    [1] Y9THsslknmz9
    [2] fITzHn0(GO!g
    [3] [cBWnN]w10]G
    [4] wu!CPhml7plf
    ```
- `-l`
  - パスワード文字列の長さを指定
  - デフォルト：12
  - ```powershell
    PS > passgen.exe -l 8
    [0] %mHD12s9
    [1] 5bd%Ce%#
    [2] ELhM@2Gc
    [3] hprEwXc7
    [4] A7N@!DPI
    ```
- `-n`
  - 生成するパスワード文字列の数を指定
  - デフォルト：5
  - ```powershell
    PS > passgen.exe -n 2
    [0] RkG4ahYDL9H8
    [1] hAsGYasIEG3l
    ```
- `-s`
  - パスワードに使用する文字種の組み合わせを指定
  - デフォルト：15（数字＋英小文字＋英大文字＋記号）
  - 下記の表を参照し、使用する文字種の数値を足した値を指定する
    - 例）英小文字(2)と英大文字(4)を使用 → 6 (2+4)
    - ```powershell
      PS > passgen.exe -s 6
      [0] mUWHQdpiABBv
      [1] AIjwtZYsejdf
      [2] aXSpldlTHnqU
      [3] reFuEFgRtTmG
      [4] fuWNpTcQtNPl
      ```

  | 数値 | 文字種 | 含まれる文字 |
  |---|----|----|
  | 1 | 数字 | 0-9 |
  | 2 | 英小文字 | a-z |
  | 4 | 英大文字 | A-Z |
  | 8 | 記号 | !@#$%^&* |


## Windows向けビルド

`passgen.exe` が出力される

- 64bit
  - ```bash
    $ GOOS=windows GOARCH=amd64 go build passgen.go
    ```
- 32bit
  - ```bash
    $ GOOS=windows GOARCH=386 go build passgen.go
    ```
