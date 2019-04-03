package prime

import "testing"

var testCases = []struct {
	x        int
	expected bool
}{
	{55, false},
	{5, true},
	{103, true},
	{10000, false},
	{1, false},
}

// TestIsPrime test
func TestIsPrime(t *testing.T) {
	for _, test := range testCases {
		if observed := IsPrime(test.x); observed != test.expected {
			t.Fatalf("ShareWith(%d) = %v, want %v", test.x, observed, test.expected)
		}
	}
	t.Logf("Pass")
}
