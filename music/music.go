package music

import "github.com/bjjb/mmmgr/audio"

// Holds data related to a Music file
type Music struct {
	Artist string
	Album  string
	Title  string
	Disc   string
	Number string
	MBID   string
}

func GuessFromPath(path string) *Music {
	a := audio.GuessFromPath(path)
	return &Music{
		Title: a.Title,
	}
}
