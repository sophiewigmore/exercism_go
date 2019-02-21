package pangram

import (
	"strings"
	"unicode"
)

// IsPangram returns if the given input string is a Pangram
func IsPangram(input string) bool {
	lettersSeen := map[rune]int{}

	input = strings.ToUpper(input)

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}

		_, ok := lettersSeen[r]

		if ok {
			lettersSeen[r] = lettersSeen[r] + 1
		}
		lettersSeen[r] = 0
	}

	return len(lettersSeen) == 26
}
