package scrabble

import "strings"

// Score is a method that computes the scrabble score of a word
func Score(words string) int {
	pointConversion := GetPointsMap()
	points := 0
	for i := 0; i < len(words); i++ {
		points += pointConversion[strings.ToUpper(words[i:i+1])]
	}
	return points
}

//GetPointsMap returns the mapping of Letter to Point conversion
func GetPointsMap() map[string]int {
	pointConversion := make(map[string]int)
	pointConversion["A"] = 1
	pointConversion["E"] = 1
	pointConversion["I"] = 1
	pointConversion["O"] = 1
	pointConversion["U"] = 1
	pointConversion["L"] = 1
	pointConversion["N"] = 1
	pointConversion["R"] = 1
	pointConversion["S"] = 1
	pointConversion["T"] = 1
	pointConversion["D"] = 2
	pointConversion["G"] = 2
	pointConversion["B"] = 3
	pointConversion["C"] = 3
	pointConversion["M"] = 3
	pointConversion["P"] = 3
	pointConversion["F"] = 4
	pointConversion["H"] = 4
	pointConversion["V"] = 4
	pointConversion["W"] = 4
	pointConversion["Y"] = 4
	pointConversion["K"] = 5
	pointConversion["J"] = 8
	pointConversion["X"] = 8
	pointConversion["Q"] = 10
	pointConversion["Z"] = 10
	return pointConversion
}
