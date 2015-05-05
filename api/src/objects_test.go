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

		Convey("should have a unique image url", func() {
			So(SceneObjects[0].Image, ShouldEqual, "http://placehold.it/140&text=Glasses")
		})
	})
}
