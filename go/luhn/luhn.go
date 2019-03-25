package luhn

import (
	"strconv"
	"strings"
)

// Valid is a function that takes in a string and returns if
// it is valid using a Luhn algorithm
func Valid(input string) bool {
	//remove spaces
	input = strings.Replace(input, " ", "", -1)
	// check if input is only numeric or not
	_, e := strconv.Atoi(input)
	if len(input) <= 1 || e != nil {
		return false
	}

	digitSum := 0
	// Step 1: double every second digit starting from the right
	for i, c := range input {
		digitAsInt, _ := strconv.Atoi(string(c))
		newNum := digitAsInt

		if (len(input)-i)%2 == 0 {
			newNum = digitAsInt * 2
			if newNum > 9 {
				newNum = newNum - 9
			}
		}
		// Step 2: Add all the digits
		// Here we will add the current number to the growing numbrt
		digitSum += newNum
	}

	//Step 3: If the digit sum is evenly divisble by 10 -> return true
	// Otherwise return true
	if digitSum%10 != 0 {
		return false
	}

	return true
}
