package files

import "testing"

func TestMimeType(t *testing.T) {
	cases := []struct{ in, want string }{
		{"foo.mp3", "audio/mpeg"},
		{"foo.ogg", "audio/ogg"},
		{"foo.opus", "audio/opus"},
		{"foo.m4a", "audio/mp4a-latm"},
	}
	for _, c := range cases {
		got := MimeType(c.in)
		if c.want != got {
			t.Errorf("MimeType(%q) == %q; got %q", c.in, c.want, got)
		}
	}
}

func TestBasicType(t *testing.T) {
	cases := []struct{ in, want string }{
		{"audio/opus", "audio"},
		{"video/webm", "video"},
		{"application/pdf", "book"},
		{"text/plain", ""},
	}
	for _, c := range cases {
		got := BasicType(c.in)
		if c.want != got {
			t.Errorf("BasicType(%q) == %q; got %q", c.in, c.want, got)
		}
	}
}
