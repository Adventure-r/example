package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	for _, v := range id {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	var sum int
	var startIndex int
	if len(id) == 1 {
		return false
	}
	if len(id)%2 == 0 {
		startIndex = 0
	} else {
		sum += int(id[0] - '0')
		startIndex = 1
	}

	for i := startIndex; i < len(id); i += 2 {
		multBy2 := int(id[i]-'0') * 2
		if multBy2 > 9 {
			multBy2 -= 9
		}
		sum += multBy2
		sum += int(id[i+1] - '0')
	}
	return sum%10 == 0
}
