package audio

import (
	"path/filepath"
	"strings"
)

// Holds data related to an audio file
type Audio struct {
	Title    string
	MimeType string
	Codec    string
}

func GuessFromPath(path string) *Audio {
	baseName := filepath.Base(path)
	n := strings.LastIndexByte(baseName, '.')
	title := baseName[:n]
	return &Audio{
		Title: title,
	}
}
