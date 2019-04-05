package grains

import (
	"errors"
)

// Calculate the number of grains of wheat on a
// chessboard given that the number on each square doubles.
// There are 64 squares on a chess board

// Total returns the total number of grains
func Total() uint64 {
	var count uint64
	var total uint64

	for i := 1; i <= 64; i++ {
		count, _ = Square(i)
		total = total + count
	}
	return total
}

// Square how many grains were on each square
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 1, errors.New("invalid input")
	}
	// 2 ^ (square number-1) = number of grains
	//return uint64(math.Pow(2, float64(square-1))), nil

	// Use bit shifting instead of math.Pow
	return uint64(1 << uint64(square-1)), nil
}
