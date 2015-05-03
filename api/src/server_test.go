package main

import (
	"net/http"
	"testing"

	"github.com/shorrockin/clikr/log"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServer(t *testing.T) {

	Convey("the server entry point", t, func() {
		initial := len(routes)

		Convey("should allow a new route to be added", func() {
			Map("GET", "/testing", func(writer http.ResponseWriter, request *http.Request) {
				log.Debug("/testing was accessed")
			})

			So(len(routes), ShouldEqual, initial+1)

			Convey("and should have the proper http method set", func() {
				So(routes[initial].method, ShouldEqual, "GET")
			})

			Convey("and should have the proper path set", func() {
				So(routes[initial].path, ShouldEqual, "/testing")
			})
		})

		Convey("should be able to create a gorilla mux router", func() {
			So(initRouter(), ShouldNotBeNil)
		})
	})
}
