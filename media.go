package main

import (
	"net/url"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Media describes a media file we have in storage
type Media struct {
	ID            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID          uuid.UUID     `json:"uuid" bson:"uuid"`
	SeriesID      bson.ObjectId `json:"seriesID" bson:"seriesID"`
	SeasonNumber  int           `json:"seasonNumber" bson:"seasonNumber"`
	EpisodeNumber int           `json:"episodeNumber" bson:"episodeNumber"`

	Release Release `json:"release" bson:"release"`

	Path string  `json:"path" bson:"path"`
	URL  url.URL `json:"url" bson:"url"`
}

// MediaModel is used to group model functions relating to Media objects
type MediaModel struct{}

// Add adds a new Media object to an existing Episode object
func (m *MediaModel) Add(me Media) error {
	var s Series
	var sm SeriesModel
	series := db.C(SeriesCollection)
	err := series.Find(bson.M{"_id": me.SeriesID}).One(&s)
	if err != nil {
		return err
	}
	s.Seasons[strconv.Itoa(me.SeasonNumber)].Episodes[strconv.Itoa(me.EpisodeNumber)].Media[me.UUID.String()] = me
	if me.UUID.String() == "" {
		me.UUID = uuid.NewV4()
	}
	return sm.Update(s)
}

// Update is an alias for the Add() function
func (m *MediaModel) Update(me Media) error {
	return m.Add(me)
}

// Delete removes a Media object from an existing Episode object
func (m *MediaModel) Delete(me Media) error {
	var s Series
	var sm SeriesModel
	series := db.C(SeriesCollection)
	err := series.Find(bson.M{"_id": me.SeriesID}).One(&s)
	if err != nil {
		return err
	}
	delete(s.Seasons[strconv.Itoa(me.SeasonNumber)].Episodes[strconv.Itoa(me.EpisodeNumber)].Media, me.UUID.String())
	return sm.Update(s)
}
