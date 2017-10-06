package handlers

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ileyd/topaz/models"
	"github.com/ileyd/topaz/utils"
	uuid "github.com/satori/go.uuid"
)

var eventModel = new(models.SonarrEventModel)
var jobModel = new(models.JobModel)
var seriesModel = new(models.SeriesModel)

// HandleSonarrEventRegistration ... self exlanatory
func HandleSonarrEventRegistration(event models.SonarrEvent) (err error) {
	eventModel.Create(event)     // register the event
	jobModel.Create(event.Job()) // register a corresponding job for the even

	// if this is a download started event rather than a download complete event, we don't need to do anything more
	if event.EventType == models.SonarrEventDownloadBegin {
		return
	}

	fmt.Println("step3")

	// if is mkv remux to mp4
	extension := filepath.Ext(event.EpisodeFile.RelativePath)
	dir := filepath.Dir(event.Series.Path + "/" + event.EpisodeFile.RelativePath)
	if extension == "mkv" || extension == ".mkv" {
		utils.RemuxMKVToMP4(dir, event.Series.Path+"/"+event.EpisodeFile.RelativePath)
	}

	fmt.Println("step4")

	series, err := seriesModel.GetOne("tvdbID", event.Series.TvdbID)
	if err != nil { // if there is an error, series probably doesn't exist so lets create it
		series = models.Series{}
		series.TVDBID = event.Series.TvdbID
		series.KitsuID, err = utils.GetKitsuIDByTitle(event.Series.Title) // unhandled error
		series.CanonicalTitle = event.Series.Title
		err = seriesModel.Create(series)                                // unhandled error
		series, err = seriesModel.GetOne("tvdbID", event.Series.TvdbID) // unhandled error
	}
	fmt.Println("step5")
	var seasonNumber = event.Episodes[0].SeasonNumber
	// if season object doesn't exist create it
	if _, ok := series.Seasons[seasonNumber]; !ok {
		var season models.Season
		season.SeasonNumber = seasonNumber
		season.SeriesID = series.ID
		season.Episodes = make(map[int]models.Episode)
		// if seasons map doesn't exist, make it
		if series.Seasons == nil {
			series.Seasons = make(map[int]models.Season)
		}
		series.Seasons[seasonNumber] = season
	}
	var episodeNumber = event.Episodes[0].EpisodeNumber
	// if episode object doesn't exist, create it
	if _, ok := series.Seasons[seasonNumber].Episodes[episodeNumber]; !ok {
		var episode models.Episode
		episode.SeasonNumber = seasonNumber
		episode.EpisodeNumber = episodeNumber
		episode.SeriesID = series.ID
		episode.Media = make(map[string]models.Media)
		series.Seasons[seasonNumber].Episodes[episodeNumber] = episode
	}

	mediaUUID := uuid.NewV4().String()
	series.Seasons[seasonNumber].Episodes[episodeNumber].Media[mediaUUID] = models.Media{
		SeriesID:      series.ID,
		UUID:          mediaUUID,
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

	err = seriesModel.Update(series)

	log.Println("step6", series, "==", event)

	return err
}
