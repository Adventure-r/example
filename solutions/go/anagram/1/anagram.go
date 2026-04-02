package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	anagrams := make([]string, 0, len(candidates)/2)
	subject = strings.ToLower(subject)
	for i, cand := range candidates {
		if len(subject) != len(cand) {
			continue
		}
		cand = strings.ToLower(cand)
		if cand == subject {
			continue
		}
		rCand := []rune(cand)
		length := len(rCand)
		letterCounter := 0
		rSubject := []rune(subject)
		for _, letter := range rCand {
			for i, subLetter := range rSubject {
				if letter == subLetter {
					letterCounter++
					rSubject[i] = ' ' // Чото надо улучшить
					break
				}
			}
			if letterCounter == length {
				anagrams = append(anagrams, candidates[i])
			}
		}
	}
	return anagrams
}
