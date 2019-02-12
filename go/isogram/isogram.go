package isogram

import "unicode"

// IsIsogram returns if a string is an isogram
func IsIsogram(input string) bool {
	if len(input) == 0 {
		return true
	}
	m := make(map[rune]bool)
	letters := []rune(input)

	for _, letter := range letters {
		letter = unicode.ToLower(letter)
		if letter == ' ' || letter == '-' {
			continue
		}
		if m[letter] {
			return false
		}
		m[letter] = true

	}
	return true
}
