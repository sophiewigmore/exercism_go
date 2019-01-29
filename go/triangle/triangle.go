package triangle

import "math"

// Kind stores the type of triangle.
type Kind int

const (
	// NaT is not a triangle
	NaT Kind = 0
	// Equ = equilateral
	Equ Kind = 1
	// Iso = isosceles
	Iso Kind = 2
	// Sca = scalene
	Sca Kind = 3
)

// KindFromSides returns the type of triangle based off of the sides
func KindFromSides(a, b, c float64) Kind {
	if SidesAreBadArgs(a, b, c) || RuleOfInequality(a, b, c) {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

// SidesAreBadArgs is a method that returns if any of the sides are of length 0 or less
func SidesAreBadArgs(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a+b+c) || math.IsInf(a+b+c, 1)
}

// RuleOfInequality is a method that returns if the sides of a triangle disobey triangle equality rules.
// the sum of the lengths of any two sides must be
// greater than or equal to the length of the remaining side.
func RuleOfInequality(a, b, c float64) bool {
	return a+b < c || a+c < b || b+c < a
}
