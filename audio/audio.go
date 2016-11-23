package audio

import (
	"github.com/bjjb/mmmgr/files"
	"path/filepath"
	"strings"
)

// Holds data related to an audio file
type Audio struct {
	File		 *files.File
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

// Makes a new Audio from a files.File.
func New(path string) *Audio {
	r := new(Audio)
	return r
}
