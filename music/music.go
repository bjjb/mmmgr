package music

import (
	"github.com/bjjb/mmmgr/audio"
	"github.com/bjjb/mmmgr/files"
)

// Holds data related to a Music file
type Music struct {
	File	 *files.File
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

// Makes a new Music from a files.File.
func New(path string) *Music {
	r := new(Music)
	return r
}
