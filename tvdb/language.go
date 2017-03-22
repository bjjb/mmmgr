package tvdb

// A Language encapsulates a language supported by The TVDB.
type Language struct {
	ID          int    `json:"id"`
	Abbr        string `json:"abbreviation"`
	Name        string `json:"name"`
	EnglishName string `json:"englishName"`
}
