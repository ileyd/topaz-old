package utils

import (
	"github.com/ileyd/topaz/models"
	"github.com/ileyd/topaz/sonarr"
)

var seriesModel = new(models.SeriesModel)

func UpdateSeriesFromSonarr() (err error) {
	// get all series to iterate over
	series, err := sonarr.GetClient().GetAllSeries()
	if err != nil {
		return err
	}

	// iterate over all series
	for _, s := range series {
		seriesObject, err := seriesModel.CreateSeriesIfNotExists(s)
		if err != nil {
			return err
		}

		episodes, err := sonarr.GetClient().GetAllEpisodes(s.ID)
		if err != nil {
			return err
		}

		// get all episode files for current series
		episodeFiles, err := sonarr.GetClient().GetAllEpisodeFiles(s.ID)
		if err != nil {
			return err
		}

		var oldEpnums = make(map[int]bool)
		// iterate over all episodeFiles for current series
		for _, ef := range episodeFiles {
			if err = CreateSeasonIfNotExists(seriesObject, ef.SeasonNumber); err != nil {
				return err
			}

			var epNum int // holds decided-on episode number
			// iterate over all episodes for current series until we fine a suitable candidate
			for _, e := range episodes {
				if e.EpisodeFileID != ef.ID || oldEpnums[e.EpisodeNumber] {
					continue
				}
				epNum = e.EpisodeNumber
				oldEpnums[epNum] = true
			}

		}
	}

	return nil
}
