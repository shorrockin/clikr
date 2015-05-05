package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProfiles(t *testing.T) {
	Convey("should be able to create and load a profile", t, func() {
		profile := GetUserProfile("abc")
		So(profile, ShouldNotBeNil)

		again := GetUserProfile("abc")
		So(profile, ShouldEqual, again)
	})

	Convey("should be able to like objects, which strengthen their bonds", t, func() {
		s1 := &SceneObject{ObjectID: "1"}
		s2 := &SceneObject{ObjectID: "1"}
		s3 := &SceneObject{ObjectID: "1"}

		LikeObject("abc", s1)
		LikeObject("abc", s2)

		So(GetBond(s1, s2).Strength, ShouldEqual, 1)
		So(GetBond(s2, s1).Strength, ShouldEqual, 1)
		So(GetBond(s1, s3).Strength, ShouldEqual, 0)

		LikeObject("abc", s3)

		So(GetBond(s1, s3).Strength, ShouldEqual, 1)
		So(GetBond(s2, s3).Strength, ShouldEqual, 1)
		So(GetBond(s1, s2).Strength, ShouldEqual, 1)

		LikeObject("abc", s2)

		So(GetBond(s1, s3).Strength, ShouldEqual, 1)
		So(GetBond(s2, s3).Strength, ShouldEqual, 2)
		So(GetBond(s1, s2).Strength, ShouldEqual, 2)

		LikeObject("efg", s1)
		LikeObject("efg", s3)

		So(GetBond(s1, s3).Strength, ShouldEqual, 2)
		So(GetBond(s2, s3).Strength, ShouldEqual, 2)
		So(GetBond(s1, s2).Strength, ShouldEqual, 2)
	})
}
