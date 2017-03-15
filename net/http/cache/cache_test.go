package cache

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bjjb/mmmgr/net/http/toys"
)

func TestCache(t *testing.T) {
	url := "http://example.com"
	client := &http.Client{
		Transport: &Cache{RoundTripper: new(toys.Counter)},
	}

	get := func(url string) *http.Response {
		r, err := client.Get(url)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}
	bodyText := func(resp *http.Response) string {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		return string(b)
	}

	b1, b2 := bodyText(get(url)), bodyText(get(url))
	if b1 != b2 {
		t.Errorf("%q != %q", b1, b2)
	}
}
