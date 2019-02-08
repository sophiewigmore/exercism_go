package isogram

import "strings"

// IsIsogram returns if a string is an isogram
func IsIsogram(input string) bool {
	if len(input) == 0 {
		return true
	}
	m := make(map[rune]int)
	letters := []rune(strings.ToLower(input))

	for _, letter := range letters {
		if letter != '-' && letter != ' ' {
			_, ok := m[letter]
			if !ok {
				m[letter] = 0
			} else {
				return false
			}
		}
	}
	return true
}
