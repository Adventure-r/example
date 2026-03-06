package rnatranscription

import "strings"

func ToRNA(dna string) string {
	runes := []rune(dna)
	var result strings.Builder
	for _, r := range runes {
		switch r {
		case 'G':
			result.WriteRune('C')
		case 'C':
			result.WriteRune('G')
		case 'T':
			result.WriteRune('A')
		case 'A':
			result.WriteRune('U')
		}
	}
	return result.String()
}
