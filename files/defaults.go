package files

/*
DefaultPatterns is the default master map of media patterns.
*/
var DefaultPatterns = map[string][]string{
	"tv": []string{
		`%show%\.S%season%E%episode%(?:\.%title%)?\.%source%\.%vcodec%-%rgroup%`,
		`%show%\W\(?%year%\)?\WS%season%E%episode%\W%title%`,
		`%show%\WS%season%E%episode%\W%title%`,
	},
	"movie": []string{
		`%title%\W\(?%year%\)?/%title%\W\(?%year%\)?`,
		`%title%\W\(?%year%\)?`,
	},
	"music": []string{
		`%artist%/%album%/%track%\W%title%`,
	},
	"book": []string{},
}

/*
DefaultReplacements is a map of regular expression snippets which can be used
to abbreviate or specify filename patterns.
*/
var DefaultReplacements = map[string]string{
	"album":   `[^/]+`,
	"artist":  `[^/]+`,
	"date":    `(?:\d{4}-\d{2}-\d{2}|\d{2}\.\d{2}\.\d{4})`,
	"disc":    `\d{2}`,
	"episode": `\d{2}`,
	"part":    `\d+`,
	"season":  `\d{2}`,
	"show":    `[^/]+`,
	"title":   `[^/]+`,
	"track":   `\d{2}`,
	"year":    `\d{4}`,
	"rgroup":  `[^ ./]+`,
	"vcodec":  `x264`,
	"source":  `WEBRip`,
}
