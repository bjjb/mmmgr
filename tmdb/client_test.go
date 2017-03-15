package tmdb

import (
	"fmt"
	"testing"

	"github.com/bjjb/mmmgr/cfg"
)

func init() {
}

type TMDB struct {
	*Client
}

func (t *TMDB) PostToken() {
}

type Client struct {
	APIKey string `json:"apikey"`
}

func TestClient(t *testing.T) {
	client := new(Client)
	if err := cfg.UnmarshalKey("tmdb", &client); err != nil {
		t.Fatal(err)
	}
	t.Log(client)
}

func ExampleTMDB_PostToken() {
	fmt.Println("ok")
	// Output:
	// ok
}