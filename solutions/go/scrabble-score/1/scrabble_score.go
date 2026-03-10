package scrabblescore

import "strings"

func Score(word string) int {
	word = strings.ToUpper(word)
	runes := []rune(word)
	points := 0
	for _, r := range runes {
		switch r {
		case 'A':
			points += 1
		case 'E':
			points += 1
		case 'I':
			points += 1
		case 'O':
			points += 1
		case 'U':
			points += 1
		case 'L':
			points += 1
		case 'N':
			points += 1
		case 'R':
			points += 1
		case 'S':
			points += 1
		case 'T':
			points += 1
		case 'D':
			points += 2
		case 'G':
			points += 2
		case 'B':
			points += 3
		case 'C':
			points += 3
		case 'M':
			points += 3
		case 'P':
			points += 3
		case 'F':
			points += 4
		case 'H':
			points += 4
		case 'V':
			points += 4
		case 'W':
			points += 4
		case 'Y':
			points += 4
		case 'K':
			points += 5
		case 'J':
			points += 8
		case 'X':
			points += 8
		case 'Q':
			points += 10
		case 'Z':
			points += 10

		}
	}
	return points
}
