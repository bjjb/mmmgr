package mock

import (
	"net/http"
	"regexp"
)

/*
A RequestMatcher is a set of regular expressions compiled from a Request.
*/
type RequestMatcher struct {
	Method, URL, Body *regexp.Regexp
	Header            map[*regexp.Regexp]*regexp.Regexp
}

/*
Match returns a Match if all of the components in the Request match their
corresponding expressions in the RequestMatcher.
*/
func (rm *RequestMatcher) Match(r *http.Request) Match {
	match := make(Match)
	if match.doString(rm.Method, r.Method) &&
		match.doURL(rm.URL, r.URL) &&
		match.doReader(rm.Body, r.Body) &&
		match.doHeader(rm.Header, r.Header) {
		return match
	}
	return nil
}
