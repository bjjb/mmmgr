package cache

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

type incrt byte // incrementing RountTripper

// RoundTrip always gives a new response and a nil error
func (rt *incrt) RoundTrip(req *http.Request) (*http.Response, error) {
	v := byte(*rt)
	resp := &http.Response{
		StatusCode: 200,
		Status:     fmt.Sprintf("200 OK (%d)", v),
		Header:     http.Header{"X-Value": []string{strconv.Itoa(int(v))}},
		Body:       ioutil.NopCloser(bytes.NewReader([]byte{v + '0'})),
	}
	(*rt)++
	return resp, nil
}

func TestCache(t *testing.T) {
	url := "http://example.com"
	client := &http.Client{Transport: &Cache{RoundTripper: new(incrt)}}

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
