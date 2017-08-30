package models

import "time"

// Sonarr EventTypes
const (
	SonarrEventGrab     = "Grab"
	SonarrEventDownload = "Download"
	SonarrEventUpgrade  = SonarrEventDownload
	SonarrEventRename   = "Rename"
)

// Sonarr SeriesTypes
const (
	SonarrSeriesAnime    = "Anime"
	SonarrSeriesDaily    = "Daily"
	SonarrSeriesStandard = "Standard"
)

// SonarrEvent contains env vars for when triggered to handle an event
type SonarrEvent struct {
	EventType string `json:"eventtype" form:"eventtype"`
	IsUpgrade bool   `json:"isupgrade" form:"isupgrade"`

	DownloadID     string `json:"downloadid" form:"downloadid"` // the hash of the torrent/NZB file downloaded (used to uniquely identify the download in the download client)
	DownloadClient string `json:"downloadclient" form:"downloadclient"`

	SeriesID       string `json:"seriesid" form:"seriesid"` // internal sonarr id of the series
	SeriesTitle    string `json:"seriestitle" form:"seriestitle"`
	SeriesTVDBID   string `json:"seriestvdbid" form:"seriestvdbid"`
	SeriesTVMazeID string `json:"seriestvmazeid" form:"seriestvmazeid"`
	SeriesIMDBID   string `json:"seriesimdbid" form:"seriesimdbid"`
	SeriesType     string `json:"seriestype" form:"seriestype"` // Anime, Daily, or Standard
	SeriesPath     string `json:"seriespath" form:"seriespath"` // full path to series

	ReleaseEpisodeCount       int64               `json:"releaseepisodecount" form:"releaseepisodecount"`
	ReleaseSeasonNumber       int64               `json:"releaseseasonnumber" form:"releaseseasonnumber"`
	ReleaseEpisodeNumbers     []int64             `json:"releaseepisodenumbers" form:"releaseepisodenumbers"`
	ReleaseEpisodeAirDates    map[int64]time.Time `json:"releaseepisodeairdates" form:"releaseepisodeairdates"`       // from original network
	ReleaseEpisodeAirDatesUTC map[int64]time.Time `json:"releaseepisodeairdatesutc" form:"releaseepisodeairdatesutc"` // above in UTC
	ReleaseEpisodeTitles      map[int64]string    `json:"releaseepisodetitles" form:"releaseepisodetitles"`           // pipe separated
	ReleaseTitle              string              `json:"releasetitle" form:"releasetitle"`
	ReleaseIndexer            string              `json:"releaseindexer" form:"releaseindexer"`
	ReleaseSize               string              `json:"releasesize" form:"releasesize"`
	ReleaseQuality            string              `json:"releasequality" form:"releasequality"`
	ReleaseQualityVersion     string              `json:"releasequalityversion" form:"releasequalityversion"`
	ReleaseGroup              string              `json:"releasegroup" form:"releasegroup"`

	EpisodeFileID                 string              `json:"episodefileid" form:"episodefileid"` // internal sonarr id
	EpisodeFileRelativePath       string              `json:"episodefilerelativepath" form:"episodefilerelativepath"`
	EpisodeFilePath               string              `json:"episodefilepath" form:"episodefilepath"`
	EpisodeFileEpisodeCount       int64               `json:"episodefileepisodecount" form:"episodefileepisodecount"`
	EpisodeFileSeasonNumber       int64               `json:"episodefileseasonnumber" form:"episodefileseasonnumber"`
	EpisodeFileEpisodeNumbers     []int64             `json:"episodefileepisodenumbers" form:"episodefileepisodenumbers"`
	EpisodeFileEpisodeAirDates    map[int64]time.Time `json:"episodefileepisodeairdates" form:"episodefileepisodeairdates"`
	EpisodeFileEpisodeAirDatesUTC map[int64]time.Time `json:"episodefileepisodeairdatesutc" form:"episodefileepisodeairdatesutc"`
	EpisodeFileEpisodeTitles      map[int64]string    `json:"episodefileepisodetitles" form:"episodefileepisodetitles"`
	EpisodeFileQuality            string              `json:"episodefilequality" form:"episodefilequality"`
	EpisodeFileQualityVersion     string              `json:"episodefilequalityversion" form:"episodefilequalityversion"`
	EpisodeFileReleaseGroup       string              `json:"episodefilereleasegroup" form:"episodefilereleasegroup"`
	EpisodeFileSceneName          string              `json:"episodefilescenename" form:"episodefilescenename"`
	EpisodeFileSourcePath         string              `json:"episodefilesourcepath" form:"episodefilesourcepath"`
	EpisodeFileSourceFolder       string              `json:"episodefilesourcefolder" form:"episodefilesourcefolder"`

	DeletedRelativePaths string `json:"deletedrelativepaths" form:"deletedrelativepaths"`
	DeletedPaths         string `json:"deletedpaths" form:"deletedpaths"`
}
