package mock

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestReadJSON(t *testing.T) {
	json := `[
	{
		"In":{"Method":"GET","URL":"http://example.com/"},
		"Out":{"Status":"200 OK","StatusCode":200,"Body":"index"}
	},
	{
		"In":{"Method":"POST","URL":"http://example.com/","Body":"(?P<body>)"},
		"Out":{"Status":"201 Created","StatusCode":201,"Body":"{{.body}}"}
	}
]`
	mocks, err := ReadJSON(strings.NewReader(json))
	if err != nil {
		t.Fatal(err)
	}

	if mocks == nil {
		t.Fatal("Read returned nil")
	}

	for _, mock := range *mocks {
		if mock.In == nil {
			t.Error("Request cannot be nil")
		}
		if mock.Out == nil {
			t.Error("Response cannot be nil")
		}
	}
}

/*
This example shows how to read in JSON to create a MockSet, to use it to match
requests, and to have the mocks return responses with tailored content.
*/
func Example_readingJSON() {
	// In this example, the Reader is from a string, but it could as easily be
	// from a file or a network stream.
	json := `[
	{
		"In":{"Method":"GET","URL":"http://example.com/"},
		"Out":{"Status":"200 OK","StatusCode":200,"Body":"index"}
	},
	{
		"In":{"Method":"POST","URL":"http://example.com/","Body":"(?P<x>h.*o)"},
		"Out":{"Status":"201 Created","StatusCode":201,"Body":"{{.x}}"}
	}
]`

	// Read in the JSON
	mocks, err := ReadJSON(strings.NewReader(json))
	if err != nil {
		log.Fatal(err)
	}

	// Make a http.Client which uses the MockSet as its RoundTripper
	client := &http.Client{Transport: mocks}

	// This one should not match.
	if _, err := client.Head("http://foo.com/"); err == nil {
		log.Fatal("Expected unmocked request to fail")
	} else {
		fmt.Println("HEAD => unmocked")
	}

	// This should be mocked by the first definition
	if resp, err := client.Get("http://example.com/"); err == nil {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			// the response body was defined in the first definition's Out template
			fmt.Printf("GET => %s\n", string(body))
			_ = resp.Body.Close()
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}

	// This should be mocked by the second definition
	if resp, err := client.Post(
		"http://example.com/",
		"text/plain",
		strings.NewReader("hello"),
	); err == nil {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			// the response body is actually extracted from the request
			fmt.Printf("POST => %s\n", string(body))
			_ = resp.Body.Close()
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}

	// Output:
	// HEAD => unmocked
	// GET => index
	// POST => hello
}
