package movies

import (
	"github.com/bjjb/mmmgr/guessit"
	"github.com/bjjb/mmmgr/files"
)

// Holds data related to a Movie file
type Movie struct {
	Title  string				`json:"title"`
	Year   int					`json:"year"`
	TMDBID string				`json:"tmdb_id"`
	IMDBID string				`json:"imdb_id"`
	*files.File
}

// Guesses the name of a TV series from a path - will return nil if the guess
// failed, or if the type wasn't guessed as "tv".
func GuessFromPath(path string) *Movie {
	f := files.New(path)
	if g := guessit.Guessit(f.Path); g != nil && g.Type == "movie" {
		return &Movie{
			File:	 f,
			Title: g.Title,
			Year:  g.Year,
		}
	}
	return nil
}

// Makes a new Movie from a path
func New(path string) *Movie {
	movie := GuessFromPath(path)
	return movie
}
