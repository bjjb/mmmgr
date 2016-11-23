package files

import (
	"mime"
	"path/filepath"
	"strings"
	"log"
	"github.com/bjjb/mmmgr/guessit"
)

// A File just wraps a file on the local filesystem
type File struct {
	Path      string `json:"path"`
	MimeType  string `json:"mime_type"`
	MediaType string `json:"type"`
}

// Constructs a new File
func New(path string) *File {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("couldn't get absolute path of %q - %v", path, err)
	}
	return &File{absPath, mimeType(path), mediaType(path)}
}

// mimeType takes a path and returns the mime-type.
func mimeType(path string) string {
	return mime.TypeByExtension(filepath.Ext(path))
}

// For the given path, returns one of the available types ('video', 'tv',
// 'movie', 'audio', 'music' or 'book'), or '' (if none of the above).
func mediaType(path string) string {
	mime := mimeType(path)
	// Guess based on the first part of the mime-type (audio or video)
	switch strings.Split(mime, "/")[0] {
	case "audio":
		// TODO: music
		return "audio"
	case "video":
		guessedType := guessit.Guessit(path).Type
		switch guessedType {
		case "episode":
			return "tv"
		case "movie":
			return "movie"
		default:
			log.Fatalf("Unexpected type from Guessit: %q", guessedType)
		}
		return "video"
	}
	// TODO: books (based on the rest of the mime type)
	return ""
}
