package utils

import (
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
	page, err := kitsu.GetAnimePage(title)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(page.Data[0].ID, 10, 64)
	intID := int(id)
	return intID, err
}
