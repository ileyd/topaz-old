package models

import (
	"github.com/ileyd/topaz/db"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Media describes a media file we have in storage
type Media struct {
	ID            string `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID          string `json:"uuid" bson:"uuid"`
	SeriesID      string `json:"seriesID" bson:"seriesID"`
	SeasonNumber  int    `json:"seasonNumber" bson:"seasonNumber"`
	EpisodeNumber int    `json:"episodeNumber" bson:"episodeNumber"`

	Release Release `json:"release" bson:"release"`

	Path string `json:"path" bson:"path"`
	URL  string `json:"url" bson:"url"`
}

// MediaModel is used to group model functions relating to Media objects
type MediaModel struct{}

func (m *MediaModel) Add(me Media) error {
	var s Series
	var sm SeriesModel
	series := db.GetDB().C(SeriesCollection)
	err := series.Find(bson.M{"_id": me.SeriesID}).One(&s)
	if err != nil {
		return err
	}
	s.Seasons[me.SeasonNumber].Episodes[me.EpisodeNumber].Media[me.UUID] = me
	if me.UUID == "" {
		me.UUID = uuid.NewV4().String()
	}
	return sm.Update(s)
}

func (m *MediaModel) Update(me Media) error {
	return m.Add(me)
}

func (m *MediaModel) Delete(me Media) error {
	var s Series
	var sm SeriesModel
	series := db.GetDB().C(SeriesCollection)
	err := series.Find(bson.M{"_id": me.SeriesID}).One(&s)
	if err != nil {
		return err
	}
	delete(s.Seasons[me.SeasonNumber].Episodes[me.EpisodeNumber].Media, me.UUID)
	return sm.Update(s)
}
