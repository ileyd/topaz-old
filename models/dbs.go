package models

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"

	"github.com/ileyd/kitsu"
	"github.com/nemith/tvdb"
)

func GetTVDBIDByTitle(title string) (id int, err error) {
	client := tvdb.NewClient("1AF5B7DB27302EB5")
	series, err := client.SearchSeries(title, "en")
	id = series[0].ID
	return id, err
}

func GetKitsuIDByTitle(title string) (int, error) {
	page, err := kitsu.GetAnimePage(url.QueryEscape("anime/?filter[text]=" + title))
	log.Println("GKIBT-1", err, title)
	j, _ := json.Marshal(page)
	log.Println("GKIBT page", string(j))
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(page.Data[0].ID, 10, 64)
	log.Println("GKIBT-2", err)
	intID := int(id)
	return intID, err
}
