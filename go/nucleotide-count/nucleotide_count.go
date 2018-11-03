package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// NewHistogram initializes an instance of Histogram.
func NewHistogram() Histogram {
	return Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
}

// DNA is a list of nucleotides.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := NewHistogram()
	for _, s := range d {
		if _, ok := h[s]; !ok {
			return nil, fmt.Errorf("strand with invalid nucleotides")
		}
		h[s]++
	}
	return h, nil
}
