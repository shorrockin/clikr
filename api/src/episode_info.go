package main

import "github.com/shorrockin/clikr/log"

type EpisodeInfo struct {
	slug    string      `json:"slug"`
	season  uint        `json:"season"`
	episode uint        `json:"episode"`
	fps     uint        `json:"fps"`
	scenes  []SceneInfo `json:"scenes"`
}

type SceneInfo struct {
	image string `json:"json"`
	frame uint   `json:"frame"`
}

func RetrieveEpisodeInfo(variables map[string]string) interface{} {
	log.Debug("handling request to retrieve series information")

	return EpisodeInfo{
		slug:    "new_girl",
		episode: 21,
		season:  4,
		fps:     24,
		scenes: []SceneInfo{
			{frame: 49, image: "/images/new_girl/s4e21/scene00049.jpeg"},
			{frame: 145, image: "/images/new_girl/s4e21/scene00145.jpeg"},
			{frame: 193, image: "/images/new_girl/s4e21/scene00193.jpeg"},
			{frame: 289, image: "/images/new_girl/s4e21/scene00289.jpeg"},
			{frame: 673, image: "/images/new_girl/s4e21/scene00673.jpeg"},
			{frame: 721, image: "/images/new_girl/s4e21/scene00721.jpeg"},
			{frame: 817, image: "/images/new_girl/s4e21/scene00817.jpeg"},
			{frame: 1057, image: "/images/new_girl/s4e21/scene01057.jpeg"},
			{frame: 2017, image: "/images/new_girl/s4e21/scene02017.jpeg"},
			{frame: 2257, image: "/images/new_girl/s4e21/scene02257.jpeg"},
			{frame: 3025, image: "/images/new_girl/s4e21/scene03025.jpeg"},
			{frame: 3745, image: "/images/new_girl/s4e21/scene03745.jpeg"},
			{frame: 3841, image: "/images/new_girl/s4e21/scene03841.jpeg"},
			{frame: 4513, image: "/images/new_girl/s4e21/scene04513.jpeg"},
			{frame: 4657, image: "/images/new_girl/s4e21/scene04657.jpeg"},
			{frame: 4945, image: "/images/new_girl/s4e21/scene04945.jpeg"},
			{frame: 4993, image: "/images/new_girl/s4e21/scene04993.jpeg"},
			{frame: 5233, image: "/images/new_girl/s4e21/scene05233.jpeg"},
			{frame: 5569, image: "/images/new_girl/s4e21/scene05569.jpeg"},
			{frame: 5665, image: "/images/new_girl/s4e21/scene05665.jpeg"},
			{frame: 6529, image: "/images/new_girl/s4e21/scene06529.jpeg"},
			{frame: 7201, image: "/images/new_girl/s4e21/scene07201.jpeg"},
			{frame: 7585, image: "/images/new_girl/s4e21/scene07585.jpeg"},
			{frame: 7681, image: "/images/new_girl/s4e21/scene07681.jpeg"},
			{frame: 8929, image: "/images/new_girl/s4e21/scene08929.jpeg"},
			{frame: 9793, image: "/images/new_girl/s4e21/scene09793.jpeg"},
			{frame: 10081, image: "/images/new_girl/s4e21/scene10081.jpeg"},
			{frame: 11089, image: "/images/new_girl/s4e21/scene11089.jpeg"},
			{frame: 11185, image: "/images/new_girl/s4e21/scene11185.jpeg"},
			{frame: 11329, image: "/images/new_girl/s4e21/scene11329.jpeg"},
			{frame: 11521, image: "/images/new_girl/s4e21/scene11521.jpeg"},
			{frame: 12433, image: "/images/new_girl/s4e21/scene12433.jpeg"},
			{frame: 13201, image: "/images/new_girl/s4e21/scene13201.jpeg"},
			{frame: 13441, image: "/images/new_girl/s4e21/scene13441.jpeg"},
			{frame: 14257, image: "/images/new_girl/s4e21/scene14257.jpeg"},
			{frame: 14401, image: "/images/new_girl/s4e21/scene14401.jpeg"},
			{frame: 15505, image: "/images/new_girl/s4e21/scene15505.jpeg"},
			{frame: 17089, image: "/images/new_girl/s4e21/scene17089.jpeg"},
			{frame: 18145, image: "/images/new_girl/s4e21/scene18145.jpeg"},
			{frame: 18193, image: "/images/new_girl/s4e21/scene18193.jpeg"},
			{frame: 18289, image: "/images/new_girl/s4e21/scene18289.jpeg"},
			{frame: 18337, image: "/images/new_girl/s4e21/scene18337.jpeg"},
			{frame: 18913, image: "/images/new_girl/s4e21/scene18913.jpeg"},
			{frame: 21025, image: "/images/new_girl/s4e21/scene21025.jpeg"},
			{frame: 21745, image: "/images/new_girl/s4e21/scene21745.jpeg"},
			{frame: 22801, image: "/images/new_girl/s4e21/scene22801.jpeg"},
			{frame: 23473, image: "/images/new_girl/s4e21/scene23473.jpeg"},
			{frame: 24097, image: "/images/new_girl/s4e21/scene24097.jpeg"},
			{frame: 24193, image: "/images/new_girl/s4e21/scene24193.jpeg"},
			{frame: 25297, image: "/images/new_girl/s4e21/scene25297.jpeg"},
			{frame: 25633, image: "/images/new_girl/s4e21/scene25633.jpeg"},
			{frame: 26497, image: "/images/new_girl/s4e21/scene26497.jpeg"},
			{frame: 26545, image: "/images/new_girl/s4e21/scene26545.jpeg"},
			{frame: 26881, image: "/images/new_girl/s4e21/scene26881.jpeg"},
			{frame: 27073, image: "/images/new_girl/s4e21/scene27073.jpeg"},
			{frame: 27121, image: "/images/new_girl/s4e21/scene27121.jpeg"},
			{frame: 27169, image: "/images/new_girl/s4e21/scene27169.jpeg"},
			{frame: 27217, image: "/images/new_girl/s4e21/scene27217.jpeg"},
			{frame: 27313, image: "/images/new_girl/s4e21/scene27313.jpeg"},
		},
	}

}
