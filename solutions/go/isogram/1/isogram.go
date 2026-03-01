package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ReplaceAll(strings.ToLower(word), "-", "")
	word = strings.ReplaceAll(word, " ", "")
	for i, v := range word {
		for innerI, innerV := range word {
			if i == innerI {
				continue
			}
			if v == innerV {
				return false
			}
		}
	}
	return true
}
