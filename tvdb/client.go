package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/bjjb/mmmgr/cfg"
)

const (
	// Endpoint is the location of the TVDB API
	Endpoint = "https://api.thetvdb.com"
	// Version is the content-type of the API, which includes the version number
	Version = "application/vnd.thetvdb.v2"
)

// DefaultClient is a convenient Client, configured (in init()) using user's
// configuration values, which is used in the Command and in tests.
var DefaultClient = new(Client)

// Debug is used for logging.
var Debug = func(fmt string, rest ...interface{}) {}

// httpClient is the underlying HTTP Client for sending requests to The TVDB.
// Replace this if you want to mock requests.
var httpClient = http.DefaultClient

// A Client exposes methods talk to TheTVDB JSON API securely.
type Client struct {
	APIKey           string `json:"apikey"`
	UserName         string `json:"username"`
	UserKey          string `json:"userkey"`
	token, tokenFile string
}

// login logs the *Client c into TheTVDB using its credentials which were
// supplied to NewClient. You probably won't need to call this yourself.
func (c *Client) login() error {
	body := map[string]string{
		"apikey":   c.APIKey,
		"username": c.UserName,
		"userkey":  c.UserKey,
	}
	buffer := new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, Endpoint+"/login", buffer)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", Version)
	req.Header.Add("Content-Type", "application/json")
	return c.doTokenRequest(req)
}

func (c *Client) refreshToken() error {
	values := &url.Values{"token": {c.token}}
	url := Endpoint + "/refresh_token?" + values.Encode()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", Version)
	req.Header.Add("Authorization", "Bearer "+c.token)
	return c.doTokenRequest(req)
}

func (c *Client) doTokenRequest(req *http.Request) error {
	Debug("doTokenRequest(%v)", req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"login failed: %d [%s]",
			resp.StatusCode,
			resp.Status,
		)
	}
	result := make(map[string]string)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	c.token = result["token"]
	return nil
}

func (c *Client) authorize() error {
	if c.token == "" {
		return c.login()
	}
	return nil
}

// Get authorizes the client (if needed), GETS from the URL, and returns the
// body, or an error.
func (c *Client) Get(url string) (io.Reader, error) {
	if err := c.authorize(); err != nil {
		return nil, err
	}
	url = fmt.Sprintf("%s/%s", Endpoint, url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", Version)
	req.Header.Add("Authorization", "Bearer "+c.token)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %d [%s]", url, resp.StatusCode, resp.Status)
	}
	return bytes.NewBuffer(body), nil
}

// Languages gets a list of the *Languages supported by TheTVDB.
func (c *Client) Languages() ([]*Language, error) {
	r, err := c.Get("languages")
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Data []*Language `json:"data"`
	})
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func init() {
	if err := cfg.UnmarshalKey("tvdb", DefaultClient); err != nil {
		log.Fatal(err)
	}
}

// A SeriesSearchResult is returned (in slices) by the SearchSeries* methods.
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

// SearchSeriesByName queries The TVDB for series matching the given name.
func (c *Client) SearchSeriesByName(name string) ([]SeriesSearchResult, error) {
	return c.SearchSeries(&url.Values{"name": {name}})
}

// SearchSeriesByImdbID queries The TVDB for series matching the given IMDB
// ID.
func (c *Client) SearchSeriesByImdbID(id string) ([]SeriesSearchResult, error) {
	return c.SearchSeries(&url.Values{"imdbId": {id}})
}

// SearchSeriesByZap2itID queries The TVDB for series matching the given
// Zap2It ID.
func (c *Client) SearchSeriesByZap2itID(id string) ([]SeriesSearchResult, error) {
	return c.SearchSeries(&url.Values{"zap2itId": {id}})
}

// SearchSeries queries The TVDB for series matching the given url.Values.
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

// SearchSeriesParams gets a list of the params applicable for SearchSeries.
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
