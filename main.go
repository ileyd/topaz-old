package main

import "log"

func main() {
	if err := loadConfig(); err != nil {
		panic(err)
	}

	if err := initSonarrClient(); err != nil {
		panic(err)
	}

	if err := initDb(); err != nil {
		panic(err)
	}

	log.Println("main error", updateSeriesFromSonarr())
}
