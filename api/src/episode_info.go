package main

import (
	"math/rand"

	"github.com/shorrockin/clikr/log"
)

type EpisodeInfo struct {
	Slug    string      `json:"slug"`
	Name    string      `json:"name"`
	Season  uint        `json:"season"`
	Episode uint        `json:"episode"`
	FPS     uint        `json:"fps"`
	Scenes  []SceneInfo `json:"scenes"`
}

type SceneInfo struct {
	Image        string            `json:"image"`
	Frame        uint              `json:"frame"`
	Interactions []InteractionInfo `json:"interactions"`
}

type InteractionInfo struct {
	ObjectID  uint64      `json:"id"`
	PositionX int         `json:"x"`
	PositionY int         `json:"y"`
	Size      int         `json:"size"`
	Object    SceneObject `json:"object"`
}

func createInteractions(seed int64) []InteractionInfo {
	r := rand.New(rand.NewSource(seed))
	amount := (r.Intn(4) + 1)
	arry := make([]InteractionInfo, amount, amount)

	for count := 0; count < amount; count++ {
		arry[count] = InteractionInfo{
			ObjectID:  NextId(),
			PositionX: r.Intn(600) + 40,
			PositionY: r.Intn(300) + 30,
			Size:      r.Intn(3) + 1,
			Object:    SceneObjects[r.Intn(len(SceneObjects)-1)],
		}
	}

	return arry
}

// used to store the static test data
var newGirlEpisodeInfo EpisodeInfo

