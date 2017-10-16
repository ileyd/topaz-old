package models

// Release describes a "release", a term with specific meaning in the warez scene
type Release struct {
	Quality        string `json:"quality" bson:"quality"`
	QualityVersion int    `json:"qualityVersion" bson:"qualityVersion"`
	ReleaseGroup   string `json:"releaseGroup" bson:"releaseGroup"`
	ReleaseName    string `json:"releaseName" bson:"releaseName"`
	Indexer        string `json:"indexer" bson:"indexer"`
	Size           int64  `json:"size" bson:"size"`
}
