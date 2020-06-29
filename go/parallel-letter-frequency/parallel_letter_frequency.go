package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency
func ConcurrentFrequency(words []string) FreqMap {
	freqMap := make(chan FreqMap)

	composite := FreqMap{}

	for _, s := range words {
		go func(s string) {
			freqMap <- Frequency(s)
		}(s)

		for k, v := range <-freqMap {
			composite[k] += v
		}
	}

	return composite
}
