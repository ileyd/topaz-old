package main

import (
	"log"

	"github.com/jinzhu/configor"
	"gopkg.in/mgo.v2"
)

// topazConfig defines the configuration required for the topaz application
type topazConfig struct {
	B2     b2Config
	DB     dbConfig
	Sonarr sonarrConfig
}

// dbConfig defines the configuration required to connect to the database
type dbConfig struct {
	Addrs          []string `required:"true"`
	ReplicaSetName string
	AuthDB         string `default:"admin"`
	Username       string `required:"true"`
	Password       string `required:"true"`
	AppDB          string `required:"true" default:"topaz"`
}

// sonarrConfig defines the configuration required to connect to Sonarr
type sonarrConfig struct {
	Addr   string `required:"true" default:"http://localhost:8989"`
	APIKey string `required:"true"`
}

// b2Config defines the configuration required to construct streamable Backblaze B2 URLs
type b2Config struct {
	BaseURL            string `required:"true"`
	StrippedComponents int    `required:"true" default:"2"`
}

var config topazConfig

// loadConfig loads the application configuration from disk
func loadConfig() error {
	err := configor.Load(&config, "config.json")
	if err == nil {
		log.Println("successfully loaded config.json")
	}
	return err
}

// DialInfo returns an mgo.DialInfo object based on the parameters specified in the application configuration
func (c dbConfig) DialInfo() *mgo.DialInfo {
	return &mgo.DialInfo{
		Addrs:          c.Addrs,
		ReplicaSetName: c.ReplicaSetName,
		Database:       c.AuthDB,
		Username:       c.Username,
		Password:       c.Password,
	}
}
