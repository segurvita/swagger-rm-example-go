# swagger-rm-example-go
<div style="text-align:right">Language: <a href="README.md">English</a> | <i>日本語</i></div>



swaggerのコードからexampleを取り除きます。Goで書かれています。



## 目的

Amazon API GatewayにSwaggerファイルをインポートしたところ、以下のようなエラーが発生しました。

```
errors : [Invalid model schema specified. Unsupported keyword(s): ["example"]]
```

このエラーを回避することが目的です。



## Contributing

バグ報告やプルリクエストは、https://github.com/segurvita/swagger-rm-example-go にお願いします。



## 開発環境構築 

開発にはGoが必要です。 `go` コマンドが動作する環境を用意してください。

```bash
# パッケージ管理はdepを想定しています。は以下のコマンドでインストールできます。
go get -u github.com/golang/dep/cmd/dep

# 本プロジェクトをインストールします。
go get -u github.com/segurvita/swagger-rm-example-go

# プロジェクトフォルダに移動します。
cd $GOPATH/github.com/segurvita/swagger-rm-example-go

# 依存パッケージをインストールします。
dep ensure

# Main処理を実行します。
go run main.go

# テストコードを実行します。
go test ./lib/swagger/swagger_test.go
```



## 作成者

[@segur_vita](https://twitter.com/segur_vita)



## ライセンス

[MIT License](https://opensource.org/licenses/MIT)


