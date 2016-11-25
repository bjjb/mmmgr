/*
Contains structures and functions for dealing with files on the local file
system, in terms of mmmgr, such as determining the mime-type and the media
type. For more detailed information, look at guess, or the tv, movies, books
or music packages.
*/
package files

import (
	"log"
	"mime"
	"path/filepath"
	"regexp"
	"strings"
)

// A File contains information about a file's absolute path, media type and
// mime-type.
type File struct {
	Path      string `json:"path"`
	MimeType  string `json:"mime_type"`
	MediaType string `json:"type"`
}

type reList []*regexp.Regexp

func init() {
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
// or `book`, unless it can determine that a video is 'movie' or 'tv', or that
// an 'audio' is 'music'.
func mediaType(path string) string {
	mime := mimeType(path)
	base := strings.Split(mime, "/")[0]
	rest := strings.Split(mime, "/")[1]
	// Guess based on the first part of the mime-type (audio or video)
	switch base {
	case "audio", "video":
		return base
	default:
		switch rest {
		case "epub+zip", "x-mobi-pocketbook", "pdf":
			return "book"
		}
	}
	// TODO: books (based on the rest of the mime type)
	return ""
}

// Compiles a list of patterns into regular expressions
func compileREs(patterns []string) *reList {
	list := make(reList, len(patterns))
	for _, pattern := range patterns {
		list = append(list, regexp.MustCompile(pattern))
	}
	return &list
}
