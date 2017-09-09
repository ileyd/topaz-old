package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/ileyd/topaz/db"
	"github.com/ileyd/topaz/models"
	"github.com/ileyd/topaz/utils"
	uuid "github.com/satori/go.uuid"
)

var eventModel = new(models.SonarrEventModel)
var jobModel = new(models.JobModel)
var seriesModel = new(models.SeriesModel)

func HandleSonarrEventRegistration(event models.SonarrEvent) (err error) {
	eventModel.Create(event)
	jobModel.Create(event.Job())

	if event.EventType == models.SonarrEventDownloadBegin {
		return
	}

	fmt.Println("step3")

	extension := filepath.Ext(event.EpisodeFile.RelativePath)
	dir := filepath.Dir(event.Series.Path + "/" + event.EpisodeFile.RelativePath)
	if extension == "mkv" || extension == ".mkv" {
		utils.RemuxMKVToMP4(dir, filepath.Base(event.EpisodeFile.RelativePath))
	}

	fmt.Println("step4")

	series, err2 := seriesModel.GetOne("tvdbID", event.Series.TvdbID)
	if err2 == db.ErrNotFound {
		series = models.Series{}
		series.TVDBID = event.Series.TvdbID
		series.KitsuID, err = utils.GetKitsuIDByTitle(event.Series.Title) // unhandled error
		series.CanonicalTitle = event.Series.Title
	}
	fmt.Println("step5")
	var seasonNumber = event.Episodes[0].SeasonNumber
	if _, ok := series.Seasons[seasonNumber]; !ok {
		var season models.Season
		season.SeasonNumber = seasonNumber
		season.SeriesID = series.ID
		season.Episodes = make(map[int]models.Episode)
		if series.Seasons == nil {
			series.Seasons = make(map[int]models.Season)
		}
		series.Seasons[seasonNumber] = season
	}
	var episodeNumber = event.Episodes[0].EpisodeNumber
	if _, ok := series.Seasons[seasonNumber].Episodes[episodeNumber]; !ok {
		var episode models.Episode
		episode.SeasonNumber = seasonNumber
		episode.EpisodeNumber = episodeNumber
		episode.SeriesID = series.ID
		episode.Media = make(map[string]models.Media)
		series.Seasons[seasonNumber].Episodes[episodeNumber] = episode
	}
	series.Seasons[seasonNumber].Episodes[episodeNumber].Media[uuid.NewV4().String()] = models.Media{
		SeriesID:      series.ID,
		SeasonNumber:  seasonNumber,
		EpisodeNumber: episodeNumber,
		Path:          event.EpisodeFile.Path,
		Release: models.Release{
			Quality:        event.EpisodeFile.Quality,
			QualityVersion: event.EpisodeFile.QualityVersion,
			ReleaseGroup:   event.EpisodeFile.ReleaseGroup,
			ReleaseName:    event.EpisodeFile.SceneName,
		},
	}
	if err2 == db.ErrNotFound {
		err = seriesModel.Create(series)
	} else {
		err = seriesModel.Update(series)
	}

	return err
}
