package mock

import (
	"net/url"
	"regexp"
	"strings"
	"testing"
)

func TestMatch(t *testing.T) {
	t.Run("doString", testMatchDoString)
	t.Run("doStrings", testMatchDoStrings)
	t.Run("doReader", testMatchDoReader)
	t.Run("doHeader", testMatchDoHeader)
	t.Run("doURL", testMatchDoURL)
}

func testMatchDoString(t *testing.T) {
	cases := []struct {
		re, in string
		want   map[string]string
	}{
		{`(?P<foo>.*)`, `hello`, map[string]string{"foo": "hello"}},
		{`(?P<x>\d)(?P<y>\w)`, `1a`, map[string]string{"x": "1", "y": "a"}},
	}
	for _, c := range cases {
		m := Match{}
		if !m.doString(regexp.MustCompile(c.re), c.in) {
			t.Errorf("doString(%q, %q) failed", c.re, c.in)
		}
		for k := range c.want {
			if c.want[k] != m[k] {
				t.Errorf("doString(%q, %q)[%q] == %q; got %q",
					c.re, c.in, k, c.want[k], m[k])
			}
		}
	}
}

func testMatchDoStrings(t *testing.T) {
	cases := []struct {
		re   string
		in   []string
		want map[string]string
	}{
		{`(?P<foo>.*)`, []string{"hi", "bye"}, map[string]string{"foo": "hi"}},
		{`(?P<x>\d)(?P<y>\w)`, []string{"1a", "2b"}, map[string]string{"x": "1", "y": "a"}},
	}

	for _, c := range cases {
		m := Match{}
		if !m.doStrings(regexp.MustCompile(c.re), c.in) {
			t.Errorf("doStrings(%q, %q) failed", c.re, c.in)
		}
		for k := range c.want {
			if c.want[k] != m[k] {
				t.Errorf("doStrings(%q, %q)[%q] == %q; got %q", c.re, c.in, k, c.want[k], m[k])
			}
		}
	}
}

func testMatchDoReader(t *testing.T) {
	cases := []struct {
		re, in string
		want   map[string]string
	}{
		{`(?P<foo>.*)`, `hello`, map[string]string{"foo": "hello"}},
		{`(?P<x>\d)(?P<y>\w)`, `1a`, map[string]string{"x": "1", "y": "a"}},
	}
	for _, c := range cases {
		m := Match{}
		if !m.doReader(regexp.MustCompile(c.re), strings.NewReader(c.in)) {
			t.Errorf("doString(%q, %q) failed", c.re, c.in)
		}
		for k := range c.want {
			if c.want[k] != m[k] {
				t.Errorf("doString(%q, %q)[%q] == %q; got %q",
					c.re, c.in, k, c.want[k], m[k])
			}
		}
	}
}

func testMatchDoHeader(t *testing.T) {
	cases := []struct {
		re   map[string]string
		in   map[string][]string
		want map[string]string
	}{
		{
			map[string]string{`(X|Y)`: `Foo (?P<foo>\w+)`},
			map[string][]string{"X": []string{"Foo Bar"}},
			map[string]string{"foo": "Bar"},
		},
		{
			map[string]string{`(X|Y)`: `Foo (?P<foo>\w+)`},
			map[string][]string{"Y": []string{"Foo Baz"}},
			map[string]string{"foo": "Baz"},
		},
	}
	for _, c := range cases {
		m := Match{}
		re := map[*regexp.Regexp]*regexp.Regexp{}
		for k := range c.re {
			re[regexp.MustCompile(k)] = regexp.MustCompile(c.re[k])
		}
		if !m.doHeader(re, c.in) {
			t.Errorf("doString(%q, %q) failed", c.re, c.in)
		}
		for k := range c.want {
			if c.want[k] != m[k] {
				t.Errorf("doString(%q, %q)[%q] == %q; got %q",
					c.re, c.in, k, c.want[k], m[k])
			}
		}
	}
}

func testMatchDoURL(t *testing.T) {
	cases := []struct {
		re, in string
		want   map[string]string
	}{
		{
			`(?P<scheme>https?)://(?P<host>[^/]+).*`,
			"http://foo.com/x",
			map[string]string{"scheme": "http", "host": "foo.com"},
		},
		{
			`(?P<scheme>https?)://(?P<host>[^/]+).*`,
			"https://bar.com/x",
			map[string]string{"scheme": "https", "host": "bar.com"},
		},
	}
	for _, c := range cases {
		m := Match{}
		url, err := url.Parse(c.in)
		if err != nil {
			t.Error(err)
		}
		if !m.doURL(regexp.MustCompile(c.re), url) {
			t.Errorf("doURL(%q, %q) failed", c.re, c.in)
		}
		for k := range c.want {
			if c.want[k] != m[k] {
				t.Errorf("doURL(%q, %q)[%q] == %q; got %q",
					c.re, c.in, k, c.want[k], m[k])
			}
		}
	}
}
