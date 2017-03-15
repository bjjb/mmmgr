package toys

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// A Counter is an incrementing RountTripper.
type Counter byte

// RoundTrip always gives a new response and a nil error
func (rt *Counter) RoundTrip(req *http.Request) (*http.Response, error) {
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