func init() {
	newGirlEpisodeInfo = EpisodeInfo{
		Slug:    "new_girl",
		Name:    "New Girl",
		Episode: 21,
		Season:  4,
		FPS:     24,
		Scenes: []SceneInfo{
			{Frame: 49, Image: "/images/new_girl/s4e21/scene00049.jpeg", Interactions: createInteractions(49)},
			{Frame: 145, Image: "/images/new_girl/s4e21/scene00145.jpeg", Interactions: createInteractions(145)},
			{Frame: 193, Image: "/images/new_girl/s4e21/scene00193.jpeg", Interactions: createInteractions(193)},
			{Frame: 289, Image: "/images/new_girl/s4e21/scene00289.jpeg", Interactions: createInteractions(289)},
			{Frame: 673, Image: "/images/new_girl/s4e21/scene00673.jpeg", Interactions: createInteractions(673)},
			{Frame: 721, Image: "/images/new_girl/s4e21/scene00721.jpeg", Interactions: createInteractions(721)},
			{Frame: 817, Image: "/images/new_girl/s4e21/scene00817.jpeg", Interactions: createInteractions(817)},
			{Frame: 1057, Image: "/images/new_girl/s4e21/scene01057.jpeg", Interactions: createInteractions(1057)},
			{Frame: 2017, Image: "/images/new_girl/s4e21/scene02017.jpeg", Interactions: createInteractions(2017)},
			{Frame: 2257, Image: "/images/new_girl/s4e21/scene02257.jpeg", Interactions: createInteractions(2257)},
			{Frame: 3025, Image: "/images/new_girl/s4e21/scene03025.jpeg", Interactions: createInteractions(3025)},
			{Frame: 3745, Image: "/images/new_girl/s4e21/scene03745.jpeg", Interactions: createInteractions(3745)},
			{Frame: 3841, Image: "/images/new_girl/s4e21/scene03841.jpeg", Interactions: createInteractions(3841)},
			{Frame: 4513, Image: "/images/new_girl/s4e21/scene04513.jpeg", Interactions: createInteractions(4513)},
			{Frame: 4657, Image: "/images/new_girl/s4e21/scene04657.jpeg", Interactions: createInteractions(4657)},
			{Frame: 4945, Image: "/images/new_girl/s4e21/scene04945.jpeg", Interactions: createInteractions(4945)},
			{Frame: 4993, Image: "/images/new_girl/s4e21/scene04993.jpeg", Interactions: createInteractions(4993)},
			{Frame: 5233, Image: "/images/new_girl/s4e21/scene05233.jpeg", Interactions: createInteractions(5233)},
			{Frame: 5569, Image: "/images/new_girl/s4e21/scene05569.jpeg", Interactions: createInteractions(5569)},
			{Frame: 5665, Image: "/images/new_girl/s4e21/scene05665.jpeg", Interactions: createInteractions(5665)},
			{Frame: 6529, Image: "/images/new_girl/s4e21/scene06529.jpeg", Interactions: createInteractions(6529)},
			{Frame: 7201, Image: "/images/new_girl/s4e21/scene07201.jpeg", Interactions: createInteractions(7201)},
			{Frame: 7585, Image: "/images/new_girl/s4e21/scene07585.jpeg", Interactions: createInteractions(7586)},
			{Frame: 7681, Image: "/images/new_girl/s4e21/scene07681.jpeg", Interactions: createInteractions(7681)},
			{Frame: 8929, Image: "/images/new_girl/s4e21/scene08929.jpeg", Interactions: createInteractions(8929)},
			{Frame: 9793, Image: "/images/new_girl/s4e21/scene09793.jpeg", Interactions: createInteractions(9793)},
			{Frame: 10081, Image: "/images/new_girl/s4e21/scene10081.jpeg", Interactions: createInteractions(10081)},
			{Frame: 11089, Image: "/images/new_girl/s4e21/scene11089.jpeg", Interactions: createInteractions(11089)},
			{Frame: 11185, Image: "/images/new_girl/s4e21/scene11185.jpeg", Interactions: createInteractions(11185)},
			{Frame: 11329, Image: "/images/new_girl/s4e21/scene11329.jpeg", Interactions: createInteractions(11329)},
			{Frame: 11521, Image: "/images/new_girl/s4e21/scene11521.jpeg", Interactions: createInteractions(11521)},
			{Frame: 12433, Image: "/images/new_girl/s4e21/scene12433.jpeg", Interactions: createInteractions(12433)},
			{Frame: 13201, Image: "/images/new_girl/s4e21/scene13201.jpeg", Interactions: createInteractions(13201)},
			{Frame: 13441, Image: "/images/new_girl/s4e21/scene13441.jpeg", Interactions: createInteractions(13441)},
			{Frame: 14257, Image: "/images/new_girl/s4e21/scene14257.jpeg", Interactions: createInteractions(14257)},
			{Frame: 14401, Image: "/images/new_girl/s4e21/scene14401.jpeg", Interactions: createInteractions(14401)},
			{Frame: 15505, Image: "/images/new_girl/s4e21/scene15505.jpeg", Interactions: createInteractions(15505)},
			{Frame: 17089, Image: "/images/new_girl/s4e21/scene17089.jpeg", Interactions: createInteractions(17089)},
			{Frame: 18145, Image: "/images/new_girl/s4e21/scene18145.jpeg", Interactions: createInteractions(18145)},
			{Frame: 18193, Image: "/images/new_girl/s4e21/scene18193.jpeg", Interactions: createInteractions(18193)},
			{Frame: 18289, Image: "/images/new_girl/s4e21/scene18289.jpeg", Interactions: createInteractions(18289)},
			{Frame: 18337, Image: "/images/new_girl/s4e21/scene18337.jpeg", Interactions: createInteractions(18337)},
			{Frame: 18913, Image: "/images/new_girl/s4e21/scene18913.jpeg", Interactions: createInteractions(18913)},
			{Frame: 21025, Image: "/images/new_girl/s4e21/scene21025.jpeg", Interactions: createInteractions(21025)},
			{Frame: 21745, Image: "/images/new_girl/s4e21/scene21745.jpeg", Interactions: createInteractions(21745)},
			{Frame: 22801, Image: "/images/new_girl/s4e21/scene22801.jpeg", Interactions: createInteractions(22801)},
			{Frame: 23473, Image: "/images/new_girl/s4e21/scene23473.jpeg", Interactions: createInteractions(23473)},
			{Frame: 24097, Image: "/images/new_girl/s4e21/scene24097.jpeg", Interactions: createInteractions(24097)},
			{Frame: 24193, Image: "/images/new_girl/s4e21/scene24193.jpeg", Interactions: createInteractions(24193)},
			{Frame: 25297, Image: "/images/new_girl/s4e21/scene25297.jpeg", Interactions: createInteractions(25297)},
			{Frame: 25633, Image: "/images/new_girl/s4e21/scene25633.jpeg", Interactions: createInteractions(25633)},
			{Frame: 26497, Image: "/images/new_girl/s4e21/scene26497.jpeg", Interactions: createInteractions(26497)},
			{Frame: 26545, Image: "/images/new_girl/s4e21/scene26545.jpeg", Interactions: createInteractions(26535)},
			{Frame: 26881, Image: "/images/new_girl/s4e21/scene26881.jpeg", Interactions: createInteractions(26881)},
			{Frame: 27073, Image: "/images/new_girl/s4e21/scene27073.jpeg", Interactions: createInteractions(27073)},
			{Frame: 27121, Image: "/images/new_girl/s4e21/scene27121.jpeg", Interactions: createInteractions(27121)},
			{Frame: 27169, Image: "/images/new_girl/s4e21/scene27169.jpeg", Interactions: createInteractions(27169)},
			{Frame: 27217, Image: "/images/new_girl/s4e21/scene27217.jpeg", Interactions: createInteractions(27217)},
			{Frame: 27313, Image: "/images/new_girl/s4e21/scene27313.jpeg", Interactions: createInteractions(27313)},
		},
	}
}

func RetrieveEpisodeInfo(variables map[string]string) interface{} {
	log.Debug("handling request to retrieve series information")
	return newGirlEpisodeInfo
}
