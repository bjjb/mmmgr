package tvdb

import (
	"encoding/json"
	"fmt"
)

/*
A Series contains detailed information from The TVDB about a particular
series.
*/
type Series struct {
	ID              int      `json:"id"`
	SeriesName      string   `json:"seriesName"`
	Aliases         []string `json:"aliases"`
	Banner          string   `json:"banner"`
	SeriesID        string   `json:"seriesId"`
	Status          string   `json:"status"`
	FirstAired      string   `json:"firstAired"`
	Network         string   `json:"network"`
	NetworkID       string   `json:"networkId"`
	Runtime         string   `json:"runtime"`
	Genre           []string `json:"genre"`
	Overview        string   `json:"overview"`
	LastUpdated     int      `json:"lastUpdated"`
	AirsDayOfWeek   string   `json:"airsDayOfWeek"`
	AirsTime        string   `json:"airsTime"`
	Rating          string   `json:"rating"`
	IMDBID          string   `json:"imdbId"`
	Zap2ItID        string   `json:"zap2itId"`
	Added           string   `json:"added"`
	SiteRating      float64  `json:"siteRating"`
	SiteRatingCount int      `json:"siteRatingCount"`
	*Client
}

/*
GetSeries gets specific information about a particular series by the id.
*/
func (c *Client) GetSeries(id int) (*Series, error) {
	r, err := c.Get(fmt.Sprintf("series/%d", id))
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Data   *Series `json:"data"`
		Errors *struct {
			InvalidFilters     []string `json:"invalidFilters"`
			InvalidLanguage    string   `json:"invalidLanguage"`
			InvalidQueryParams []string `json:"invalidQueryParams"`
		} `json:"errors"`
	})
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	if result.Data == nil {
		return nil, fmt.Errorf("GetSeries errors: %q", result.Errors)
	}
	result.Data.Client = c
	return result.Data, nil
}
