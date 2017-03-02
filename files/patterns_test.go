package files

import (
	"regexp"
	"testing"
)

func TestMatchFirst(t *testing.T) {
	cases := []struct {
		re            string
		in            []string
		key, expected string
	}{
		{"x", []string{"x"}, "", "x"},
		{"x", []string{"x", "y"}, "", "x"},
		{"x", []string{"y", "x"}, "", "x"},
	}
	for _, c := range cases {
		re := regexp.MustCompile(c.re)
		matches := MatchFirst(re, c.in)
		if matches == nil {
			t.Errorf("Regexp(%q).Match(%q) failed", c.re, c.in)
		}
		actual, found := matches[c.key]
		if !found {
			t.Errorf("Regexp(%q).Match(%q)[%q] not found", c.re, c.in, c.key)
		}
		if actual != c.expected {
			t.Errorf("Regexp(%q).Match(%q)[%q] == %q; got %q", c.re, c.in, c.key,
				c.expected, actual)
		}
	}
}

func TestMatch(t *testing.T) {
	cases := []struct {
		re, in, key, expected string
	}{
		{"x", "x", "", "x"},
		{"x", "-x-", "", "x"},
		{"-(?P<a>x+)-", "-x-", "a", "x"},
		{"-(?P<a>x+)-", "-xxx-", "a", "xxx"},
		{"-(?P<a>.+)-", "-xxx-", "a", "xxx"},
		{`(?P<t>.+) \((?P<y>\d{4})\)`, "Foo Bar (1999)", "t", "Foo Bar"},
		{`(?P<t>.+) \((?P<y>\d{4})\)`, "Foo Bar (1999)", "y", "1999"},
		{`(?P<t>[^/]+) \((?P<y>\d{4})\)`, "/a/b/Foo Bar (1999)", "t", "Foo Bar"},
		{`S(?P<s>\d+)E(?P<e>\d+)`, "/a/b/S.1999/S.S02E11.a.b.x", "s", "02"},
		{`S(?P<s>\d+)E(?P<e>\d+)`, "/a/b/S.1999/S.S02E11.a.b.x", "e", "11"},
		{
			`(?P<n>[^/]+)[ .]S(?P<s>\d+)E(?P<e>\d+)[ .](?P<t>.+)`,
			"/a/b/The.Show.1999/The.Show.S02E11.Foo.Bar", "", "The.Show.S02E11.Foo.Bar",
		},
		{
			`(?P<n>[^/]+)[ .]S(?P<s>\d+)E(?P<e>\d+)[ .](?P<t>.+)`,
			"/a/b/The.Show.1999/The.Show.S02E11.Foo.Bar", "n", "The.Show",
		},
		{
			`(?P<n>[^/]+)[ .]S(?P<s>\d+)E(?P<e>\d+)[ .](?P<t>.+)`,
			"/a/b/The.Show.1999/The.Show.S02E11.Foo.Bar", "s", "02",
		},
		{
			`(?P<n>[^/]+)[ .]S(?P<s>\d+)E(?P<e>\d+)[ .](?P<t>.+)`,
			"/a/b/The.Show.1999/The.Show.S02E11.Foo.Bar", "e", "11",
		},
		{
			`(?P<n>[^/]+)[ .]S(?P<s>\d+)E(?P<e>\d+)[ .](?P<t>.+)`,
			"/a/b/The.Show.1999/The.Show.S02E11.Foo.Bar", "t", "Foo.Bar",
		},
	}
	for _, c := range cases {
		re := regexp.MustCompile(c.re)
		matches := Match(re, c.in)
		if matches == nil {
			t.Errorf("Regexp(%q).Match(%q) failed", c.re, c.in)
		}
		actual, found := matches[c.key]
		if !found {
			t.Errorf("Regexp(%q).Match(%q)[%q] not found", c.re, c.in, c.key)
		}
		if actual != c.expected {
			t.Errorf("Regexp(%q).Match(%q)[%q] == %q; got %q", c.re, c.in, c.key,
				c.expected, actual)
		}
	}
}
