package file

import (
	"github.com/bjjb/mmmgr/books"
	"github.com/bjjb/mmmgr/movies"
	"github.com/bjjb/mmmgr/tv"
	"mime"
	"path/filepath"
	"strings"
)

// A File just wraps a file on the local filesystem
type File struct {
	Path      string
	MimeType  string
	MediaType string
}

// Constructs a new File
func New(path string) *File {
	return &File{path, MimeType(path), MediaType(path)}
}

// Returns true if the path is to an audio, video or book file.
func IsMedia(path string) bool {
	switch MediaType(path) {
	case "audio", "video", "book", "movie", "tv", "music":
		return true
	}
	return false
}

// MediaType takes a path and returns the type of media associated with it.
// If the file isn't any known media type, an empty string is returned.
func MimeType(path string) string {
	return mime.TypeByExtension(filepath.Ext(path))
}

// True if the given file path is a video
func IsVideo(path string) bool {
	return strings.Split(MimeType(path), "/")[0] == "video"
}

// True if the given file path looks like a TV show
func IsTV(path string) bool {
	return tv.GuessFromPath(path) != nil
}

// True if the given file path looks like a movie
func IsMovie(path string) bool {
	return movies.GuessFromPath(path) != nil
}

// True if the given file path is an audio file
func IsAudio(path string) bool {
	return strings.Split(MimeType(path), "/")[0] == "audio"
}

// True if the given file path looks like an music file
func IsMusic(path string) bool {
	return true
}

// True if the given file path is a book
func IsBook(path string) bool {
	mimeType := MimeType(path)
	for _, bookType := range books.MimeTypes {
		if mimeType == bookType {
			return true
		}
	}
	return false
}

// For the given path, returns one of the available types ('video', 'tv',
// 'movie', 'audio', 'music' or 'book'), or '' (if none of the above).
func MediaType(path string) string {
	if IsVideo(path) {
		if IsTV(path) {
			return "tv"
		}
		if IsMovie(path) {
			return "movie"
		}
		return "video"
	}
	if IsAudio(path) {
		if IsMusic(path) {
			return "music"
		}
		return "audio"
	}
	if IsBook(path) {
		return "book"
	}
	return ""
}
