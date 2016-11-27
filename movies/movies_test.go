package movies

import (
	"github.com/bjjb/mmmgr/assert"
	"github.com/bjjb/mmmgr/guess"
	"github.com/bjjb/mmmgr/video"
	"testing"
)

func TestNew(t *testing.T) {
	m := New("foo")
	assert.Equal(t, m.Video, video.New("foo"))
}

func TestNewFromGuess(t *testing.T) {
	g := guess.New("foo.avi")
	g.Format = "DVD"
	g.VideoCodec = "XviD"
	g.Container = "avi"
	g.Title = "Home Movie"
	g.Year = 1998
	languages := []string{"en", "fr", "it"}
	g.Languages = languages[:]

	v := NewFromGuess(g)

	assert.Equal(t, v.Format, "DVD")
	assert.Equal(t, v.Codec, "XviD")
	assert.Equal(t, v.Container, "avi")
	assert.Equal(t, v.Title, "Home Movie")
	assert.Equal(t, v.Year, 1998)
	assert.Equal(t, v.Languages, languages)
}
