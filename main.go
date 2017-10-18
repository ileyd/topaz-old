package main

import (
	"log"
)

func main() {
	loadConfig()
	initDb()
	initSonarrClient()
	log.Println("main error", updateSeriesFromSonarr())
}
