package protein

import (
	"errors"
)

// FromCodon translates a codon string into the amino acid string
func FromCodon(codon string) (string, error) {
	mapping := getMapping()
	aminoAcid, ok := mapping[codon]

	if !ok {
		return "", errors.New("Invalid Base")
	} else if aminoAcid == "STOP" {
		return "", errors.New("Stop Codon")
	}
	return aminoAcid, nil
}

// FromRNA translates a string of RNA codons into the polypeptide string
func FromRNA(rna string) ([]string, error) {
	polypeptide := []string{}
	err := *new(error)

	for i := 0; i < len(rna); i = i + 3 {
		codon := rna[i : i+3]
		aminoAcid, newErr := FromCodon(codon)

		if aminoAcid == "" && newErr.Error() == "Invalid Base" {
			return polypeptide, newErr
		}

		if aminoAcid == "" && newErr.Error() == "Stop Codon" {
			return polypeptide, err
		}

		polypeptide = append(polypeptide, aminoAcid)

		if newErr != err {
			err = newErr
		}
	}
	return polypeptide, err
}

// getMapping returns the map of codon strings to amino acid strings
func getMapping() map[string]string {
	mapping := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine", "UUC": "Phenylalanine",
		"UUA": "Leucine", "UUG": "Leucine",
		"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
		"UAU": "Tyrosine", "UAC": "Tyrosine",
		"UGU": "Cysteine", "UGC": "Cysteine", "UGG": "Tryptophan",
		"UAA": "STOP", "UAG": "STOP", "UGA": "STOP"}
	return mapping
}
