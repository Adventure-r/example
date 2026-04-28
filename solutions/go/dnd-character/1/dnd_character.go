package dndcharacter

import "math/rand"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	result := score - 10
	if score%2 != 0 {
		if result < 0 {
			result--
		}
	}
	return result / 2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	min := 100 // bigger than any value
	sum := 0
	for range 4 {
		n := rand.Intn(6) + 1
		if n < min {
			min = n
		}
		sum += n
	}
	sum -= min
	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
