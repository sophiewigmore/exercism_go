package strand

import "strings"

// ToRNA converts a DNA string into its RNA transcription
// See README for details
func ToRNA(dna string) string {
	transcription := ""
	for i := 0; i < len(dna); i++ {
		if strings.ToUpper(dna[i:i+1]) == "A" {
			transcription += "U"
		} else if strings.ToUpper(dna[i:i+1]) == "T" {
			transcription += "A"
		} else if strings.ToUpper(dna[i:i+1]) == "C" {
			transcription += "G"
		} else if strings.ToUpper(dna[i:i+1]) == "G" {
			transcription += "C"
		}
	}
	return transcription
}
