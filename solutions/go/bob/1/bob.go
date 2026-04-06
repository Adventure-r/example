// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	upper := strings.ToUpper(remark)
	lower := strings.ToLower(remark)
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case remark[len(remark)-1] == '?' && remark == upper && upper != lower:
		return "Calm down, I know what I'm doing!"
	case remark[len(remark)-1] == '?':
		return "Sure."
	case remark == upper && upper != lower:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
