package handlers

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/ileyd/topaz/models"
	"github.com/ileyd/topaz/utils"
	uuid "github.com/satori/go.uuid"
)

var eventModel = new(models.SonarrEventModel)
var jobModel = new(models.JobModel)
var seriesModel = new(models.SeriesModel)

var sonarrEventRegistrationChannel = make(chan models.SonarrEvent)

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
	time.Sleep(time.Second * 15) // conversion seems to get triggered too early
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
		if err != nil {
			return err
		}
		series.CanonicalTitle = event.Series.Title
		err = seriesModel.Create(series) // unhandled error
		if err != nil {
			return err
		}
		series, err = seriesModel.GetOne("tvdbID", event.Series.TvdbID) // unhandled error
		if err != nil {
			return err
		}
	}
	fmt.Println("step5")
	var seasonNumber = event.Episodes[0].SeasonNumber
	// if season object doesn't exist create it
	if _, ok := series.Seasons[string(seasonNumber)]; !ok {
		var season models.Season
		season.SeasonNumber = seasonNumber
		season.SeriesID = series.ID
		season.Episodes = make(map[string]models.Episode)
		// if seasons map doesn't exist, make it
		if series.Seasons == nil {
			series.Seasons = make(map[string]models.Season)
		}
		series.Seasons[string(seasonNumber)] = season
		err = seriesModel.Update(series)
		if err != nil {
			return err
		}
	}
	var episodeNumber = event.Episodes[0].EpisodeNumber
	// if episode object doesn't exist, create it
	if _, ok := series.Seasons[string(seasonNumber)].Episodes[string(episodeNumber)]; !ok {
		var episode models.Episode
		episode.SeasonNumber = seasonNumber
		episode.EpisodeNumber = episodeNumber
		episode.SeriesID = series.ID
		episode.Media = make(map[string]models.Media)
		series.Seasons[string(seasonNumber)].Episodes[string(episodeNumber)] = episode
		err = seriesModel.Update(series)
		if err != nil {
			return err
		}
	}

	mediaUUID := uuid.NewV4().String()
	series.Seasons[string(seasonNumber)].Episodes[string(episodeNumber)].Media[mediaUUID] = models.Media{
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
