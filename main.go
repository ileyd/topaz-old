package main

import "log"

func main() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	log.Println(config)
	initDb()
	initSonarrClient()
	// log.Println("main error", updateSeriesFromSonarr())
}
