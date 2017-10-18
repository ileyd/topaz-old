package main

import (
	"github.com/ileyd/sonarr"
)

var sonarrClient *sonarr.SonarrClient

// initSonarrClient initialises the sonarrClient variable with a new sonarr client
func initSonarrClient() (err error) {
	sonarrClient, err = sonarr.NewSonarrClient(config.Sonarr.Addr, config.Sonarr.APIKey)
	return err
}
