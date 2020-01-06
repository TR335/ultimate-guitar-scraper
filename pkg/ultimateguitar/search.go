package ultimateguitar

import (
    "encoding/json"
    "strconv"
    "fmt"
    "net/http"
)

func (s *Scraper) Search(term string, page int) (SearchResult, Pagination, error){
	searchResult := SearchResult{}
    pageInfo := Pagination{}

	urlString := fmt.Sprintf("%s%s?title=%s&page=%d&type[]=300&type[]=800&official[]=0&official[]=1&", ugAPIEndpoint, AppPaths.SEARCH, term, page)
	req, _ := http.NewRequest("GET", urlString, nil)

	// Do some minor header manipulation so we retain the case
	for key := range ugHeaders {
		req.Header[key] = []string{ugHeaders[key]}
	}
	req.Header["X-UG-CLIENT-ID"] = []string{s.DeviceID}
	req.Header["X-UG-API-KEY"] = []string{s.generateAPIKey()}

	// This header isn't sent in the app, so we remove it.
	req.Header.Del("Accept-Encoding")

	res, err := s.Client.Do(req)
	if err != nil {
		return searchResult, pageInfo, err
	}
	defer res.Body.Close()
    //get the page info from the response headers
    pageInfo.TotalTabs, _ = strconv.Atoi(res.Header["X-Pagination-Total-Count"][0])
    pageInfo.TotalPages, _ = strconv.Atoi(res.Header["X-Pagination-Page-Count"][0])
    pageInfo.ThisPage, _ = strconv.Atoi(res.Header["X-Pagination-Current-Page"][0])

	err = json.NewDecoder(res.Body).Decode(&searchResult)
	if err != nil {
		return searchResult, pageInfo, err
	}

	return searchResult, pageInfo, nil
}

// SearchResult struct - this is what we get when we search for a term.
type SearchResult struct {
	Tabs []struct {
		ID                 int     `json:"id"`
		SongID             int     `json:"song_id"`
		SongName           string  `json:"song_name"`
		ArtistName         string  `json:"artist_name"`
		Type               string  `json:"type"`
		Part               string  `json:"part"`
		Version            int     `json:"version"`
		Votes              int     `json:"votes"`
		Rating             float64 `json:"rating"`
		Date               string  `json:"date"`
		Status             string  `json:"status"`
		PresetID           int     `json:"preset_id"`
		TabAccessType      string  `json:"tab_access_type"`
		TpVersion          int     `json:"tp_version"`
		TonalityName       string  `json:"tonality_name"`
		VersionDescription string  `json:"version_description"`
		Verified           int     `json:"verified"`
		Recording          struct {
			IsAcoustic       int           `json:"is_acoustic"`
			TonalityName     string        `json:"tonality_name"`
			Performance      interface{}   `json:"performance"`
			RecordingArtists []interface{} `json:"recording_artists"`
		} `json:"recording"`
	} `json:"tabs"`
	Artists []string `json:"artists"`
}

type Pagination struct {
    TotalTabs  int
    TotalPages int
    ThisPage   int
}
