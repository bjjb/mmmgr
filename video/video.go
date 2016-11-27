package video

import (
	"github.com/bjjb/mmmgr/files"
	"github.com/bjjb/mmmgr/guess"
)

type Video struct {
	*files.File
	AudioChannels     string   `json:"audio_channels"`
	Codec             string   `json:"codec"`
	Container         string   `json:"container"`
	Duration          string   `json:"duration"`
	Format            string   `json:"format"`
	Languages         []string `json:"languages"`
	ScreenSize        string   `json:"screen_size"`
	SubtitleLanguages []string `json:"subtitle_languages"`
	Title             string   `json:"title"`
}

func New(path string) *Video {
	v := &Video{
		File:              files.New(path),
		Languages:         make([]string, 5),
		SubtitleLanguages: make([]string, 10),
	}
	return v
}

func NewFromGuess(g *guess.Guess) *Video {
	v := New(g.Path)
	v.MimeType = g.MimeType
	v.Codec = g.VideoCodec
	v.ScreenSize = g.ScreenSize
	v.Format = g.Format
	v.MediaType = g.Type
	v.Container = g.Container
	v.Languages = g.Languages[:]
	v.SubtitleLanguages = g.SubtitleLanguages[:]
	v.Title = g.Title
	return v
}
