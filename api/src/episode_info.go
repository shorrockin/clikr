package main

import "github.com/shorrockin/clikr/log"

type EpisodeInfo struct {
	Slug    string      `json:"slug"`
	Season  uint        `json:"season"`
	Episode uint        `json:"episode"`
	FPS     uint        `json:"fps"`
	Scenes  []SceneInfo `json:"scenes"`
}

type SceneInfo struct {
	Image string `json:"json"`
	Frame uint   `json:"frame"`
}

func RetrieveEpisodeInfo(variables map[string]string) interface{} {
	log.Debug("handling request to retrieve series information")

	return EpisodeInfo{
		Slug:    "new_girl",
		Episode: 21,
		Season:  4,
		FPS:     24,
		Scenes: []SceneInfo{
			{Frame: 49, Image: "/images/new_girl/s4e21/scene00049.jpeg"},
			{Frame: 145, Image: "/images/new_girl/s4e21/scene00145.jpeg"},
			{Frame: 193, Image: "/images/new_girl/s4e21/scene00193.jpeg"},
			{Frame: 289, Image: "/images/new_girl/s4e21/scene00289.jpeg"},
			{Frame: 673, Image: "/images/new_girl/s4e21/scene00673.jpeg"},
			{Frame: 721, Image: "/images/new_girl/s4e21/scene00721.jpeg"},
			{Frame: 817, Image: "/images/new_girl/s4e21/scene00817.jpeg"},
			{Frame: 1057, Image: "/images/new_girl/s4e21/scene01057.jpeg"},
			{Frame: 2017, Image: "/images/new_girl/s4e21/scene02017.jpeg"},
			{Frame: 2257, Image: "/images/new_girl/s4e21/scene02257.jpeg"},
			{Frame: 3025, Image: "/images/new_girl/s4e21/scene03025.jpeg"},
			{Frame: 3745, Image: "/images/new_girl/s4e21/scene03745.jpeg"},
			{Frame: 3841, Image: "/images/new_girl/s4e21/scene03841.jpeg"},
			{Frame: 4513, Image: "/images/new_girl/s4e21/scene04513.jpeg"},
			{Frame: 4657, Image: "/images/new_girl/s4e21/scene04657.jpeg"},
			{Frame: 4945, Image: "/images/new_girl/s4e21/scene04945.jpeg"},
			{Frame: 4993, Image: "/images/new_girl/s4e21/scene04993.jpeg"},
			{Frame: 5233, Image: "/images/new_girl/s4e21/scene05233.jpeg"},
			{Frame: 5569, Image: "/images/new_girl/s4e21/scene05569.jpeg"},
			{Frame: 5665, Image: "/images/new_girl/s4e21/scene05665.jpeg"},
			{Frame: 6529, Image: "/images/new_girl/s4e21/scene06529.jpeg"},
			{Frame: 7201, Image: "/images/new_girl/s4e21/scene07201.jpeg"},
			{Frame: 7585, Image: "/images/new_girl/s4e21/scene07585.jpeg"},
			{Frame: 7681, Image: "/images/new_girl/s4e21/scene07681.jpeg"},
			{Frame: 8929, Image: "/images/new_girl/s4e21/scene08929.jpeg"},
			{Frame: 9793, Image: "/images/new_girl/s4e21/scene09793.jpeg"},
			{Frame: 10081, Image: "/images/new_girl/s4e21/scene10081.jpeg"},
			{Frame: 11089, Image: "/images/new_girl/s4e21/scene11089.jpeg"},
			{Frame: 11185, Image: "/images/new_girl/s4e21/scene11185.jpeg"},
			{Frame: 11329, Image: "/images/new_girl/s4e21/scene11329.jpeg"},
			{Frame: 11521, Image: "/images/new_girl/s4e21/scene11521.jpeg"},
			{Frame: 12433, Image: "/images/new_girl/s4e21/scene12433.jpeg"},
			{Frame: 13201, Image: "/images/new_girl/s4e21/scene13201.jpeg"},
			{Frame: 13441, Image: "/images/new_girl/s4e21/scene13441.jpeg"},
			{Frame: 14257, Image: "/images/new_girl/s4e21/scene14257.jpeg"},
			{Frame: 14401, Image: "/images/new_girl/s4e21/scene14401.jpeg"},
			{Frame: 15505, Image: "/images/new_girl/s4e21/scene15505.jpeg"},
			{Frame: 17089, Image: "/images/new_girl/s4e21/scene17089.jpeg"},
			{Frame: 18145, Image: "/images/new_girl/s4e21/scene18145.jpeg"},
			{Frame: 18193, Image: "/images/new_girl/s4e21/scene18193.jpeg"},
			{Frame: 18289, Image: "/images/new_girl/s4e21/scene18289.jpeg"},
			{Frame: 18337, Image: "/images/new_girl/s4e21/scene18337.jpeg"},
			{Frame: 18913, Image: "/images/new_girl/s4e21/scene18913.jpeg"},
			{Frame: 21025, Image: "/images/new_girl/s4e21/scene21025.jpeg"},
			{Frame: 21745, Image: "/images/new_girl/s4e21/scene21745.jpeg"},
			{Frame: 22801, Image: "/images/new_girl/s4e21/scene22801.jpeg"},
			{Frame: 23473, Image: "/images/new_girl/s4e21/scene23473.jpeg"},
			{Frame: 24097, Image: "/images/new_girl/s4e21/scene24097.jpeg"},
			{Frame: 24193, Image: "/images/new_girl/s4e21/scene24193.jpeg"},
			{Frame: 25297, Image: "/images/new_girl/s4e21/scene25297.jpeg"},
			{Frame: 25633, Image: "/images/new_girl/s4e21/scene25633.jpeg"},
			{Frame: 26497, Image: "/images/new_girl/s4e21/scene26497.jpeg"},
			{Frame: 26545, Image: "/images/new_girl/s4e21/scene26545.jpeg"},
			{Frame: 26881, Image: "/images/new_girl/s4e21/scene26881.jpeg"},
			{Frame: 27073, Image: "/images/new_girl/s4e21/scene27073.jpeg"},
			{Frame: 27121, Image: "/images/new_girl/s4e21/scene27121.jpeg"},
			{Frame: 27169, Image: "/images/new_girl/s4e21/scene27169.jpeg"},
			{Frame: 27217, Image: "/images/new_girl/s4e21/scene27217.jpeg"},
			{Frame: 27313, Image: "/images/new_girl/s4e21/scene27313.jpeg"},
		},
	}

}
