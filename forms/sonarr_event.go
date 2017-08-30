package forms

import "time"

//SonarrEventForm ...
type SonarrEventForm struct {
	EventType string `form:"eventtype" json:"eventtype"`
	IsUpgrade bool   `form:"isupgrade" json:"isupgrade"`

	DownloadID     string `form:"downloadid" json:"downloadid"` // the hash of the torrent/NZB file downloaded (used to uniquely identify the download in the download client)
	DownloadClient string `form:"downloadclient" json:"downloadclient"`

	SeriesID       string `form:"seriesid" json:"seriesid"` // internal sonarr id of the series
	SeriesTitle    string `form:"seriestitle" json:"seriestitle"`
	SeriesTVDBID   string `form:"seriestvdbid" json:"seriestvdbid"`
	SeriesTVMazeID string `form:"seriestvmazeid" json:"seriestvmazeid"`
	SeriesIMDBID   string `form:"seriesimdbid" json:"seriesimdbid"`
	SeriesType     string `form:"seriestype" json:"seriestype"` // Anime, Daily, or Standard
	SeriesPath     string `form:"seriespath" json:"seriespath"` // full path to series

	ReleaseEpisodeCount       int64               `form:"releaseepisodecount" json:"releaseepisodecount"`
	ReleaseSeasonNumber       int64               `form:"releaseseasonnumber" json:"releaseseasonnumber"`
	ReleaseEpisodeNumbers     []int64             `form:"releaseepisodenumbers" json:"releaseepisodenumbers"`
	ReleaseEpisodeAirDates    map[int64]time.Time `form:"releaseepisodeairdates" json:"releaseepisodeairdates"`       // from original network
	ReleaseEpisodeAirDatesUTC map[int64]time.Time `form:"releaseepisodeairdatesutc" json:"releaseepisodeairdatesutc"` // above in UTC
	ReleaseEpisodeTitles      map[int64]string    `form:"releaseepisodetitles" json:"releaseepisodetitles"`           // pipe separated
	ReleaseTitle              string              `form:"releasetitle" json:"releasetitle"`
	ReleaseIndexer            string              `form:"releaseindexer" json:"releaseindexer"`
	ReleaseSize               string              `form:"releasesize" json:"releasesize"`
	ReleaseQuality            string              `form:"releasequality" json:"releasequality"`
	ReleaseQualityVersion     string              `form:"releasequalityversion" json:"releasequalityversion"`
	ReleaseGroup              string              `form:"releasegroup" json:"releasegroup"`

	EpisodeFileID                 string              `form:"episodefileid" json:"episodefileid"` // internal sonarr id
	EpisodeFileRelativePath       string              `form:"episodefilerelativepath" json:"episodefilerelativepath"`
	EpisodeFilePath               string              `form:"episodefilepath" json:"episodefilepath"`
	EpisodeFileEpisodeCount       int64               `form:"episodefileepisodecount" json:"episodefileepisodecount"`
	EpisodeFileSeasonNumber       int64               `form:"episodefileseasonnumber" json:"episodefileseasonnumber"`
	EpisodeFileEpisodeNumbers     []int64             `form:"episodefileepisodenumbers" json:"episodefileepisodenumbers"`
	EpisodeFileEpisodeAirDates    map[int64]time.Time `form:"episodefileepisodeairdates" json:"episodefileepisodeairdates"`
	EpisodeFileEpisodeAirDatesUTC map[int64]time.Time `form:"episodefileepisodeairdatesutc" json:"episodefileepisodeairdatesutc"`
	EpisodeFileEpisodeTitles      map[int64]string    `form:"episodefileepisodetitles" json:"episodefileepisodetitles"`
	EpisodeFileQuality            string              `form:"episodefilequality" json:"episodefilequality"`
	EpisodeFileQualityVersion     string              `form:"episodefilequalityversion" json:"episodefilequalityversion"`
	EpisodeFileReleaseGroup       string              `form:"episodefilereleasegroup" json:"episodefilereleasegroup"`
	EpisodeFileSceneName          string              `form:"episodefilescenename" json:"episodefilescenename"`
	EpisodeFileSourcePath         string              `form:"episodefilesourcepath" json:"episodefilesourcepath"`
	EpisodeFileSourceFolder       string              `form:"episodefilesourcefolder" json:"episodefilesourcefolder"`

	DeletedRelativePaths string `form:"deletedrelativepaths" json:"deletedrelativepaths"`
	DeletedPaths         string `form:"deletedpaths" json:"deletedpaths"`
}
