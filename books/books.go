package books

import (
	"github.com/bjjb/mmmgr/files"
)

// Holds data related to an eBook
type Book struct {
	File	 *files.File
	Title  string
	Author string
	Date   string
	Series string
}

var MimeTypes = [2]string{
	"application/x-mobipocket-ebook",
	"application/epub+zip",
}

func GuessFromPath(path string) *Book {
	return nil
}

// Makes a new Book from a files.File.
func New(path string) *Book {
	r := new(Book)
	return r
}
