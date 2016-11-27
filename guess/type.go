package guess

import (
	"log"
	"mime"
	"path/filepath"
	"regexp"
	"strings"
)

type Question func(string) Answer
type Answer int

const (
	NO int = iota
	YES
	MAYBE
)

var questions = struct{ tv, movie, music []Question }{
	[]Question{},
	[]Question{},
	[]Question{},
}

var patterns = struct{ tv, movie, music []string }{
	[]string{ // tv
		`[sS]\d{1,2}[eE]\d{1,3}`,
		`\d{1,2}[xX]\d{1,3}`,
	},
	[]string{ // movie
		`\d{4}`,
	},
	[]string{ // music
	},
}

var Patterns struct{ TV, Movie, Music []*regexp.Regexp }
var Tests struct{ TV, Movie, Music, Book []Question }

func init() {
	Patterns.TV = compilePatterns(patterns.tv)
	Patterns.Movie = compilePatterns(patterns.movie)
	Patterns.Music = compilePatterns(patterns.music)
}

func MimeType(path string) string {
	return mimeType(path)
}

func Type(path string) string {
	mime := mimeType(path)
	parts := strings.Split(mime, "/")
	switch parts[0] {
	case "audio":
		return audioType(path)
	case "video":
		return videoType(path)
	case "application":
		switch parts[1] {
		case "epub+zip", "x-mobi-pocketbook", "pdf":
			return "book"
		}
	}
	return ""
}

// If the path can be determined to be a TV show episode, returns "tv"; if
// not, but it looks like a film, returns "movie". Otherwise returns "video".
func videoType(path string) string {
	if strings.Split(mimeType(path), "/")[0] != "video" {
		log.Fatalf("guess.videoType called on a non-video file %q", path)
	}
	for _, re := range Patterns.TV {
		if re.Match([]byte(path)) {
			return "tv"
		}
	}
	for _, re := range Patterns.Movie {
		if re.Match([]byte(path)) {
			return "movie"
		}
	}
	return "video"
}

// If the path can be determined to be a song, returns "music". Otherwise,
// returns "audio"
func audioType(path string) string {
	if strings.Split(mimeType(path), "/")[0] != "audio" {
		log.Fatalf("guess.audioType called on a non-audio %q", path)
	}
	for _, re := range Patterns.Music {
		if re.Match([]byte(path)) {
			return "music"
		}
	}
	return "audio"
}

// Gets the mime-type of the file at the given path
func mimeType(path string) string {
	return mime.TypeByExtension(filepath.Ext(path))
}

// Converts a slice of strings to a slice of regexp.Regexp, with
// regexp.MustCompile.
func compilePatterns(in []string) []*regexp.Regexp {
	out := make([]*regexp.Regexp, len(in))
	for i, re := range in {
		out[i] = regexp.MustCompile(re)
	}
	return out
}
