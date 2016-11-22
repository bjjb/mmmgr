package movies

import "github.com/bjjb/mmmgr/guessit"

// Holds data related to a Movie file
type Movie struct {
	Title  string
	Year   int
	TMDBID string
	IMDBID string
}

// Guesses the name of a TV series from a path - will return nil if the guess
// failed, or if the type wasn't guessed as "tv".
func GuessFromPath(path string) *Movie {
	if g := guessit.Guessit(path); g != nil && g.Type == "movie" {
		return &Movie{
			Title: g.Title,
			Year:  g.Year,
		}
	}
	return nil
}
