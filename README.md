# EnvX - Golang Environment Variable Utility

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/envx)](https://goreportcard.com/report/github.com/goliajp/envx)
[![GoDoc](https://godoc.org/github.com/goliajp/envx?status.svg)](https://godoc.org/github.com/goliajp/envx)

EnvX is a simple and powerful Go utility library for managing environment variables in your Go applications. It makes it easy to get environment variables with default values and proper types.

## Installation

To start using EnvX, simply run:

```bash
go get -u github.com/goliajp/envx
```

## Usage

Import EnvX into your Go application:

```go
import "github.com/goliajp/envx"
```

Here are some examples of how to use EnvX:

```go
// example 1:
// Get a duration environment variable with a default value
// TICKER_INTERVAL=10s -> TickerInterval = 10 * time.Second
var TickerInterval = envx.Get("ticker_interval", 120*time.Second)

// example 2:
// Get a string environment variable with a default value
// APP_NAME=foo -> AppName = "foo"
var AppName = envx.Get("app_name", "myapp")

// example 3:
// Get a boolean environment variable with a default value
// IS_TEST_MODE=true -> IsTestMode = true
var IsTestMode = envx.Get("is_test_mode", false)

// example 4:
// Get an integer environment variable with a default value
// MYSQL_PORT=63306 -> MysqlPort = 63306
var MysqlPort = envx.Get("mysql_port", 3306)

// example 5
// Get a time.Time environment variable with a default value
// DATE_TIME="2022-12-25 12:00:00" -> DateTime = time.Date(2022, 12, 25, 12, 0, 0, 0, time.UTC)
var DateTime = envx.Get("date_time", time.Date(2023, 4, 5, 18, 7, 0, 0, time.UTC))
```

## Contributing
We welcome contributions to EnvX! Feel free to submit issues, feature requests, or pull requests on GitHub.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.