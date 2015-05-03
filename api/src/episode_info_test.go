package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEpisodeInfo(t *testing.T) {
	Convey("should be able to retrieve all the info for an episode", t, func() {
		result := RetrieveEpisodeInfo(nil).(EpisodeInfo)

		So(result.slug, ShouldEqual, "new_girl")
		So(result.episode, ShouldEqual, 21)
		So(result.season, ShouldEqual, 4)
		So(result.fps, ShouldEqual, 24)
		So(len(result.scenes), ShouldEqual, 59)
	})
}
