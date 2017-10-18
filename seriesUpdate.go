package main

import (
	"errors"
	"log"

	"github.com/ileyd/sonarr"
	uuid "github.com/satori/go.uuid"
)

var seriesModel = new(SeriesModel)
var seasonModel = new(SeasonModel)
var episodeModel = new(EpisodeModel)
var mediaModel = new(MediaModel)

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

	var continues = 0

	// iterte through all series
	for _, s := range series {
		// get all episodes so that we may find the corresponding episode information for episodeFiles
		log.Println("using series ID", s.ID)
		episodes, err := sonarrClient.GetAllEpisodes(s.ID)
		if err != nil {
			log.Println(err)
			continues++
			continue
		}
		// get all episodeFiles so that we may iterate through them all
		episodeFiles, err := sonarrClient.GetAllEpisodeFiles(s.ID)
		if err != nil {
			log.Println(err)
			continues++
			continue
		}

		// iterate through all episodeFiles
		for _, ef := range episodeFiles {
			// find matching episode and season number for a gien episodeFile
			seasonNumber, episodeNumber, err := findEpisodeFromEpisodeFile(episodes, ef.ID)
			if err != nil {
				log.Println(err)
				continues++
				continue
			}
			// get the ID for the database object representing the series we are concerned with
			dbSeries, err := seriesModel.CreateIfNotExists(s)
			if err != nil {
				log.Println("dbSeries", err)
				continues++
				continue
			}
			dbSeriesID := dbSeries.ID
			err = seasonModel.CreateIfNotExists(dbSeries, seasonNumber)
			if err != nil {
				log.Println("seasonModel CINE", err)
				continues++
				continue
			}
			err = episodeModel.CreateIfNotExists(dbSeries, seasonNumber, episodeNumber)
			if err != nil {
				log.Println("episodeModel CINE", err)
				continues++
				continue
			}

			url, err := GenerateB2URL(ef.Path)
			if err != nil {
				log.Println(err)
				continues++
				continue
			}

			var media = Media{
				SeriesID:      dbSeriesID,
				SeasonNumber:  seasonNumber,
				EpisodeNumber: episodeNumber,
				UUID:          uuid.NewV4(),
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

			mediaModel.Add(media)

		}

	}

	log.Println("Skipped", continues, "episodes due to errors")
	return nil
}
