package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/ileyd/sonarr"
)

var seriesModel = new(SeriesModel)

// findEpisodeFromEpisodeFile accepts an episodes slice and an episodeFile ID and finds the episode in the slice that corresponds to the given episodeFile ID
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

// updateSeriesFromSonarr iterates through all series in Sonarr's database and populates our database based on a processed form of this information
func updateSeriesFromSonarr() (err error) {
	// get all series so that we may loop through them
	series, err := sonarrClient.GetAllSeries()
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
		episodes, err := sonarrClient.GetAllEpisodes(s.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		// get all episodeFiles so that we may iterate through them all
		episodeFiles, err := sonarrClient.GetAllEpisodeFiles(s.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		// iterate through all episodeFiles
		for _, ef := range episodeFiles {
			// find matching episode and season number for a gien episodeFile
			seasonNumber, episodeNumber, err := findEpisodeFromEpisodeFile(episodes, ef.ID)
			if err != nil {
				log.Println(err)
				continue
			}
			// get the ID for the database object representing the series we are concerned with
			dbSeries, err := seriesModel.CreateIfNotExists(s)
			if err != nil {
				log.Println("dbSeries", err)
				// continue
			}
			dbSeriesID := dbSeries.ID

			url, err := GenerateB2URL(ef.Path)
			if err != nil {
				log.Println(err)
				continue
			}

			var media = Media{
				SeriesID:      dbSeriesID,
				SeasonNumber:  seasonNumber,
				EpisodeNumber: episodeNumber,
				Release: Release{
					Quality:        ef.Quality.Quality.Name,
					QualityVersion: ef.Quality.Quality.ID,
					ReleaseGroup:   "",
					ReleaseName:    ef.SceneName,
					Indexer:        "",
					Size:           ef.Size,
				},
				Path: ef.Path,
				URL:  *url,
			}

			jsonValue, err := json.Marshal(media)
			log.Println(string(jsonValue))
		}

	}

	return nil
}
