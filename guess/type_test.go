package guess

import (
	"github.com/bjjb/mmmgr/assert"
	"testing"
)

func TestMimeType(t *testing.T) {
	testCases := []struct{ t, x string }{
		{"foo.pdf", "application/pdf"},
		{"foo.mobi.pdf", "application/pdf"},
		{"foo", ""},
		{"", ""},
		{"     s     .ass   .sas", ""},
		{`\/`, ""},
		{"foo.mp3", "audio/mpeg"},
		{"foo.mp4", "video/mp4"},
	}
	for _, c := range testCases {
		x := MimeType(c.t)
		assert.Equal(t, c.x, x)
	}
}

func TestType(t *testing.T) {
	testCases := []struct{ t, x string }{
		{"foo.mp3", "audio"},
		{"foo.mp4", "video"},
		{"foo.pdf", "book"},
		{"foo.1999.pdf", "book"},
		{"foo.1999.mp4", "movie"},
		{"Foo (1999).mp4", "movie"},
		{"foo.s01e02.mp4", "tv"},
		{"foo.s01e02.bar.baz.mp4", "tv"},
		{"foo/s01e02-bar baz.mp4", "tv"},
		{"01x02-bar baz.mp4", "tv"},
	}
	for _, c := range testCases {
		x := Type(c.t)
		assert.Equal(t, c.x, x)
	}
}
