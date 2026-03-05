package reversestring

import (
	"strings"
)

func Reverse(input string) string {
	var revert strings.Builder
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		revert.WriteRune(runes[i])
	}
	return revert.String()
}
