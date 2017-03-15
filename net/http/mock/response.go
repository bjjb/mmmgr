package mock

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

/*
A Response is a set of templates used to respond to matching Requests.
*/
type Response struct {
	StatusCode   int
	Status, Body string
	Header       map[string]string
}

/*
Compile expects a Match, which it uses to execute the Templates in the
Response to populate a *http.Response.
*/
func (r *Response) Compile(m Match) (*http.Response, error) {
	hr := &http.Response{}
	s, err := renderInt(r.StatusCode, m)
	if err != nil {
		return nil, err
	}
	code, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	hr.StatusCode = code
	status, err := renderString(r.Status, m)
	if err != nil {
		return nil, err
	}
	hr.Status = status
	body, err := renderString(r.Body, m)
	if err != nil {
		return nil, err
	}
	hr.Body = ioutil.NopCloser(strings.NewReader(body))
	header := http.Header{}
	for _, h := range r.Header {
		name, err := renderString(h, m)
		if err != nil {
			return nil, err
		}
		value, err := renderString(r.Header[name], m)
		if err != nil {
			return nil, err
		}
		header.Add(name, value)
	}
	hr.Header = header
	return hr, nil
}

/*
String implements the Stringer interface for a Response.
*/
func (r *Response) String() string {
	return fmt.Sprintf(
		"%T@%[1]p{%d,%s,%q,%v}",
		r,
		r.StatusCode,
		r.Status,
		r.Body,
		r.Header,
	)
}

/*
renderString executes the Template with the given Match, and returns the
resulting string, or an error.
*/
func renderString(t string, m Match) (string, error) {
	templ, err := template.New(fmt.Sprintf("%x", time.Now().UnixNano())).Parse(t)
	if err != nil {
		return "", err
	}
	b := new(bytes.Buffer)
	err = templ.Execute(b, m)
	if err != nil {
		return "", err
	}
	return string(b.Bytes()), nil
}

/*
renderInt converts i to a String and calls renderString.
*/
func renderInt(i int, m Match) (string, error) {
	return renderString(strconv.Itoa(i), m)
}
