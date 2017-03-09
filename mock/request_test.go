package mock

import (
	"regexp"
	"testing"
)

func TestRequest(t *testing.T) {
	t.Run("Compile", testRequestCompile)
}

func testRequestCompile(t *testing.T) {
	cases := []struct {
		Method, URL, Body []string
		Header            []map[string]string
	}{
		{
			Method: []string{"Post", "Post"},
			URL:    []string{"http://example.com", "http://example.com"},
			Body:   []string{"(?P<foo>hello)", "(?P<foo>hello)"},
			Header: []map[string]string{{"a": "1", "b": "2"}, {"a": "1", "b": "2"}},
		},
	}
	for _, c := range cases {
		r := &Request{
			Method: c.Method[0],
			URL:    c.URL[0],
			Body:   c.Body[0],
			Header: c.Header[0],
		}
		rm, err := r.Compile()
		if err != nil {
			t.Error(err)
		}
		want := c.Method[1]
		got := rm.Method.String()
		if got != want {
			t.Errorf("rm.Method ~ %q; got %q", want, got)
		}
		want = c.URL[1]
		got = rm.URL.String()
		if got != want {
			t.Errorf("rm.URL ~ %q; got %q", want, got)
		}
		want = c.Body[1]
		got = rm.Body.String()
		if got != want {
			t.Errorf("rm.Body ~ %q; got %q", want, got)
		}
		if !headerEquals(rm.Header, c.Header[1]) {
			t.Errorf("Header %v doesn't match %q", rm.Header, c.Header[1])
		}
	}
}

func headerEquals(
	actual map[*regexp.Regexp]*regexp.Regexp,
	expected map[string]string,
) bool {
	for kr := range actual {
		s := kr.String()
		found := false
		for ks := range expected {
			if ks == s {
				found = true
				break
			}
		}
		if !found {
			return false
		}
		if actual[kr].String() != expected[s] {
			return false
		}
	}
	return true
}
