package utils

import (
	"github.com/ileyd/topaz/models"
	"github.com/ileyd/topaz/sonarr"
)

func UpdateSeriesFromSonarr() (err error) {
	series, err := sonarr.GetClient().GetAllSeries()
	if err != nil {
		return err
	}

	for _, s := range series {
		episodes := make(map[int]map[int]models.Media)

		episodeFiles, err := sonarr.GetClient().GetAllEpisodeFiles(s.ID)
		if err != nil {
			return err
		}

	}
}
