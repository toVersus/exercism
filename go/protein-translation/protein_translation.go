package protein

import "fmt"

var (
	STOP           = fmt.Errorf("terminating codon is found")
	ErrInvalidBase = fmt.Errorf("invalid RNA base sequence is found")
)

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
}

// FromCodon translates codon into protein.
func FromCodon(input string) (string, error) {
	if input == "UAA" || input == "UAG" || input == "UGA" {
		return "", STOP
	}

	if v, ok := codonToProtein[input]; ok {
		return v, nil
	}

	return "", ErrInvalidBase
}

// FromRNA translates RNA sequences into proteins.
func FromRNA(input string) ([]string, error) {
	var c string
	var rna []string
	for len(input) > 0 {
		c, input = input[0:3], input[3:]

		r, err := FromCodon(c)
		if err == ErrInvalidBase {
			return rna, err
		} else if err != nil {
			return rna, nil
		}
		rna = append(rna, r)
	}

	return rna, nil
}
