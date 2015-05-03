package log

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLog(t *testing.T) {

	Convey("the logger", t, func() {
		Convey("should allow for all the methods to be called without bombing", func() {
			Debug("this is a debug message!!")
			Info("this is a info message")
			Warn("this is a warning message")
			Error("this is a panic message")
			Temp("this is a temporary message")
		})

		Convey("should allow for the minimium logger to be set", func() {
			defer func() { MinLevel = DebugLevel }() // cleanup

			Debug("current min level is %v, setting to %v", MinLevel, FatalLevel)
			MinLevel = FatalLevel
			Panic("this should not panic, which is bad, but the min level should reset it")
			Debug("hurray, we didn't panic, however this won't get printed because the level wasn't reset by this point")
		})
	})
}
