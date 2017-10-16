package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/ileyd/sonarr"
	"github.com/ileyd/topaz/models"
	sonarrClient "github.com/ileyd/topaz/sonarr"
)

type SonarrEpisodeInfo struct {
	Episode     sonarr.Episode
	EpisodeFile sonarr.EpisodeFile
}

var seriesModel = new(models.SeriesModel)

func findEpisodeFromEpisodeFile(episodes []sonarr.Episode, efID int) (seasonNumber, episodeNumber int, err error) {
	var episode sonarr.Episode
	for _, e := range episodes {
		if e.EpisodeFileID != efID {
			continue
		}
		episode = e
		break
	}
	if episode.SeriesID == 0 {
		return -1, -1, errors.New("Unable to find matching episode from provided episodes and episode file ID")
	}
	return episode.SeasonNumber, episode.EpisodeNumber, nil
}

func UpdateSeriesFromSonarr() (err error) {
	// get all series so that we may loop through them
	series, err := sonarrClient.GetClient().GetAllSeries()
	if err != nil {
		return err
	}
	/* DEBUG */
	seriesJSON, err := json.Marshal(series)
	log.Println("seriesJSON err", err)
	log.Println("seriesJSON", string(seriesJSON))
	/* END DEBUG */

	// iterte through all series
	for _, s := range series {
		// get all episodes so that we may find the corresponding episode information for episodeFiles
		log.Println("using series ID", s.ID)
		episodes, err := sonarrClient.GetClient().GetAllEpisodes(s.ID)
		if err != nil {
			return err
		}
		// get all episodeFiles so that we may iterate through them all
		episodeFiles, err := sonarrClient.GetClient().GetAllEpisodeFiles(s.ID)
		if err != nil {
			return err
		}

		// iterate through all episodeFiles
		for _, ef := range episodeFiles {
			// find matching episode and season number for a gien episodeFile
			seasonNumber, episodeNumber, err := findEpisodeFromEpisodeFile(episodes, ef.ID)
			if err != nil {
				return err
			}
			// get the ID for the database object representing the series we are concerned with
			dbSeries, err := seriesModel.GetOne("tvdbID", s.TvdbID)
			dbSeriesID := dbSeries.ID

			var media = models.Media{
				SeriesID:      dbSeriesID,
				SeasonNumber:  seasonNumber,
				EpisodeNumber: episodeNumber,
				Release: models.Release{
					Quality:        ef.Quality.Quality.Name,
					QualityVersion: ef.Quality.Quality.ID,
					ReleaseGroup:   "",
					ReleaseName:    ef.SceneName,
					Indexer:        "",
					Size:           ef.Size,
				},
				Path: ef.Path,
				URL:  "todo",
			}

			jsonValue, err := json.Marshal(media)
			log.Println(string(jsonValue))
		}

	}

	return nil
}
