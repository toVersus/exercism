package luhn

import (
	"unicode"
)

// Valid determines whether or not it is valid per the Luhn formula.
func Valid(code string) bool {
	sum, size := 0, 0

	for i := len(code); i > 0; i-- {
		letter := rune(code[i-1])

		if unicode.IsNumber(letter) {
			digit := int(letter - '0')
			if size%2 == 1 {
				digit *= 2
			}

			if digit > 9 {
				digit -= 9
			}

			sum += digit
			size++
			continue
		}

		if !unicode.IsSpace(letter) {
			return false
		}
	}
	return sum%10 == 0 && size > 1
}
