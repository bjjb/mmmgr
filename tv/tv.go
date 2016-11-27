package tv

import (
	"github.com/bjjb/mmmgr/guess"
	"github.com/bjjb/mmmgr/video"
)

// Holds data related to a TV file (generally one episode)
type TV struct {
	*video.Video
	Season       int    `json:"season"`
	Episode      int    `json:"episode"`
	EpisodeTitle string `json:"episode_title"`
	Year         int    `json:"year"`
	Date         string `json:"date"`
	TVDBID       string `json:"tvdb_id"`
	TMDBID       string `json:"tmdb_id"`
}

func New(path string) *TV {
	return &TV{
		Video: video.New(path),
	}
}

// Makes a TV from a Guess
func NewFromGuess(g *guess.Guess) *TV {
	return &TV{
		Video:        video.NewFromGuess(g),
		Episode:      g.Episode,
		Season:       g.Season,
		EpisodeTitle: g.EpisodeTitle,
		Year:         g.Year,
		Date:         g.Date,
	}
}
