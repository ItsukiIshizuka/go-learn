# Go言語の導入
【参照サイト】<br>
[Go言語をはじめよう](https://news.mynavi.jp/techplus/article/gogogo-1/)<br>
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
    > HELLO, GOlang!!
  ```
4. Go言語でBMI値を計測し肥満度を表示するプログラム作成
  - bmi.goを作成しコンパイルして実行
  ```bash
    $ go build bmi.go
    $ ./bmi
      BMI=22.038567 肥満度=100.175307
      WEIGHT = int
      HEIGHT = int
      hm = float64
      bmi = float64
      bestW = float64
      per = float64
  ```

## わからなかったこと
### 【 hello.goで記述しているpackageとは何なのか？ 】
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
        // testingを使わないプログラムを記述
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
### 【 GOの型推論と型指定 】
- わからなかったこと
  - bmi.goで定義した定数および変数に型宣言する方法
  - 型推論が行われた後の定数・変数の型を表示する方法
- 調べたこと
  - 型宣言方法と型を表示する書式指定子
  - 参照：[Goの書式指定子_値の型](https://qiita.com/Sekky0905/items/c9cbda2498a685517ad0#t-%E5%80%A4%E3%81%AE%E5%9E%8B)
- わかったこと
  - `%T`の書式指定子で定数および変数の型を表示することができる
  - 定数および変数を定義する際に`const 型 定数名` or `var 型 変数名` とすることで型宣言ができる
  - 型宣言をしない場合は「型推論」が働き、コンパイル時に自動で型を判断し置き換えが行われる
  - 変数の宣言については以下で省略可能(基本的には型宣言するものだと思っている)
  ```go
    var test int = 1 //型宣言した形
    var test = 1     //型推論する形
    test := 1        //varも省略した形
  ```
  - 以下のようなパターンは実行結果が変わるため注意する
  ```go
    a := 10/3            //aはint型、出力は 3
    b := 10/3.0          //bはfloat64型、出力は 3.333333
    var c float64 = 10/3 //cはfloat64型、出力は 3.000000

    // aとcは計算自体は同じだが型が違う
    // bとcは計算は同じだが結果が違う
    // cはint型で計算した結果をfloat型の変数に代入しているため、小数点以下がbと違い0になっている
  ```
