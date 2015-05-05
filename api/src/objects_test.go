package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestObjects(t *testing.T) {
	Convey("the objects", t, func() {
		Convey("should all be created with sample data", func() {
			So(len(SceneObjects), ShouldEqual, 15)
		})
	})
}
