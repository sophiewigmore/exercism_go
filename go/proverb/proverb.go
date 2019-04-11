package proverb

// Proverb takes in a string array of words and returns a proverb
// "1" "2" "3" => "1,2", "2,3", "1"
// "1" => 1
func Proverb(rhyme []string) []string {
	var proverb []string

	if len(rhyme) < 1 {
		return proverb
	}

	for i, word := range rhyme {
		if i < len(rhyme)-1 {
			part := "For want of a " + word + " the " + rhyme[i+1] + " was lost."
			proverb = append(proverb, part)
		} else {
			finalPart := "And all for the want of a " + rhyme[0] + "."
			proverb = append(proverb, finalPart)
		}
	}
	return proverb
}
