package files

import (
	"encoding/json"
	"github.com/bjjb/mmmgr/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestJSONUnmarshal(t *testing.T) {
	f := new(File)
	j := []byte(`{"path": "a", "mime_type": "b", "type": "c"}`)
	if err := json.Unmarshal(j, f); err != nil {
		t.Errorf("Couldn't unmarshal JSON: %s", err)
	}
	if f.Path != "a" {
		t.Errorf(".Path: expected %q, got %q", "a", f.Path)
	}
	if f.MimeType != "b" {
		t.Errorf(".MimeType: expected %q, got %q", "b", f.MimeType)
	}
	if f.MediaType != "c" {
		t.Errorf(".MediaType: expected %q, got %q", "c", f.MediaType)
	}
}

func TestScan(t *testing.T) {
	dir, err := ioutil.TempDir("", "mmmgr-files-scan")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)
	testCases := []struct{ files []string }{
		{
			[]string{"foo.mp3", "foo.mp4", "foo.xls", "foo.pdf"},
		},
	}
	for _, test := range testCases {
		assert.WriteTempFiles(t, dir, test.files)
		files := Scan(dir)
		for f := range files {
			x := New(f.Path)
			assert.Equal(t, x.Path, f.Path)
			assert.Assert(t, "MediaType missing!", f.MediaType != "")
		}
	}
}
