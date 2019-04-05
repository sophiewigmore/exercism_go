package diffsquares

// SquareOfSum : Returns the square of the sums of the natural numbers up to N
func SquareOfSum(n int) int {
	// formula for sum of first N natual nums: (number of terms * (first term+last term))/2
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares : Returns sum of the square of natural numbers up to N
func SumOfSquares(n int) int {
	// formula for the sum of squares
	return (n * (n + 1) * ((2 * n) + 1)) / 6
}

// Difference :  Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
