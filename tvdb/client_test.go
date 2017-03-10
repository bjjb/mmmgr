package tvdb

import (
	"github.com/bjjb/mmmgr/mock"
	"log"
	"net/http"
	"os"
	"testing"
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
	t.Run("Languages", testLanguages)
	t.Run("SearchSeriesByName", testSearchSeriesByName)
}

func testLanguages(t *testing.T) {
	if _, err := DefaultClient.Languages(); err != nil {
		t.Error(err)
	}
}

func testSearchSeriesByName(t *testing.T) {
	if _, err := DefaultClient.SearchSeriesByName("Westworld"); err != nil {
		t.Error(err)
	}
}
