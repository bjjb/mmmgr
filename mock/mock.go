/*
Package mock is a simple web mocking library for go. A mock.Set is a
http.RoundTripper that can be used in place of the default http.Client's
Transport.
*/
package mock

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
A Mock contains a Request to compare to *http.Requests and the http.Response
and error value to return if it matches.
*/
type Mock struct {
	// In is the set of regular expression strings against which to match
	// requests
	In *Request
	// Out is the set of template strings with which to generate the response
	Out *Response
	// Err is a template string for generating an error; "" implies no error.
	Err string
	// RequestMatcher is used to match against http.Requests
	*RequestMatcher
}

/*
String implements the Stringer interface for a Mock.
*/
func (m *Mock) String() string {
	return fmt.Sprintf(
		"%T@%[1]p{%s,%s,%s,%p}",
		m,
		m.In,
		m.Out,
		m.Err,
		m.RequestMatcher,
	)
}

/*
A Set is a set of Mocks which implements RoundTripper.
*/
type Set []*Mock

/*
RoundTrip implements the http.RoundTripper interface, which allows a
http.Client to use a MockSet instead of its default Transport.
*/
func (ms Set) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		defer req.Body.Close()
	}
	for _, mock := range ms {
		if mock.RequestMatcher == nil {
			rm, err := mock.In.Compile()
			if err != nil {
				return nil, err
			}
			mock.RequestMatcher = rm
		}
		if match := mock.RequestMatcher.Match(req); match != nil {
			if mock.Err != "" {
				return nil, fmt.Errorf(mock.Err)
			}
			resp, err := mock.Out.Compile(match)
			if err != nil {
				return nil, err
			}
			return resp, nil
		}
	}
	return nil, ErrUnmocked{req, ms}
}

/*
String implements the Stringer interface for a MockSet.
*/
func (ms Set) String() string {
	s := "["
	for _, mock := range ms {
		s += mock.String()
	}
	return s + "]"
}

/*
ErrUnmocked is an error created when a request is unmocked
*/
type ErrUnmocked struct {
	*http.Request
	Set
}

/*
Error implemets the errors.Error interface for an ErrUnmocked.
*/
func (e ErrUnmocked) Error() string {
	r, err := ParseRequest(e.Request)
	if err != nil {
		return fmt.Sprintf(
			"unmocked request %v <unparseable! %s>, (%d mocks)",
			e.Request,
			err,
			len(e.Set),
		)
	}
	return fmt.Sprintf("Unmocked: %s, (%d mocks)", r, len(e.Set))
}

/*
Read reads a Set (in JSON format) from the io.Reader r and returns it (or the
resultant error). The JSON should be an array of objects, whose keys
correspond to a Mock (where In is a Request, Out is a Response, and Err is an
error). The In strings will be compiled to regular expressions, and the Out
strings (and Err string) will be treated as Go templates. Names submatches in
the In strings may be used as variables in the Out and Err templates.
	Note:
Since the JSON is compiled into regular expressions, you should escape
special characters such as ? and . if you want them to match, like this:
	[
		{
			"In": { "URL": "http://www.example.com/path\\?foo=bar" },
			"Out": { "StatusCode": 200, "Status": "200 OK", "Body": "OK" }
		}
	]
*/
func ReadJSON(r io.Reader) (*Set, error) {
	result := Set{}
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
