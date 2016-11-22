package video

import "github.com/bjjb/mmmgr/guessit"

type Video struct {
	Codec     string
	Container string
	Duration  string
	Format    string
	MimeType  string
	Title     string
	Type      string
}

func GuessFromPath(path string) *Video {
	if g := guessit.Guessit(path); g != nil {
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
