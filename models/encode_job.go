package models

import (
	"time"
)

// EncodeJob describes an encoding job
type EncodeJob struct {
	UUID    string
	KitsuID string
	TVDBID  string

	timeStarted  time.Time
	timeFinished time.Time

	SourceReleaseName string
	SourceResolution  string // 720p, 1080p, etc
	SourceType        string // BluRay Rip, WEB-DL, etc

	EncodeCommand
}
