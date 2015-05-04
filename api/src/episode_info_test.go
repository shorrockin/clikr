package main

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEpisodeInfo(t *testing.T) {
	Convey("should be able to retrieve all the info for an episode", t, func() {
		result := RetrieveEpisodeInfo(nil).(EpisodeInfo)

		So(result.Slug, ShouldEqual, "new_girl")
		So(result.Episode, ShouldEqual, 21)
		So(result.Season, ShouldEqual, 4)
		So(result.FPS, ShouldEqual, 24)
		So(len(result.Scenes), ShouldEqual, 59)
	})

	Convey("should be able to render the episode info as json", t, func() {
		// info := RetrieveEpisodeInfo(nil).(EpisodeInfo)
		info := EpisodeInfo{Slug: "new_girl"}
		So(info, ShouldNotBeNil)
		So(info.Slug, ShouldEqual, "new_girl")

		bytes, error := json.Marshal(info)
		So(bytes, ShouldNotBeNil)
		So(error, ShouldBeNil)

		var unmarshelled EpisodeInfo
		err := json.Unmarshal(bytes, &unmarshelled)

		So(err, ShouldBeNil)
		So(unmarshelled, ShouldNotBeNil)
		So(unmarshelled.Slug, ShouldEqual, info.Slug)
	})
}
