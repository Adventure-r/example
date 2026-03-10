// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	runes := []rune(s)
	var abbr strings.Builder
	if len(runes) > 0 {
		abbr.WriteRune(runes[0])
	}
	for i, r := range runes {
		if !unicode.IsLetter(r) && r != '\'' && unicode.IsLetter(runes[i+1]) {
			abbr.WriteRune(runes[i+1])
		}
	}
	result := strings.ToUpper(abbr.String())
	return result
}
