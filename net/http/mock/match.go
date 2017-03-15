package mock

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

/*
A Match is a map of replacement strings for Response templates, which have a
number of methods that allow them to be built up from Regexps.
*/
type Match map[string]string

/*
doString adds to m submatches found when matching s to r; returns false if s
doesn't match r.
*/
func (m Match) doString(r *regexp.Regexp, s string) bool {
	if matches := r.FindStringSubmatch(s); matches != nil {
		for i, k := range r.SubexpNames()[1:] {
			m[k] = matches[i+1]
		}
		return true
	}
	return false
}

/*
doStrings applies doString to all of the strings, and returns true if any of
them match (strings after the first matching string are ignored).
*/
func (m Match) doStrings(r *regexp.Regexp, strings []string) bool {
	matched := false
	for _, s := range strings {
		if m.doString(r, s) {
			matched = true
			break
		}
	}
	return matched
}

/*
doReader applies doString to the contents of r.
*/
func (m Match) doReader(re *regexp.Regexp, r io.Reader) bool {
	if re == nil && r == nil {
		return true
	}
	s := ""
	if r != nil {
		data, err := ioutil.ReadAll(r)
		if err != nil {
			panic("error reading io")
		}
		s = string(data)
	}
	return m.doString(re, s)
}

/*
doHeader applies matches to the keys and values of header. Only the first
matching value (per key) is considered.
*/
func (m Match) doHeader(
	rexmap map[*regexp.Regexp]*regexp.Regexp,
	header http.Header,
) bool {
	for r := range rexmap {
		matched := false
		for h := range header {
			if m.doString(r, h) {
				if m.doStrings(rexmap[r], header[h]) {
					matched = true
					break
				}
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

/*
doURL applies doString to url.String()
*/
func (m Match) doURL(re *regexp.Regexp, url *url.URL) bool {
	s := ""
	if url != nil {
		s = url.String()
	}
	return m.doString(re, s)
}
