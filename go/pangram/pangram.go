package pangram

import (
	"strings"
)

// IsPangram determines if a sentence is a pangram.
func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for c := 'a'; c < 'z'; c++ {
		if !strings.ContainsRune(input, c) {
			return false
		}
	}
	return true
}
