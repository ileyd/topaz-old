package sonarr

import "time"

type SonarrFolder struct {
	Path            string `json:"path"`
	FreeSpace       int64  `json:"freeSpace"`
	UnmappedFolders []struct {
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"unmappedFolders"`
	ID int `json:"id"`
}

type SonarrSeries struct {
	Title             string        `json:"title,omitempty"`
	AlternateTitles   []interface{} `json:"alternateTitles,omitempty"`
	SortTitle         string        `json:"sortTitle,omitempty"`
	TotalEpisodeCount int           `json:"totalEpisodeCount,omitempty"`

	SeasonCount      int    `json:"seasonCount,omitempty"`
	EpisodeCount     int    `json:"episodeCount,omitempty"`
	EpisodeFileCount int    `json:"episodeFileCount,omitempty"`
	Status           string `json:"status,omitempty"`
	Overview         string `json:"overview,omitempty"`
	Network          string `json:"network,omitempty"`
	Images           []struct {
		CoverType string `json:"coverType,omitempty"`
		URL       string `json:"url,omitempty"`
	} `json:"images,omitempty"`
	RemotePoster string `json:"remotePoster,omitempty"`
	Seasons      []struct {
		SeasonNumber int  `json:"seasonNumber,omitempty"`
		Monitored    bool `json:"monitored,omitempty"`
	} `json:"seasons,omitempty"`
	Year              int           `json:"year,omitempty"`
	QualityProfileID  int           `json:"qualityProfileId,omitempty"`
	SeasonFolder      bool          `json:"seasonFolder,omitempty"`
	Monitored         bool          `json:"monitored,omitempty"`
	UseSceneNumbering bool          `json:"useSceneNumbering,omitempty"`
	Runtime           int           `json:"runtime,omitempty"`
	TvdbID            int           `json:"tvdbId,omitempty"`
	TvRageID          int           `json:"tvRageId,omitempty"`
	SeriesType        string        `json:"seriesType,omitempty"`
	CleanTitle        string        `json:"cleanTitle,omitempty"`
	ImdbID            string        `json:"imdbId,omitempty"`
	TitleSlug         string        `json:"titleSlug,omitempty"`
	Path              string        `json:"path,omitempty"`
	FirstAired        time.Time     `json:"firstAired,omitempty"`
	LastInfoSync      time.Time     `json:"lastInfoSync,omitempty"`
	SizeOnDisk        int           `json:"sizeOnDisk,omitempty"`
	TvMazeID          int           `json:"tvMazeId,omitempty"`
	Genres            []interface{} `json:"genres,omitempty"`
	Tags              []interface{} `json:"tags,omitempty"`
	Added             time.Time     `json:"added,omitempty"`
	ID                int           `json:"id,omitempty"`
}
