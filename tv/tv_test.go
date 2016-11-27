package tv

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
	g.EpisodeTitle = "Justified"
	g.Title = "Blaj"
	g.Year = 2001
	g.Season = 2
	g.Episode = 3

	v := NewFromGuess(g)

	assert.Equal(t, v.EpisodeTitle, "Justified")
	assert.Equal(t, v.Title, "Blaj")
	assert.Equal(t, v.Year, 2001)
	assert.Equal(t, v.Season, 2)
	assert.Equal(t, v.Episode, 3)
}
