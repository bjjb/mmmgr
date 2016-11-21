package file

import (
	"github.com/bjjb/mmmgr/books"
	"mime"
	"path/filepath"
	"regexp"
	"strings"
)

type File struct {
	Path      string
	MimeType  string
	MediaType string
	Error     error
}

// Constructs a new File
func New(path string) *File {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return &File{Error: err}
	}
	return &File{absPath, MimeType(absPath), MediaType(absPath), nil}
}

// Returns true if the path is to an audio, video or book file.
func IsMedia(path string) bool {
	switch MediaType(path) {
	case "audio", "video", "book", "movie", "tv", "music":
		return true
	}
	return false
}

// Holds data related to a video file
type Video struct {
	File
	Duration string
}

// Holds data related to an audio file
type Audio struct {
	File
	Duration string
}

// Holds data related to an eBook
type Book struct {
	File
	Title  string
	Author string
	Date   string
	Series string
}

// Holds data related to a Music file
type Music struct {
	Audio
	Artist string
	Album  string
	Title  string
	Disc   string
	Number string
	MBID   string
}

// Holds data related to a TV file
type TV struct {
	Video
	Title  string
	Show   string
	Season string
	Number string
	Year   string
	TVDBID string
	TMDBID string
}

// Holds data related to a Movie file
type Movie struct {
	Video
	Title  string
	Year   string
	TMDBID string
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

// True if the given file path is a TV show
func IsTV(path string) bool {
	tvPatterns := []string{
		`[sS]\d{1,2}[eE]\d{1,2}`,
		`\d{1,2}[x.]\d{1,2}`,
	}
	baseName := filepath.Base(path)
	//parentDir := filepath.Base(filepath.Dir(path))
	for _, pattern := range tvPatterns {
		if regexp.MustCompile(pattern).MatchString(baseName) {
			return true
		}
	}
	return false
}

// True if the given file path looks like a movie
func IsMovie(path string) bool {
	return false
}

// True if the given file path is an audio file
func IsAudio(path string) bool {
	return strings.Split(MimeType(path), "/")[0] == "audio"
}

// True if the given file path looks like an music file
func IsMusic(path string) bool {
	return false
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
