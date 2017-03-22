// Package cache provides a pluggable rfc7234 compliant cache for HTTP
// clients.
package cache

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// A Cache is a concurrency-safe HTTP client cache
type Cache struct {
	http.RoundTripper
	*log.Logger
	table         map[string]*response
	skipRequests  []skipRequest
	skipResponses []skipResponse
}

// A Store implements methods to store and retrieve a *http.Response by a key.
type Store interface {
	Fetch(*http.Request) *http.Response
	Put(*http.Request, *http.Response)
}

// Fetch returns the matching, non-expired HTTP response; returns nil, nil if
// it's not cached.
func (c *Cache) Fetch(req *http.Request) (*http.Response, error) {
	if c.skipRequest(req) {
		return nil, nil
	}
	if c.table == nil {
		c.table = make(map[string]*response)
		c.log("%s MISS\n", key(req))
		return nil, nil
	}
	key := key(req)
	if entry := c.table[key]; entry != nil {
		c.log("%s HIT\n", key)
		return deserialize(entry)
	}
	c.log("%s MISS\n", key)
	return nil, nil
}

// Store stores the HTTP response in the cache
func (c *Cache) Store(req *http.Request, resp *http.Response) error {
	if c.skipResponse(req, resp) {
		return nil
	}
	c.log("%s STORING\n", key(req))
	if c.table == nil {
		c.table = make(map[string]*response)
	}
	entry, err := serialize(resp)
	if err != nil {
		return err
	}
	c.table[key(req)] = entry
	return nil
}

// RoundTrip implements *http.RoundTripper for a Cache, so you can plug it
// into a client as a middleware.
func (c *Cache) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := c.Fetch(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return resp, nil
	}
	resp, err = c.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if err := c.Store(req, resp); err != nil {
		return nil, err
	}
	return resp, err
}

func key(req *http.Request) string {
	b := []byte(fmt.Sprintf("%s%s%s", req.Method, req.URL.String(), req.Header))
	return fmt.Sprintf("%x", md5.Sum(b))
}

type response struct {
	StatusCode int
	Status     string
	Header     http.Header
	Body       []byte
}

func serialize(in *http.Response) (*response, error) {
	body := []byte{}
	if _, err := io.Copy(bytes.NewBuffer(body), in.Body); err != nil {
		return nil, err
	}
	if err := in.Body.Close(); err != nil {
		return nil, err
	}
	return &response{
		StatusCode: in.StatusCode,
		Status:     in.Status,
		Header:     in.Header,
		Body:       body,
	}, nil
}

func deserialize(in *response) (*http.Response, error) {
	return &http.Response{
		StatusCode: in.StatusCode,
		Status:     in.Status,
		Header:     in.Header,
		Body:       ioutil.NopCloser(bytes.NewBuffer(in.Body)),
	}, nil
}

func (c *Cache) logger() *log.Logger {
	if c.Logger == nil {
		c.Logger = log.New(ioutil.Discard, "", 0)
	}
	return c.Logger
}

func (c *Cache) log(fmt string, rest ...interface{}) {
	logger := c.logger()
	if logger == nil {
		return
	}
	logger.Printf(fmt, rest)
}

type skipRequest func(*http.Request) bool
type skipResponse func(*http.Request, *http.Response) bool

func (c *Cache) skipRequest(req *http.Request) bool {
	for _, f := range c.skipRequests {
		if f(req) {
			return true
		}
	}
	return false
}

func (c *Cache) skipResponse(req *http.Request, resp *http.Response) bool {
	if c.skipRequest(req) {
		return true
	}
	for _, f := range c.skipResponses {
		if f(req, resp) {
			return true
		}
	}
	return false
}

func requestNotUnderstood(req *http.Request) bool { return false }

func foo() skipRequest {
	return requestNotUnderstood
}

var skipRequests = []skipRequest{foo()}
var skipResponses = []skipResponse{}
