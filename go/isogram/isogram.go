package isogram

import (
	"strings"
)

// IsIsogram returns true if a word or phrase is an isogram.
func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	letterSet := make(map[rune]struct{})
	for _, r := range input {
		if isWhitespace(r) || isHyphen(r) {
			continue
		}

		if _, ok := letterSet[r]; !ok {
			letterSet[r] = struct{}{}
			continue
		}

		return false
	}

	return true
}

func isWhitespace(r rune) bool {
	if r == ' ' {
		return true
	}
	return false
}

func isHyphen(r rune) bool {
	if r == '-' {
		return true
	}
	return false
}
