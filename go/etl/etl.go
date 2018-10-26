package etl

import (
	"strings"
)

// Transform computes the scrabble score for given word.
func Transform(old map[int][]string) map[string]int {
	new := make(map[string]int)
	for score, letters := range old {
		for _, l := range letters {
			new[strings.ToLower(l)] = score
		}
	}
	return new
}
