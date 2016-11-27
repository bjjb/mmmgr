package movies

import (
	"github.com/bjjb/mmmgr/guess"
	"github.com/bjjb/mmmgr/video"
)

// Holds data related to a Movie file
type Movie struct {
	*video.Video
	Year   int    `json:"year"`
	TMDBID string `json:"tmdb_id"`
	IMDBID string `json:"imdb_id"`
}

func New(path string) *Movie {
	m := &Movie{
		Video: video.New(path),
	}
	return m
}

func NewFromGuess(g *guess.Guess) *Movie {
	return &Movie{
		Video: video.NewFromGuess(g),
		Year:  g.Year,
	}
}
