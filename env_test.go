package envx

import (
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {

	// int

	Convey("int default", t, func() {
		val := Get("INT", 1)
		So(val, ShouldEqual, val)
	})
	Convey("int", t, func() {
		if err := os.Setenv("INT", "1"); err != nil {
			t.Errorf("set environment variable failed: %v", err)
		}
		val := Get("INT", 0)
		So(val, ShouldEqual, 1)
	})

	// float64

	Convey("float64 default", t, func() {
		val := Get("FLOAT", 1.1)
		So(val, ShouldEqual, val)
	})
	Convey("float64", t, func() {
		if err := os.Setenv("FLOAT", "1.1"); err != nil {
			t.Errorf("set environment variable failed: %v", err)
		}
		val := Get("FLOAT", 1.1)
		So(val, ShouldEqual, 1.1)
	})

	// bool

	Convey("bool default", t, func() {
		val := Get("BOOL", true)
		So(val, ShouldBeTrue)
	})
	Convey("bool", t, func() {
		if err := os.Setenv("BOOL", "true"); err != nil {
			t.Errorf("set environment variable failed: %v", err)
		}
		val := Get("BOOL", false)
		So(val, ShouldBeTrue)
	})

	// string

	Convey("string default", t, func() {
		val := Get("STRING", "foo")
		So(val, ShouldEqual, "foo")
	})
	Convey("string", t, func() {
		if err := os.Setenv("STRING", "foo"); err != nil {
			t.Errorf("set environment variable failed: %v", err)
		}
		val := Get("STRING", "bar")
		So(val, ShouldEqual, "foo")
	})

	// time.Duration

	Convey("time.Duration default", t, func() {
		val := Get("DURATION", time.Second*1)
		So(val, ShouldEqual, time.Second*1)
	})
	Convey("time.Duration", t, func() {
		if err := os.Setenv("DURATION", "1s"); err != nil {
			t.Errorf("set environment variable failed: %v", err)
		}
		val := Get("DURATION", time.Second*10)
		So(val, ShouldEqual, time.Second*1)
	})

	// time.Time

	Convey("time.Time default", t, func() {
		val := Get("TIME", time.Date(2023, 4, 5, 18, 7, 0, 0, time.UTC))
		So(val, ShouldEqual, time.Date(2023, 4, 5, 18, 7, 0, 0, time.UTC))
	})
	Convey("time.Time", t, func() {
		if err := os.Setenv("TIME", "1949-10-01 09:00:00"); err != nil {
			t.Errorf("set environment variable failed: %v", err)
		}
		val := Get("TIME", time.Date(2023, 4, 5, 18, 7, 0, 0, time.UTC))
		So(val, ShouldEqual, time.Date(1949, 10, 1, 9, 0, 0, 0, time.UTC))
	})

	// config
	Convey("config", t, func() {
		Config.SetIsLog(true)
		So(Config.IsLog, ShouldBeTrue)
		Config.SetIsLog(false)
		So(Config.IsLog, ShouldBeFalse)
	})
}
