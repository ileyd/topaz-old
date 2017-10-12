package models

import (
	"github.com/ileyd/topaz/db"
	"gopkg.in/mgo.v2/bson"
)

// Collection names
const (
	SeriesCollection = "series"
)

// Series describes a series we have in our database
type Series struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	KitsuID        int    `json:"kitsuID" bson:"kitsuID"`
	TVDBID         int    `json:"tvdbID" bson:"tvdbID"`
	CanonicalTitle string `json:"canonicalTitle" bson:"canonicalTitle"`

	SeasonCount int               `json:"seasonCount" bson:"seasonCount"`
	Seasons     map[string]Season `json:"seasons" bson:"seasons"`
}

// Season describes an series' season
type Season struct {
	ID           bson.ObjectId             `json:"_id,omitempty" bson:"_id,omitempty"`
	SeriesID     string             `json:"seriesID" bson:"seriesID"`
	SeasonNumber int                `json:"seasonNumber" bson:"seasonNumber"`
	EpisodeCount int                `json:"episodeCount" bson:"episodeCount"`
	Episodes     map[string]Episode `json:"episodes" bson:"episodes"`
}

// Episode describes a media file relating to an anime episode
type Episode struct {
	ID            bson.ObjectId           `json:"_id,omitempty" bson:"_id,omitempty"`
	SeriesID      string           `json:"seriesID" bson:"seriesID"`
	SeasonNumber  int              `json:"seasonNumber" bson:"seasonNumber"`
	EpisodeNumber int              `json:"episodeNumber" bson:"episodeNumber"`
	MediaCount    int              `json:"mediaCount" bson:"mediaCount"`
	Media         map[string]Media `json:"media" bson:"media"`
}

// SeriesModel is used to group model functions relating to Series objects
type SeriesModel struct{}

// Create ...
func (m *SeriesModel) Create(s Series) error {
	series := db.GetDB().C(SeriesCollection)
	return series.Insert(s)
}

// Update ...
func (m *SeriesModel) Update(s Series) error {
	series := db.GetDB().C(SeriesCollection)
	return series.Update(bson.M{"_id": s.ID}, s)
}

// Delete ...
func (m *SeriesModel) Delete(selectorKey string, selectorValue interface{}) error {
	series := db.GetDB().C(SeriesCollection)
	return series.Remove(bson.M{selectorKey: selectorValue})
}

// GetIDFromKitsuID gets the series' mongodb Object ID based on a specified KitsuID value
func (m *SeriesModel) GetIDFromKitsuID(KID string) (ID bson.ObjectId, err error) {
	var s Series
	s, err = m.GetOne("kitsuID", KID)
	return s.ID, err
}

// GetOne ...
func (m *SeriesModel) GetOne(selectorKey string, selectorValue interface{}) (Series, error) {
	series := db.GetDB().C(SeriesCollection)
	var s Series
	err := series.Find(bson.M{selectorKey: selectorValue}).One(&s)
	return s, err
}

// GetAll ...
func (m *SeriesModel) GetAll() ([]Series, error) {
	series := db.GetDB().C(SeriesCollection)
	var s []Series
	err := series.Find(nil).All(&s)
	return s, err
}

// SeasonModel is used to group model functions relating to Season objects
type SeasonModel struct{}

func (m *SeasonModel) Add(s Season) error {
	var srm SeriesModel
	sr, err := srm.GetOne("_id", s.SeriesID)
	if err != nil {
		return err
	}
	sr.Seasons[string(s.SeasonNumber)] = s
	return srm.Update(sr)
}

// EpisodeModel is used to group model functions relating to Episode objects
type EpisodeModel struct{}
