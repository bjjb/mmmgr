package tv

import (
	"github.com/bjjb/mmmgr/files"
	"github.com/bjjb/mmmgr/guessit"
)

// Holds data related to a TV file (generally one episode)
type TV struct {
	Show    string `json:"show"`
	Season  int    `json:"season"`
	Episode int    `json:"episode"`
	Title   string `json:"title"`
	Year    int    `json:"year"`
	TVDBID  string `json:"tvdb_id"`
	TMDBID  string `json:"tmdb_id"`
	*files.File
}

// Guesses the name of a TV series from a path - will return nil if the guess
// failed, or if the type wasn't guessed as "tv".
func GuessFromPath(path string) *TV {
	f := files.New(path)
	if g := guessit.Guessit(f.Path); g != nil && g.Type == "episode" {
		return &TV{
			File:    f,
			Show:    g.Title,
			Season:  g.Season,
			Episode: g.Episode,
			Title:   g.EpisodeTitle,
			Year:    g.Year,
		}
	}
	return nil
}

// Makes a new TV from a files.File.
func New(path string) *TV {
	tv := GuessFromPath(path)
	return tv
}
