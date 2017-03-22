package tvdb

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/bjjb/mmmgr/net/http/mock"
)

func init() {
	httpClient = new(http.Client)
	file, err := os.Open("mocks.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()
	mocks, err := mock.ReadJSON(file)
	if err != nil {
		log.Fatal(err)
	}
	httpClient.Transport = mocks
}

func TestLanguages(t *testing.T) {
	c := new(Client)
	t.Run("Languages", testLanguages(c))
	t.Run("SearchSeriesByName", testSearchSeriesByName(c))
}

func testLanguages(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := DefaultClient.Languages(); err != nil {
			t.Error(err)
		}
	}
}

func testSearchSeriesByName(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := DefaultClient.SearchSeriesByName("Westworld"); err != nil {
			t.Error(err)
		}
	}
}
