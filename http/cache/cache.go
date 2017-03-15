// Package cache helps to cut down on network requests.
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
	table map[string]*response
}

// Fetch returns the matching, non-expired HTTP response
func (c *Cache) Fetch(req *http.Request) (*http.Response, error) {
	if c.table == nil {
		c.table = make(map[string]*response)
		log.Printf("%s MISS\n", key(req))
		return nil, nil
	}
	if entry := c.table[key(req)]; entry != nil {
		log.Printf("%s HIT\n", key(req))
		return deserialize(entry)
	}
	log.Printf("%s MISS\n", key(req))
	return nil, nil
}

// Store stores the HTTP response in the cache
func (c *Cache) Store(req *http.Request, resp *http.Response) error {
	log.Printf("%s STORING\n", key(req))
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
