package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bjjb/mmmgr/config"
	"io"
	"io/ioutil"
	"log"
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
	defer func() { _ = resp.Body.Close() }()
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
Get authorizes the client (if needed), GETS from the URL, and returns the
body, or an error.
*/
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
		return nil, fmt.Errorf("get %s not 200: [%s] (%s)", url, string(body))
	}
	return bytes.NewBuffer(body), nil
}

/*
Languages gets a list of the *Languages supported by TheTVDB.
*/
func (c *Client) Languages() ([]Language, error) {
	r, err := c.Get("languages")
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Data []Language `json:"data"`
	})
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func init() {
	auth := config.TVDB
	DefaultClient = NewClient(auth["apikey"], auth["username"], auth["userkey"])
}
