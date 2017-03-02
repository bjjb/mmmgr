package files

import (
	"regexp"
	"strings"
)

/*
Patterns is the master map of media patterns.
*/
var Patterns = map[string][]string{
	"tv": []string{
		`%show%\W\(?%year%\)?\WS%season%E%episode%\W%title%`,
		`%show%\WS%season%E%episode%\W%title%`,
	},
	"movie": []string{
		`%title%\W\(?%year%\)?/%title%\W\(?%year%\)?`,
		`%title%\W\(?%year%\)?`,
	},
	"music": []string{},
	"book":  []string{},
}

/*
Types enumerates the media types of each basic type.
*/
var Types = map[string][]string{
	"video": []string{"tv", "movie"},
	"audio": []string{"music"},
	"book":  []string{"book"},
}

/*
Replacements is a map of regular expression snippets which can be used to
abbreviate or specify filename patterns.
*/
var Replacements = map[string]string{
	"year":    `\d{4}`,
	"title":   `[^/]+`,
	"season":  `\d{2}`,
	"part":    `\d+`,
	"episode": `\d{2}`,
	"date":    `(?:\d{4}-\d{2}-\d{2}|\d{2}\.\d{2}\.\d{4})`,
	"show":    `[^/]+`,
	"disc":    `\d{2}`,
	"track":   `\d{2}`,
}

/*
expandPattern Expands a given pattern string into a *regexp.Regexp by looking
through the keys in the replacement map (rm) and replacing all occurances of
that replacement string (surrounded by '%') with the corresponding value
enclosed in a submatch expression, and then compiling the resultant regular
expression (or panicking). For example:

	expandPattern("-%x% %y%-", map[string]string{"x": "fo+", "y": "bar"})

expands to a *regexp.Regexp whose text would be: "-(?P<x>fo+) (?P<y>bar)-"
*/
func expandPattern(p string, rm map[string]string) string {
	s := p
	for r := range rm {
		t := "(?P<" + r + ">" + rm[r] + ")"
		s = strings.Replace(s, "%"+r+"%", t, -1)
	}
	return s
}

/*
Match matches s to a regular expression, returning a map of strings from the
submatches. Returns nil if re doesn't match s.
*/
func Match(re *regexp.Regexp, s string) map[string]string {
	if matches := re.FindStringSubmatch(s); matches != nil {
		results := map[string]string{}
		for i, k := range re.SubexpNames() {
			results[k] = matches[i]
		}
		return results
	}
	return nil
}

/*
MatchFirst calls Match on the list of strings, returning the first successful
set of matches.
*/
func MatchFirst(re *regexp.Regexp, ss []string) map[string]string {
	for _, s := range ss {
		if r := Match(re, s); r != nil {
			return r
		}
	}
	return nil
}

/*
Regexps is a map of regular expressions per media type, which is compiled at
init-time from the expanded contents of Patterns.
*/
var Regexps = map[string][]*regexp.Regexp{}

/*
Compiles Regexps as specified in Patterns and Replacements.
*/
func init() {
	for mediaType := range Patterns {
		list := []*regexp.Regexp{}
		for _, p := range Patterns[mediaType] {
			s := expandPattern(p, Replacements)
			list = append(list, regexp.MustCompile(s))
		}
		Regexps[mediaType] = list
	}
}
