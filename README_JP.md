# EnvX - Golang 環境変数ユーティリティ

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/envx)](https://goreportcard.com/report/github.com/goliajp/envx)
[![GoDoc](https://godoc.org/github.com/goliajp/envx?status.svg)](https://godoc.org/github.com/goliajp/envx)

---
[ENGLISH](README.md)
[简体中文](README_CN.md)

EnvX は、Go アプリケーションで環境変数を管理するためのシンプルで強力なユーティリティです。デフォルト値と適切な型を持つ環境変数を簡単に取得できます

## インストール

EnvX を使い始めるには、以下のコマンドを実行してください:

```bash
go get -u github.com/goliajp/envx
```

## 使い方

Go アプリケーションに EnvX をインポートします:

```go
import "github.com/goliajp/envx"
```

EnvX の使い方の例は以下の通りです:

```go
// 例 1：
// デフォルト値付きの duration 型環境変数を取得
// TICKER_INTERVAL=10s -> TickerInterval = 10 * time.Second
var TickerInterval = envx.Get("ticker_interval", 120*time.Second)

// 例 2：
// デフォルト値付きの string 型環境変数を取得
// APP_NAME=foo -> AppName = "foo"
var AppName = envx.Get("app_name", "myapp")

// 例 3：
// デフォルト値付きの boolean 型環境変数を取得
// IS_TEST_MODE=true -> IsTestMode = true
var IsTestMode = envx.Get("is_test_mode", false)

// 例 4：
// デフォルト値付きの int 型環境変数を取得
// MYSQL_PORT=63306 -> MysqlPort = 63306
var MysqlPort = envx.Get("mysql_port", 3306)

// 例 5：
// デフォルト値付きの time.Time 型環境変数を取得
// DATE_TIME="2022-12-25 12:00:00" -> DateTime = time.Date(2022, 12, 25, 12, 0, 0, 0, time.UTC)
var DateTime = envx.Get("date_time", time.Date(2023, 4, 5, 18, 7, 0, 0, time.UTC))
```

## 貢献
EnvX へのコントリビューションを歓迎します！GitHub で問題や機能リクエスト、プルリクエストを提出してください

## ライセンス
このプロジェクトは、MIT ライセンスでライセンスされています。詳細については、 [LICENSE](LICENSE) ファイルを参照してください