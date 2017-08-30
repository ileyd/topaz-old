package models

import "time"

// Season describes an collection of media files relating to an anime season
type Season struct {
	UUID    string // unique collection id
	KitsuID string
	TVDBID  string

	SourcePath string
	TargetPath string

	Quality        string
	SourceType     string
	QualityVersion string

	ReleaseGroup string
	ReleaseName  string

	EpisodeCount   int
	SeasonNumber   int
	EpisodeNumbers []int
	EpisodeTitles  map[int]string
	AirDates       map[int]time.Time
}

// Episode describes a media file relating to an anime episode
type Episode struct {
	UUID    string // unique collection id
	KitsuID string
	TVDBID  string

	SourcePath string
	TargetPath string

	Quality        string
	SourceType     string
	QualityVersion string

	ReleaseGroup string
	ReleaseName  string

	SeasonNumber  int
	EpisodeNumber int
	EpisodeTitle  string
	AirDate       time.Time
}
