package files

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Replacements is the list of substitutions which can be inserted into
// patterns.
var Replacements map[string]string

// Patterns is set of a lists of file matching patterns, per media type.
var Patterns map[string][]string

// Types enumerates the media types of each basic type.
var Types = map[string][]string{
	"video": []string{"tv", "movie"},
	"audio": []string{"music"},
	"book":  []string{"book"},
}

// expandPattern Expands a given pattern string into a *regexp.Regexp by
// looking through the keys in the replacement map (rm) and replacing all
// occurances of that replacement string (surrounded by '%') with the
// corresponding value enclosed in a submatch expression, and then compiling
// the resultant regular expression (or panicking). For example:
//   expandPattern("-%x% %y%-", map[string]string{"x": "fo+", "y": "bar"})
// expands to a *regexp.Regexp whose text would be: "-(?P<x>fo+) (?P<y>bar)-"
func expandPattern(p string, rm map[string]string) string {
	s := p
	for r := range rm {
		t := "(?P<" + r + ">" + rm[r] + ")"
		s = strings.Replace(s, "%"+r+"%", t, -1)
	}
	return s
}

// Match matches s to a regular expression, returning a map of strings from
// the submatches. Returns nil if re doesn't match s.
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

// MatchFirst calls Match on the list of strings, returning the first
// successful set of matches.
func MatchFirst(re *regexp.Regexp, ss []string) map[string]string {
	for _, s := range ss {
		if r := Match(re, s); r != nil {
			return r
		}
	}
	return nil
}

// Regexps is a map of regular expressions per media type, which is compiled
// at init-time from the expanded contents of Patterns.
var Regexps = map[string][]*regexp.Regexp{}

// Compiles Regexps as specified in Patterns and Replacements.
func init() {
	Replacements = DefaultReplacements
	Patterns = DefaultPatterns
	for mediaType := range Patterns {
		list := []*regexp.Regexp{}
		for _, p := range Patterns[mediaType] {
			s := "(?i:" + expandPattern(p, Replacements) + ")"
			re, err := regexp.Compile(s)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error compiling regexp %q", s)
			} else {
				list = append(list, re)
			}
		}
		Regexps[mediaType] = list
	}
}
