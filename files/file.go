// Package files contains structures and functions for dealing with files on
// the local file system, in terms of mmmgr, such as determining the mime-type
// and the media type, and extracting media information from file names.
package files

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// A File contains information about a file's absolute path, media types, and
// information gleaned from the path.
type File struct {
	path                                    string
	absPath, mimeType, basicType, mediaType string
	mediaInfo                               map[string]string
	fileInfo                                os.FileInfo
}

// Path gets the File's path
func (f *File) Path() string {
	if f.path == "" {
		log.Fatalf("File's path is not set")
	}
	return f.path
}

// AbsPath gets the absolute path to the File
func (f *File) AbsPath() string {
	if f.absPath != "" {
		return f.absPath
	}
	path, err := AbsolutePath(f.Path())
	if err != nil {
		log.Fatalf("Couldn't get absolute path of %q - %v", f.path, err)
	}
	return path
}

// Ancestry gets the filepath ancestry of the file without an extension.. See
// Ancestry.
func (f *File) Ancestry() []string {
	return Ancestry(StripExt(f.AbsPath()))
}

// Exists tells whether a file actually exists.
func (f *File) Exists() bool {
	return f.FileInfo() != nil
}

// IsFile returns true only if the File exists and is an ordinary file
func (f *File) IsFile() bool {
	return !f.FileInfo().IsDir()
}

// FileInfo gets the underlying file's info (from os.Stat).
func (f *File) FileInfo() os.FileInfo {
	if f.fileInfo == nil {
		fileInfo, err := os.Stat(f.Path())
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
			log.Fatalf("Failed to stat file %q", f.path)
		}
		f.fileInfo = fileInfo
	}
	return f.fileInfo
}

// MimeType gets the underlying file's mime type.
func (f *File) MimeType() string {
	if f.mimeType == "" {
		f.mimeType = MimeType(f.Path())
	}
	return f.mimeType
}

// BasicType gets the underlying file's basic type ("video", "audio", "book"
// or "").
func (f *File) BasicType() string {
	if f.basicType == "" {
		if f.MimeType() != "" {
			f.basicType = BasicType(f.MimeType())
		}
	}
	return f.basicType
}

// MediaType returns the media type of the file, one of "book", "movie", "tv",
// "music" or "".
func (f *File) MediaType() string {
	if f.mediaType == "" {
		f.mediaType, f.mediaInfo = f.getMediaInfo()
	}
	return f.mediaType
}

// MediaInfo returns the mediaInfo map of the file (or nil, if none could be
// found).
func (f *File) MediaInfo() map[string]string {
	if f.mediaInfo == nil {
		f.mediaType, f.mediaInfo = f.getMediaInfo()
	}
	return f.mediaInfo
}

// getMediaInfo looks up sets of Regexps for the File's basicType, and tries
// to match a regular expression in one of the sets therein to the File's path
// ancestry. If found, returns the mediaType (as a string) and the mediaInfo
// (as a map of strings). Return an empty string and nil if no matching media
// info could be found.
func (f *File) getMediaInfo() (string, map[string]string) {
	if types := Types[f.BasicType()]; types != nil {
		for _, mediaType := range types {
			if set := Regexps[mediaType]; set != nil {
				for _, r := range set {
					if mediaInfo := MatchFirst(r, f.Ancestry()); mediaInfo != nil {
						return mediaType, mediaInfo
					}
				}
			}
		}
	}
	return "", nil
}

// New constructs a new File from a path p.
func New(p string) *File {
	f := new(File)
	f.path = p
	return f
}

// AbsolutePath gets the absolute path of a file path p.
func AbsolutePath(p string) (string, error) {
	return filepath.Abs(p)
}

// Stat gets an os.FileInfo for a path p.
func Stat(p string) (os.FileInfo, error) {
	return os.Stat(p)
}

// IsDir returns true if the path exists and points to a directory.
func IsDir(path string) bool {
	if info, err := Stat(path); err == nil {
		return info.IsDir()
	}
	return false
}

// IsMediaFile returns true if the path exists and points to a media file.
func IsMediaFile(path string) bool {
	f := New(path)
	return f.Exists() && f.IsFile() && f.MediaType() != ""
}

// Scan scans the given directory for media files, and returns a channel which
// gets a *File or an error for every media file found.
func Scan(path string) <-chan *File {
	out := make(chan *File)
	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f := New(path)
		if f.MediaType() == "" {
			return nil
		}
		out <- f
		return nil
	}
	go func() {
		if err := filepath.Walk(path, walker); err != nil {
			log.Fatalf("Error walking filepath %q: %v", path, err)
		}
		close(out)
	}()
	return out
}

// Exists returns true if path exists, false otherwise.
func Exists(path string) bool {
	_, err := Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Fatalf("Couldn't determine if %q exists - %v", path, err)
	panic("unreachable")
}

// StripExt returns the pathname without the extension.
func StripExt(p string) string {
	return strings.TrimSuffix(p, filepath.Ext(p))
}

// Ancestry returns a list of paths of increasing specificity for the given
// path.  For example, given "/a/b/c.x", it returns {"c.x", "b/c.x",
// "a/b/c.x", "/a/b/c.x"}. This allows partial paths to be tested against
// regular expressions before trying ever more specific paths.
func Ancestry(p string) []string {
	stack := []string{}
	parts := strings.Split(filepath.Clean(p), "/")
	n := len(parts)
	for i := range parts[:n] {
		stack = append(stack, filepath.Join(parts[n-i-1:]...))
	}
	return stack
}

// findMatch looks through a slice of *regexp.Regexps for the one which
// matches the given path, by checking the path's ancestry (i.e., starting
// with the basename without the extension, then adding parent directories).
// Returns the matching *regexp.Regexp if found, otherwise returns nil.
func findMatch(path string, rexes []*regexp.Regexp) (*regexp.Regexp, string) {
	ancestries := Ancestry(path)
	for _, rex := range rexes {
		for _, p := range ancestries {
			if rex.MatchString(p) {
				return rex, p
			}
		}
	}
	return nil, ""
}
