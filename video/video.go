package video

import (
	"github.com/bjjb/mmmgr/files"
	"github.com/bjjb/mmmgr/guess"
)

type Video struct {
	File      *files.File
	Codec     string
	Container string
	Duration  string
	Format    string
	MimeType  string
	Title     string
	Type      string
}

func GuessFromPath(path string) *Video {
	if g := guess.New(path); g != nil {
		return &Video{
			Codec:     g.VideoCodec,
			Container: g.Container,
			MimeType:  g.MimeType,
			Format:    g.Format,
			Title:     g.Title,
			Type:      g.Type,
		}
	}
	return nil
}

// Makes a new Video from a files.File.
func New(path string) *Video {
	r := new(Video)
	return r
}
