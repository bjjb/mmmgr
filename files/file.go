package files

import (
	"mime"
	"path/filepath"
	"strings"
	"log"
)

// A File contains information about a file's absolute path, media type and
// mime-type.
type File struct {
	Path      string `json:"path"`
	MimeType  string `json:"mime_type"`
	MediaType string `json:"type"`
}

// Constructs a new File from a path
func New(path string) *File {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("couldn't get absolute path of %q - %v", path, err)
	}
	mime := mimeType(absPath)
	return &File{absPath, mime, mediaType(path)}
}

// Gets the mime-type of the file at the given path
func mimeType(path string) string {
	return mime.TypeByExtension(filepath.Ext(path))
}

// Gets the media type of the file at the given path. Will be `video`, `audio`
// or `book`.
func mediaType(path string) string {
	mime := mimeType(path)
	base := strings.Split(mime, "/")[0]
	// Guess based on the first part of the mime-type (audio or video)
	switch base {
	case "audio", "video":
		return base
	default:
		if _, found := bookMimeTypes[mime]; found {
			return "book"
		}
	}
	// TODO: books (based on the rest of the mime type)
	return ""
}

var bookMimeTypes = map[string]string{}
