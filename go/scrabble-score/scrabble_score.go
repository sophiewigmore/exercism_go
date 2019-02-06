package scrabble

import "unicode"

// pointConversion is the map for letter -> points conversion
var pointConversion = GetPointsMap()

// Score is a method that computes the scrabble score of a word
func Score(words string) int {
	points := 0
	// turn the string into runes
	wordAsRunes := []rune{}
	for _, r := range words {
		wordAsRunes = append(wordAsRunes, r)
	}
	for i := 0; i < len(wordAsRunes); i++ {
		points += pointConversion[unicode.ToUpper(wordAsRunes[i])]
	}
	return points
}

//GetPointsMap returns the mapping of Letter to Point conversion
func GetPointsMap() map[rune]int {
	pointConversion := map[rune]int{
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
		'D': 2, 'G': 2,
		'B': 3, 'C': 3, 'M': 3, 'P': 3,
		'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
		'K': 5,
		'J': 8, 'X': 8,
		'Q': 10, 'Z': 10}
	return pointConversion
}
