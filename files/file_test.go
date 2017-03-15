package files

import (
	"testing"
)

func TestNew(t *testing.T) {
	type info map[string]string
	type result struct {
		mediaType string
		info      info
	}
	type test struct {
		in  string
		out *result
	}

	cases := []test{
		{"foo.mov", &result{"", info{}}},
		{"foo.1999.mov", &result{"movie", info{"title": "foo", "year": "1999"}}},
		{"ass.1999/foo.1999.mov", &result{"movie", info{"title": "foo", "year": "1999"}}},
		{"ass.1999/foo.1999.HDtv.mov", &result{"movie", info{"title": "foo", "year": "1999"}}},
		{"foo.S02E12.blah.webm", &result{"tv", info{"show": "foo", "season": "02", "episode": "12", "title": "blah"}}},
		{"foo.1999.S02E12.blah.webm", &result{"tv", info{"show": "foo", "season": "02", "episode": "12", "title": "blah", "year": "1999"}}},
		{"Downloads/Cool.Show.S04E03.WEBRip.x264-FUM[ettv]/Cool.Show.S04E03.WEBRip.x264-FUM[ettv].mp4",
			&result{"tv", info{"show": "Cool.Show", "title": "", "source": "WEBRip", "vcodec": "x264", "rgroup": "FUM[ettv]"}}},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			expected := c.out
			f := New(c.in)

			t.Run("MediaType", func(t *testing.T) {
				got := f.MediaType()
				if got != expected.mediaType {
					t.Errorf("wanted %q; got %q", expected.mediaType, got)
				}
			})

			t.Run("MediaInfo", func(t *testing.T) {
				for k := range c.out.info {
					t.Run(k, func(t *testing.T) {
						got := f.MediaInfo()[k]
						if got != expected.info[k] {
							t.Errorf("wanted %q; got %q", expected.info[k], got)
						}
					})
				}
			})
		})
	}
}
