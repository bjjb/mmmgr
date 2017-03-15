package mock

import (
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func TestRequestMatcher(t *testing.T) {
	rm := &RequestMatcher{
		Method: regexp.MustCompile(`(?i:(?P<method>Post|Put))`),
		URL: regexp.MustCompile(
			`(?P<scheme>https?)://(?P<host>[^/]+)(?P<path>/[^?]+)?`),
		Body: regexp.MustCompile("--(?P<x>.+)--"),
		Header: map[*regexp.Regexp]*regexp.Regexp{
			regexp.MustCompile(`Accept`): regexp.MustCompile(`json`),
		},
	}
	body := strings.NewReader("--FOO--")
	r, err := http.NewRequest(http.MethodPost, "https://example.com/a/b", body)
	if err != nil {
		t.Error(err)
	}
	r.Header.Add("Accept", "application/json")
	m := rm.Match(r)
	if m["method"] != "POST" {
		t.Errorf("method: %q != %q", m["method"], "POST")
	}
	if m["scheme"] != "https" {
		t.Errorf("scheme: %q != %q", m["scheme"], "https")
	}
	if m["host"] != "example.com" {
		t.Errorf("host: %q != %q", m["host"], "example.com")
	}
	if m["path"] != "/a/b" {
		t.Errorf("path: %q != %q", m["path"], "/a/b")
	}
	if m["x"] != "FOO" {
		t.Errorf("x: %q != %q", m["x"], "FOO")
	}
}
