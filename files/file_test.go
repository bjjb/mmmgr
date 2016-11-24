package files

import (
  "testing"
  "encoding/json"
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
