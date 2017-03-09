package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bjjb/mmmgr/config"
	"net/http"
	"net/url"
	"time"
)

/*
Endpoint is the location of the TVDB API
*/
var Endpoint = "https://api.thetvdb.com"

/*
Version is the content-type of the API, which includes the version number.
*/
var Version = "application/vnd.thetvdb.v2"

/*
TokenHoursToLive is the number of hours for which a JWT token is expected to
remain valid.
*/
var TokenHoursToLive = 24.0

/*
DefaultClient is a convenient Client, configured (in init()) using user's
configuration values, which is used in the Command and in tests.
*/
var DefaultClient *Client

/*
httpClient is the underlying HTTP Client for sending requests to The TVDB.
Replace this if you want to mock requests.
*/
var httpClient = http.DefaultClient

/*
A Client exposes methods talk to TheTVDB JSON API securely.
*/
type Client struct {
	apikey, username, userkey, token string
	loggedInAt                       time.Time
}

/*
NewClient gets a new *Client which can log in to TheTVDB as needed.
*/
func NewClient(apikey, username, userkey string) *Client {
	c := new(Client)
	c.apikey = apikey
	c.username = username
	c.userkey = userkey
	return c
}

/*
login logs the *Client c into TheTVDB using its credentials which were
supplied to NewClient. You probably won't need to call this yourself.
*/
func (c *Client) login() error {
	body := map[string]string{
		"apikey":   c.apikey,
		"username": c.username,
		"userkey":  c.userkey,
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

/*
refreshToken updates the token of *Client c. It must already have a token, and
its token must be valid.
*/
func (c *Client) refreshToken() error {
	values := url.Values{"token": {c.token}}
	url := Endpoint + "/refresh_token?" + values.Encode()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", Version)
	req.Header.Add("Authorization", "Bearer "+c.token)
	return c.doTokenRequest(req)
}

/*
doTokenRequest is used by login and refreshToken to send the preconfigured
request, parse the response, and either set the *Client's token and
loggedInAt, or return the appropriate error.
*/
func (c *Client) doTokenRequest(req *http.Request) error {
	now := time.Now()
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed: %s", resp.Status)
	}
	result := make(map[string]string)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	c.token = result["token"]
	c.loggedInAt = now
	return nil
}

/*
authorize checks the *Client c for a token and does nothing if it is found and
valid; otherwise it refreshes a token which is about to expire, or tries to
log in if the token is expired or blank.
*/
func (c *Client) authorize() error {
	if c.token == "" {
		return c.login()
	}
	if time.Now().Sub(c.loggedInAt).Hours() >= TokenHoursToLive {
		return c.refreshToken()
	}
	return nil
}

/*
A Language encapsulates a language supported by The TVDB.
*/
type Language struct {
	ID          int    `json:"id"`
	Abbr        string `json:"abbreviation"`
	Name        string `json:"name"`
	EnglishName string `json:"englishName"`
}

/*
Languages gets a list of the *Languages supported by TheTVDB.
*/
func (c *Client) Languages() ([]Language, error) {
	if err := c.authorize(); err != nil {
		return nil, err
	}
	url := Endpoint + "/languages"
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
	defer resp.Body.Close()
	result := new(struct {
		Data []Language `json:"data"`
	})
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

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
	return c.SearchSeries(url.Values{"name": {name}})
}

/*
SearchSeriesByImdbID queries The TVDB for series matching the given IMDB ID.
*/
func (c *Client) SearchSeriesByImdbID(id string) ([]SeriesSearchResult, error) {
	return c.SearchSeries(url.Values{"imdbId": {id}})
}

/*
SearchSeriesByZap2itID queries The TVDB for series matching the given Zap2It
ID.
*/
func (c *Client) SearchSeriesByZap2itID(id string) ([]SeriesSearchResult, error) {
	return c.SearchSeries(url.Values{"zap2itId": {id}})
}

/*
SearchSeries queries The TVDB for series matching the given url.Values.
*/
func (c *Client) SearchSeries(values url.Values) ([]SeriesSearchResult, error) {
	if err := c.authorize(); err != nil {
		return nil, err
	}
	url := Endpoint + "/search/series?" + values.Encode()
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
	defer resp.Body.Close()
	result := new(struct {
		Data []SeriesSearchResult `json:"data"`
	})
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func init() {
	auth := config.TVDB
	DefaultClient = NewClient(auth["apikey"], auth["username"], auth["userkey"])
}
