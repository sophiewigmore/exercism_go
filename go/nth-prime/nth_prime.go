package prime

// Nth takes in an int, n, and returns the nth prime number
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	// starting at the first prime number, keep going to next prime until at nth one
	for i := 2; ; i++ {
		if IsPrime(i) && n > 0 {
			n--
			if n == 0 {
				return i, true
			}
		}

	}
}

// IsPrime returns if an int, x, is prime and returns true if it
func IsPrime(x int) bool {
	if x < 2 {
		return false
	}
	// the only even prime number is 2
	if x%2 == 0 {
		return x == 2
	}
	// check if number is divisible by any number between 3 and x-1.
	for i := 3; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
