/*
Contains structures and functions for dealing with files on the local file
system, in terms of mmmgr, such as determining the mime-type and the media
type. For more detailed information, look at guess, or the tv, movies, books
or music packages.
*/
package files

import (
	"github.com/bjjb/mmmgr/guess"
	"log"
	"path/filepath"
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
	mime := guess.MimeType(absPath)
	media := guess.Type(absPath)
	return &File{absPath, mime, media}
}
