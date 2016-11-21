package file

import "testing"

func TestNew(t *testing.T) {
	expected := &File{"foo.mp4", "video/mp4", "video", nil}
	actual := New("foo.mp4")
	if actual.MimeType != expected.MimeType {
		t.Errorf("New(%q).MimeType is %q; expected %q", "foo.mp4", actual.MimeType, expected.MimeType)
	}
	if actual.MediaType != expected.MediaType {
		t.Errorf("New(%q).MediaType is %q; expected %q", "foo.mp4", actual.MediaType, expected.MediaType)
	}
	if actual.Path != expected.Path {
		t.Errorf("New(%q).Path is %q; expected %q", "foo.mp4", actual.Path, expected.Path)
	}
	if actual.Error != expected.Error {
		t.Errorf("New(%q).Error is %q; expected %q", "foo.mp4", actual.Error, expected.Error)
	}
}
