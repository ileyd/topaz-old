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

func InitSonarrEventRegistration() {
	go func() {
		for event := range sonarrEventRegistrationChannel {
			log.Println(HandleSonarrEventRegistration(event))
		}
	}()
}

func TriggerSonarrEventRegistration(event models.SonarrEvent) (err error) {
	// send in a goroutine so the event returns immediately
	go func() {
		sonarrEventRegistrationChannel <- event
	}()
	// no error checking, probably wont be a problem unless
	// something explodes or the api is used incorrectly
	return nil
}

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
		time.Sleep(time.Second * 15) // conversion seems to get triggered too early
		utils.RemuxMKVToMP4(dir, event.Series.Path+"/"+event.EpisodeFile.RelativePath)
	}

	return err
}
