package sonarr

import (
	"fmt"

	"github.com/juju/errors"
)

func (sc *SonarrClient) SeriesLookup(term string) ([]SonarrSeries, error) {

	if term == "" {
		return nil, errors.New("No term specified")
	}

	rv := &[]SonarrSeries{}

	err := sc.DoRequest("GET", "series/lookup", map[string]string{"term": term}, nil, rv)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to lookup series")
	}

	return *rv, nil
}

func (sc *SonarrClient) CreateSeries(tvdbId int) (SonarrSeries, error) {
	//first lookup data
	slrs, err := sc.SeriesLookup(fmt.Sprintf("tvdb:%v", tvdbId))

	if len(slrs) == 0 {
		return SonarrSeries{}, errors.Errorf("No series found with tvdbid %v", tvdbId)
	}

	if err != nil {
		return SonarrSeries{}, errors.Annotate(err, "Failed to get series data to add ")
	}

	rootFolder, err := sc.RootFolder()

	if len(rootFolder) == 0 {
		return SonarrSeries{}, errors.Errorf("No root folder found")
	}

	path := rootFolder[0].Path + slrs[0].Title

	//alter some stuff
	slrs[0].QualityProfileID = 1
	slrs[0].Path = path

	var spr_out SonarrSeries

	//then add it
	err = sc.DoRequest("POST", "series", nil, slrs[0], &spr_out)

	if err != nil {
		return SonarrSeries{}, errors.Annotate(err, "Failed to add series")
	}

	return spr_out, nil
}

func (sc *SonarrClient) RootFolder() ([]SonarrFolder, error) {
	var rv []SonarrFolder

	//then add it
	err := sc.DoRequest("GET", "rootfolder", nil, nil, &rv)

	if err != nil {
		return []SonarrFolder{}, errors.Annotate(err, "Failed to determine root folder")
	}

	return rv, nil
}
