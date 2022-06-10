# Go言語の導入
【参照サイト】<br>
[Go言語をはじめのよう](https://news.mynavi.jp/techplus/article/gogogo-1/)<br>
<br>
主にGo言語の導入から、簡単なプログラムをコンパイルして実行するまで行う

## 手順
1. Go言語のインストール
  - macを使用しているためHomebrewを使用してインストールします
  ```bash
  brew install go
  ```
2. Go言語でHelloWorld!
  - hello.goファイルを作成しプログラムを実行
  - `$ go run hello.go` でプログラム実行
3. hello.goをビルドコマンドでコンパイルし再度実行する
  ```bash
    $ go build hello.go
    $ ./hello
    ## > HELLO, GOlang!!
  ```

## わからなかったこと
### 【hello.goで記述しているpackageとは何なのか？】
- 試したこと
  - 8行目 `func main()` を `func test()` に変更 -> エラーになった
  ```bash
    $ go run hello.go
    # command-line-arguments
    runtime.main_main·f: function main is undeclared in the main package
  ```
  - `func test()` のまま 2行目 `package main` を `package test` に変更 -> エラーになった
  ```bash
    $ go run hello.go
    package command-line-arguments is not a main package
  ```
- 調べたこと
  - 参照：[go run と go buildの違い](https://nununu.hatenablog.jp/entry/2016/09/20/210000)
    - packageやimportについても簡単に説明してくれていた
    - packageについて
      - 全てのGOプログラムは何かしらのパッケージの所属している
      - 原則としてのルールが3つ存在する
        - 1つのファイルに複数のパッケージを設定してはいけない
        - 1つのディレクトに複数のパッケージを置いてはならない
        - `package main` にある `main関数` がエンドポイントとされる(ここから処理が開始する)
    - importについて
      - プログラムの中で使用するパッケージを宣言する
      - プログラム内で使用されないパッケージを宣言していた場合はエラーがプログラム実行時にエラーが発生する
      ```go
        //hello.go
        package main

        // fmt と testingを宣言
        import (
	        "fmt"
          "testing"
        )

        func main() {
	        fmt.Printf("HELLO, GOlang!!\n")
        }
      ```
      ```bash
        $ go run hello.go
        # command-line-arguments
        ./hello.go:7:2: imported and not used: "testing"
      ```
    - buildについて
      - カレントディレクトリ以下をコンパイルしてバイナリファイルを作成する
    - runについて
      - 指定したファイルのみのコンパイルから実行までを行ってれる
      - buildと違いカレントディレクトリ以下を読み込むわけではないので、mainパッケージに所属する別関数などがあった場合は無視されてしまう
- わかったこと
  - package mainと使用するパッケージをimportする
  - 関数はmain関数を作成してその中にプログラムを記述する(package main で main関数じゃなきゃ動かない)
  - 簡単に実行したいときはrunコマンドで、ちゃんとしたいときはbuildコマンドを使用する
  - この辺りは全て曖昧なので追々学んでいく