package tvdb

import (
	"encoding/json"
	"fmt"
	"net/url"
)

/*
A SeriesSearchResult is returned (in slices) by the SearchSeries* methods.
*/
type SeriesSearchResult struct {
	Aliases    []string `json:"aliases"`
	Banner     string   `json:"banner"`
	FirstAired string   `json:"firstAired"`
	ID         int      `json:"id"`
	Network    string   `json:"network"`
	Overview   string   `json:"overview"`
	Name       string   `json:"seriesName"`
	Status     string   `json:"status"`
}

/*
SearchSeriesByName queries The TVDB for series matching the given name.
*/
func (c *Client) SearchSeriesByName(name string) ([]SeriesSearchResult, error) {
	return c.SearchSeries(&url.Values{"name": {name}})
}

/*
SearchSeriesByImdbID queries The TVDB for series matching the given IMDB ID.
*/
func (c *Client) SearchSeriesByImdbID(id string) ([]SeriesSearchResult, error) {
	return c.SearchSeries(&url.Values{"imdbId": {id}})
}

/*
SearchSeriesByZap2itID queries The TVDB for series matching the given Zap2It
ID.
*/
func (c *Client) SearchSeriesByZap2itID(id string) ([]SeriesSearchResult, error) {
	return c.SearchSeries(&url.Values{"zap2itId": {id}})
}

/*
SearchSeries queries The TVDB for series matching the given url.Values.
*/
func (c *Client) SearchSeries(values *url.Values) ([]SeriesSearchResult, error) {
	r, err := c.Get("search/series?" + values.Encode())
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Data []SeriesSearchResult `json:"data"`
	})
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

/*
SearchSeriesParams gets a list of the params applicable for SearchSeries.
*/
func (c *Client) SearchSeriesParams() ([]string, error) {
	r, err := c.Get("search/series/params")
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Data struct {
			Params []string `json:"params"`
		} `json:"data"`
	})
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data.Params, nil
}

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
