package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)
	alphabet := make(map[rune]bool)
	characterCount := 0
	for _, v := range input {
		if !unicode.IsLetter(v) {
			continue
		}
		if _, exists := alphabet[v]; !exists {
			alphabet[v] = true
			characterCount++
		}
	}
	if characterCount != 26 {
		return false
	}
	return true
}
