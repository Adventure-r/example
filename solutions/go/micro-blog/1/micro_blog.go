package microblog

import (
	"strings"
)

func Truncate(phrase string) string {
	var result strings.Builder
	runes := []rune(phrase)
	for i := 0; i < len(runes) && i < 5; i++ {
		result.WriteRune(runes[i])
	}
	return result.String()
}
