package main

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBonds(t *testing.T) {
	Convey("when retrieving a bond", t, func() {
		s1 := &SceneObject{ObjectID: "1"}
		s2 := &SceneObject{ObjectID: "2"}
		s3 := &SceneObject{ObjectID: "3"}

		startLength := len(Bonds)

		Convey("should be able to create a bond array for a new scene", func() {
			GetBond(s1, s2)
			So(len(Bonds), ShouldEqual, startLength+1)

			GetBond(s1, s2) // should not increment it further
			So(len(Bonds), ShouldEqual, startLength+1)
		})

		Convey("should be able to return the bond object if it does exist", func() {
			bond := GetBond(s1, s3)
			So(bond, ShouldNotBeNil)
			So(bond.With, ShouldEqual, s3)
			So(bond.Strength, ShouldEqual, 0)
		})
	})

	Convey("when strengthening a bond", t, func() {
		Convey("should increase the bond between two objects", func() {
			s1 := &SceneObject{ObjectID: "1"}
			s2 := &SceneObject{ObjectID: "2"}
			s3 := &SceneObject{ObjectID: "3"}

			str := StrengthenBond(s1, s2)
			So(str, ShouldEqual, 1)

			str = StrengthenBond(s1, s2) // inverse should be the same
			So(str, ShouldEqual, 2)

			str = StrengthenBond(s1, s3)
			So(str, ShouldEqual, 1)

			str = StrengthenBond(s1, s2) // while adding to the bond[] we should clear old values
			So(str, ShouldEqual, 3)
		})

		Convey("should ensure things are sorted in order", func() {
			s1 := &SceneObject{ObjectID: "1"}
			s2 := &SceneObject{ObjectID: "2"}
			s3 := &SceneObject{ObjectID: "3"}

			So(StrengthenBond(s1, s2), ShouldEqual, 1)
			So(StrengthenBond(s1, s2), ShouldEqual, 2)
			So(StrengthenBond(s1, s3), ShouldEqual, 1)

			So(Bonds[s1][0].With, ShouldEqual, s2)
			So(Bonds[s1][1].With, ShouldEqual, s3)

			So(StrengthenBond(s1, s3), ShouldEqual, 2)
			So(StrengthenBond(s1, s3), ShouldEqual, 3)

			So(Bonds[s1][0].With, ShouldEqual, s3)
			So(Bonds[s1][1].With, ShouldEqual, s2)
		})
	})

	Convey("when registering a click", t, func() {
		Convey("it should return an array of recommendations based on bond strength", func() {
			r := rand.New(rand.NewSource(1))
			obj1 := randomSceneObject(r)
			obj2 := randomSceneObject(r)
			obj3 := randomSceneObject(r)

			// just make sure none of our randoms selected the same thing, since seeded
			// if it works once - should continue to work
			So(obj1, ShouldNotEqual, obj2)
			So(obj1, ShouldNotEqual, obj3)
			So(obj2, ShouldNotEqual, obj3)

			fmt.Printf("wtf %v\n", obj1)

			vars := map[string]string{"object_id": obj1.ObjectID, "X-Clikr-User": "abc"}
			result := RegisterObjectClick(vars).([]*SceneObject)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 0)

			vars = map[string]string{"object_id": obj2.ObjectID, "X-Clikr-User": "abc"}
			result = RegisterObjectClick(vars).([]*SceneObject)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 1)
			So(result[0], ShouldEqual, obj1)

			vars = map[string]string{"object_id": obj3.ObjectID, "X-Clikr-User": "abc"}
			result = RegisterObjectClick(vars).([]*SceneObject)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 2)
			So(result[0], ShouldEqual, obj1)
			So(result[1], ShouldEqual, obj2)
		})
	})

}
