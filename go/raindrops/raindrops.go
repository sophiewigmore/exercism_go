package raindrops

import (
	"strconv"
)

//Convert : converts a number to a string:
/**
* If the number has 3 as a factor, output 'Pling'.
* If the number has 5 as a factor, output 'Plang'.
* If the number has 7 as a factor, output 'Plong'.
* If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
 */
func Convert(x int) string {
	converted := ""
	if x%3 == 0 {
		converted += "Pling"
	}
	if x%5 == 0 {
		converted += "Plang"
	}
	if x%7 == 0 {
		converted += "Plong"
	}
	if len(converted) == 0 {
		converted = strconv.Itoa(x)
	}
	return converted
}
