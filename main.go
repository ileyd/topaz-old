package main

import (
	"log"

	sonarrClient "github.com/ileyd/topaz/sonarr"
)

func main() {
	sonarrClient.InitSonarrClient("http://localhost:8989", "7cda98a81bc44aa68d9447d1957cb29b")
	log.Println("main error", UpdateSeriesFromSonarr())
}
