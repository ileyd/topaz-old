package sonarr

import (
	"github.com/ileyd/sonarr"
)

var Client *sonarr.SonarrClient

func InitSonarrClient(address, apiKey string) (err error) {
	Client, err = sonarr.NewSonarrClient(address, apiKey)
	return err
}

func GetClient() *sonarr.SonarrClient {
	return Client
}
