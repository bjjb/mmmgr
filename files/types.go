package files

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"
)

// MimeTypes is an map of overridden mime types (per file extension).
var MimeTypes = map[string]string{
	".opus": "audio/opus",
}

// MimeType returns the (best-guess) MIME type for the file. Extracts the
// path's file extension (including the '.'), and looks it up in MimeTypes. If
// not found, it uses mime.TypeByExtension, stripping off anything after the
// first ';'.
func MimeType(path string) string {
	ext := filepath.Ext(path)
	if m, found := MimeTypes[ext]; found {
		return m
	}
	return strings.Split(mime.TypeByExtension(ext), ";")[0]
}

// BasicType determines the basic type corresponding to a particular mime
// type.  This will be one of "video", "audio", "book" or "" (the latter if
// the mime type isn't that of an interesting media).
func BasicType(mimeType string) string {
	parts := strings.Split(mimeType, "/")
	if len(parts) != 2 {
		panic(fmt.Sprintf("invalid MIME type %q", mimeType))
	}
	switch parts[0] {
	case "audio", "video":
		return parts[0]
	case "application":
		switch strings.Split(parts[1], ";")[0] {
		case "epub+zip", "x-mobi-pocketbook", "x-mobipocket-ebook", "pdf":
			return "book"
		}
	}
	return ""
}
