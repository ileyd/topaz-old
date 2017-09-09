package models

import (
	"crypto/sha1"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/ileyd/topaz/db"
)

// Sonarr EventTypes
const (
	SonarrEventDownloadBegin    = "Grab"
	SonarrEventDownloadComplete = "Download"
	SonarrEventUpgrade          = "Download"
	SonarrEventRename           = "Rename"
)

// Sonarr SeriesTypes
const (
	SonarrSeriesAnime    = "Anime"
	SonarrSeriesDaily    = "Daily"
	SonarrSeriesStandard = "Standard"
)

// MongoDB collection names
const (
	SonarrEventsCollection = "events"
)

// SonarrEvent describes a Sonarr event
type SonarrEvent struct {
	Episodes []struct {
		ID             int       `json:"id"`
		EpisodeNumber  int       `json:"episodeNumber"`
		SeasonNumber   int       `json:"seasonNumber"`
		Title          string    `json:"title"`
		AirDate        string    `json:"airDate"`
		AirDateUtc     time.Time `json:"airDateUtc"`
		Quality        string    `json:"quality"`
		QualityVersion int       `json:"qualityVersion"`
		ReleaseGroup   string    `json:"releaseGroup"`
		SceneName      string    `json:"sceneName"`
	} `json:"episodes" bson:"episodes"`
	EpisodeFile struct {
		ID             int    `json:"id"`
		RelativePath   string `json:"relativePath"`
		Path           string `json:"path"`
		Quality        string `json:"quality"`
		QualityVersion int    `json:"qualityVersion"`
		ReleaseGroup   string `json:"releaseGroup"`
		SceneName      string `json:"sceneName"`
	} `json:"episodeFile"`
	Release struct {
		Quality        string `json:"quality"`
		QualityVersion int    `json:"qualityVersion"`
		ReleaseGroup   string `json:"releaseGroup"`
		ReleaseTitle   string `json:"releaseTitle"`
		Indexer        string `json:"indexer"`
		Size           int    `json:"size"`
	} `json:"release" bson:"release"`
	IsUpgrade bool   `json:"isUpgrade"`
	EventType string `json:"eventType" bson:"eventType"`
	Series    struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Path   string `json:"path"`
		TvdbID int    `json:"tvdbId"`
	} `json:"series" bson:"series"`
}

// EpisodeInfoID returns a hash of a concatenation of some information that is consistent between Grab/Download Sonarr event pairs
func (e *SonarrEvent) EpisodeInfoID() string {
	var info string
	for _, ep := range e.Episodes {
		info = info + string(ep.ID) + string(ep.EpisodeNumber) + string(ep.SeasonNumber) + ep.ReleaseGroup
	}
	hash := sha1.Sum([]byte(info))
	return string(hash[:20])
}

func (e *SonarrEvent) Job() (job Job) {
	var jm JobModel
	switch e.EventType {
	case SonarrEventDownloadBegin:
		job.Series.CanonicalTitle = e.Series.Title
		job.Series.TVDBID = e.Series.TvdbID
		job.EpisodeInfoID = e.EpisodeInfoID()
		job.TimeStarted = time.Now()
		job.Type = FetchJob
		job.Episodes = make(map[int]map[int]bool)
		for _, ep := range e.Episodes {
			if job.Episodes[ep.SeasonNumber] == nil {
				job.Episodes[ep.SeasonNumber] = make(map[int]bool)
			}
			job.Episodes[ep.SeasonNumber][ep.EpisodeNumber] = true
		}
		break
	case SonarrEventDownloadComplete:
		job, _ = jm.GetByEpisodeInfoID(e.EpisodeInfoID())
		job.Completed = true
		job.TimeFinished = time.Now()
		break
	}

	return job
}

// SonarrEventModel ...
type SonarrEventModel struct{}

// Create ...
func (m *SonarrEventModel) Create(event SonarrEvent) (err error) {
	events := db.GetDB().C(SonarrEventsCollection)
	return events.Insert(event)
}

// GetAll ...
func (m *SonarrEventModel) GetAll() []SonarrEvent {
	var results []SonarrEvent
	err := db.GetDB().C(SonarrEventsCollection).Find(nil).All(&results)
	log.Println(err)
	return results
}

// Get ...
func (m *SonarrEventModel) Get(id int64) SonarrEvent {
	var result SonarrEvent
	err := db.GetDB().C(SonarrEventsCollection).Find(bson.M{"_id": id}).One(&result)
	log.Println(err)
	return result
}
