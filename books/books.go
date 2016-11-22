package books

// Holds data related to an eBook
type Book struct {
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
