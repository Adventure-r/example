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
	var abbr strings.Builder
	abbr.Grow(len(s) / 3)
	lastWasSeparator := true
	for _, r := range s {
		if !unicode.IsLetter(r) && r != '\'' {
			lastWasSeparator = true
		} else if lastWasSeparator {
			abbr.WriteRune(r)
			lastWasSeparator = false
		}
	}
	result := strings.ToUpper(abbr.String())
	return result
}
