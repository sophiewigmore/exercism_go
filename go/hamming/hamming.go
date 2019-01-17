package hamming

import (
	"errors"
)

// Distance: calculate the hamming distance between two strings
// If the strings are unequal length, throw an error
func Distance(a, b string) (int, error) {
	distance := 0
	// Check that the strings are of equal length
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				distance++
			}
		}
		return distance, nil
	}
	// If strings are of unequal length, throw error
	return distance, errors.New("error - couldn't calculate string length")
}
