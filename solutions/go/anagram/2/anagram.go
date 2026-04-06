package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	anagrams := make([]string, 0)

	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func isAnagram(source, target string) bool {
	if len(source) != len(target) {
		return false
	}
	source = strings.ToLower(source)
	target = strings.ToLower(target)

	if source == target {
		return false
	}
	sourceMap := countLetters(source)
	targetMap := countLetters(target)
	for k, v := range sourceMap {
		if targetMap[k] != v {
			return false
		}
	}
	return true
}

func countLetters(source string) map[rune]uint8 {
	letters := make(map[rune]uint8, len(source))
	for _, l := range source {
		letters[l]++
	}
	return letters
}
