package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestResponse(t *testing.T) {
	t.Run("Compile", testResponseCompile)
}

func testResponseCompile(t *testing.T) {
	cases := []struct {
		in, want *Response
		Match
	}{
		{
			in:    &Response{200, "{{.code}} {{.status}}", "--{{.body}}--", map[string]string{"Foo": "{{.foo}}"}},
			want:  &Response{200, "200 OK", "--hello--", map[string]string{"Foo": "Bar"}},
			Match: Match{"code": "200", "status": "OK", "body": "hello", "foo": "Bar"},
		},
	}

	for _, c := range cases {
		r, err := c.in.Compile(c.Match)
		if err != nil {
			t.Fatal(err)
		}
		if err := assertEqual(r, c.want); err != nil {
			t.Error(err)
		}
	}
}

func assertEqual(got *http.Response, want *Response) error {
	if got.StatusCode != want.StatusCode {
		return fmt.Errorf("StatusCode %d != %d", got.StatusCode, want.StatusCode)
	}
	if got.Status != want.Status {
		return fmt.Errorf("Status %q != %q", got.Status, want.Status)
	}
	body, err := ioutil.ReadAll(got.Body)
	if err != nil {
		return err
	}
	if string(body) != want.Body {
		return fmt.Errorf("Body %q != %q", string(body), want.Body)
	}
	for name := range got.Header {
		for _, value := range got.Header[name] {
			if value != want.Header[name] {
				return fmt.Errorf("Header[%q]: %q != %q", name, value, want.Header[name])
			}
		}
	}
	return nil
}
