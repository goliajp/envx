package envx

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Get the system environment variables, with default values
// now support [T int | float64 | bool | string | time.Duration | time.Time ]
//
// example 1: var TickerInterval = envx.Get("ticker_interval", 120 * time.Second)
// example 2: var AppName = envx.Get("app_name", "myapp")
// example 3: var IsTestMode = envx.Get("is_test_mode", false)
// example 4: var MysqlPort = envx.Get("mysql_port", 3306)
// example 5: var DateTime = envx.Get("date_time", time.Date(2023, 4, 5, 18, 7, 0, 0, time.UTC))
func Get[T int | float64 | bool | string | time.Duration | time.Time](k string, v T) T {
	// default value
	str := os.Getenv(strings.ToUpper(k))
	if str == "" {
		fmt.Printf("%s: %v (default)\n", strings.ToLower(k), v)
		return v
	}

	// parse env
	var vi any = &v
	switch rv := vi.(type) {
	case *string:
		*rv = str
	case *float64:
		fv, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Fatalf("convert str to float64 failed: %v", err)
		}
		*rv = fv
	case *int:
		iv, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatalf("convert str to int failed: %v", err)
		}
		*rv = int(iv)
	case *bool:
		if str == "1" || str == "TRUE" || str == "true" || str == "ON" || str == "on" {
			*rv = true
		} else {
			*rv = false
		}
	case *time.Duration:
		d, err := time.ParseDuration(str)
		if err != nil {
			log.Fatalf("convert str to time.Duration failed: %v", err)
		}
		*rv = d
	case *time.Time:
		t, err := time.Parse(time.DateTime, str)
		if err != nil {
			log.Fatalf("convert str to time.Time failed: %v", err)
		}
		*rv = t
	default:
		log.Fatalln("unsupported type")
	}

	// log
	if Config.IsLog {
		fmt.Printf("%s: %v\n", strings.ToLower(k), v)
	}
	return v
}
