package luhn

import (
	"strconv"
)

// Valid is a function that takes in a string and returns if
// it is valid using a Luhn algorithm
func Valid(input string) bool {
	// check if input is only numeric or not
	s, e := strconv.Atoi(input)

	if len(input) <= 1 || e != nil {
		return false
	}

	runes := []rune(input)
	// Step 1: double every second digit starting from the right
	for i := len(runes) - 2; i >= 0; i = i - 2 {
		newNum := runes[i] * 2
		if newNum > 9 {
			newNum = newNum - 9
		}
		runes[i] = newNum

		// Step 2: Add all the digits
		// Here we will add the current letter, and the one previous to it

	}

}
