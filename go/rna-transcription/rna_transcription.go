package strand

import (
	"bytes"
	"unicode"
)

var transcriptor = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA transcripts DNA standard into RNA complement.
func ToRNA(dna string) string {
	var rna bytes.Buffer
	for _, d := range dna {
		d = unicode.ToUpper(d)
		if r, ok := transcriptor[d]; ok {
			rna.WriteRune(r)
		}
	}
	return rna.String()
}
