package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// A Request contains all the input strings against which to match a
// http.Request. It is generally compiled to a RequestMatcher which does the
// actual matching.
type Request struct {
	Method, URL, Body string
	Header            map[string]string
}

// Compile builds a RequestMatcher for the given Request.
func (r *Request) Compile() (*RequestMatcher, error) {
	var method, url, body, k, v *regexp.Regexp
	var err error

	compile := func(s string) (*regexp.Regexp, error) {
		if s == "" {
			s = ".*"
		}
		return regexp.Compile(s)
	}

	if method, err = compile(r.Method); err == nil {
		if url, err = compile(r.URL); err == nil {
			if body, err = compile(r.Body); err == nil {
				header := map[*regexp.Regexp]*regexp.Regexp{}
				for name := range r.Header {
					if k, err = compile(name); err != nil {
						return nil, err
					}
					if v, err = compile(r.Header[name]); err != nil {
						return nil, err
					}
					header[k] = v
				}
				return &RequestMatcher{method, url, body, header}, nil
			}
		}
	}
	return nil, err
}

// ParseRequest Makes a Request from a http.Request. This simplifies matching
// and printing the http.Request's values without worrying about reading the
// body or parsing the URL or headers. The http.Request's Body will be read at
// this point.
func ParseRequest(hr *http.Request) (*Request, error) {
	r := &Request{"", "", "", make(map[string]string)}
	r.Method = hr.Method
	if hr.URL != nil {
		r.URL = hr.URL.String()
	}
	if hr.Body != nil {
		body, err := ioutil.ReadAll(hr.Body)
		if err != nil {
			return nil, err
		}
		r.Body = string(body)
	}
	for name := range hr.Header {
		if len(hr.Header[name]) > 0 {
			r.Header[name] = hr.Header[name][0]
		}
	}
	return r, nil
}

// String implements the Stringer interface for a Request.
func (r *Request) String() string {
	trim := func(s string, max int, suffix string) string {
		count := 0
		for i := range s {
			count++
			if count > max {
				return s[:i] + suffix
			}
		}
		return s
	}
	return fmt.Sprintf(
		"%T@%[1]p{%s,%s,%s,%s}",
		r,
		r.Method,
		r.URL,
		trim(r.Body, 17, "..."),
		r.Header,
	)
}
