package tv

import "github.com/bjjb/mmmgr/guessit"

// Holds data related to a TV file (generally one episode)
type TV struct {
	Show    string
	Season  int
	Episode int
	Title   string
	Year    int
	TVDBID  string
	TMDBID  string
}

// Guesses the name of a TV series from a path - will return nil if the guess
// failed, or if the type wasn't guessed as "tv".
func GuessFromPath(path string) *TV {
	if g := guessit.Guessit(path); g != nil && g.Type == "episode" {
		return &TV{
			Show:    g.Title,
			Season:  g.Season,
			Episode: g.Episode,
			Title:   g.EpisodeTitle,
			Year:    g.Year,
		}
	}
	return nil
}
