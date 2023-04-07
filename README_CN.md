# EnvX - Golang 环境变量工具库

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/envx)](https://goreportcard.com/report/github.com/goliajp/envx)
[![GoDoc](https://godoc.org/github.com/goliajp/envx?status.svg)](https://godoc.org/github.com/goliajp/envx)

---
[ENGLISH](README.md)
[日本語](README_JP.md)

EnvX 是一个简单且强大的Go环境变量管理工具库，用于在您的 Go 应用程序中管理环境变量。它可以轻松地获取带有默认值和正确类型的环境变量

## 安装

要开始使用 EnvX，请运行以下命令：

```bash
go get -u github.com/goliajp/envx
```

## 使用方法

在您的 Go 应用程序中导入 EnvX：

```go
import "github.com/goliajp/envx"
```

以下是一些使用 EnvX 的示例：

```go
// 示例 1：
// 获取带有默认值的 duration 类型环境变量
// TICKER_INTERVAL=10s -> TickerInterval = 10 * time.Second
var TickerInterval = envx.Get("ticker_interval", 120*time.Second)

// 示例 2：
// 获取带有默认值的 string 类型环境变量
// APP_NAME=foo -> AppName = "foo"
var AppName = envx.Get("app_name", "myapp")

// 示例 3：
// 获取带有默认值的 boolean 类型环境变量
// IS_TEST_MODE=true -> IsTestMode = true
var IsTestMode = envx.Get("is_test_mode", false)

// 示例 4：
// 获取带有默认值的 int 类型环境变量
// MYSQL_PORT=63306 -> MysqlPort = 63306
var MysqlPort = envx.Get("mysql_port", 3306)

// 示例 5：
// 获取带有默认值的 time.Time 类型环境变量
// DATE_TIME="2022-12-25 12:00:00" -> DateTime = time.Date(2022, 12, 25, 12, 0, 0, 0, time.UTC)
var DateTime = envx.Get("date_time", time.Date(2023, 4, 5, 18, 7, 0, 0, time.UTC))
```

## 贡献
我们欢迎大家为 EnvX 贡献！请在 GitHub 上提交问题，功能请求或拉取请求

## 许可证
本项目采用MIT许可证授权-有关详细信息，请参阅 [LICENSE](LICENSE) 文件