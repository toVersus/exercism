package luhn

import (
	"unicode"
)

// Valid determines whether or not it is valid per the Luhn formula.
func Valid(code string) bool {
	sum, size, mul := 0, 0, 1

	for i := len(code); i > 0; i-- {
		letter := rune(code[i-1])

		if unicode.IsNumber(letter) {
			digit := int(letter-'0') * mul
			if digit > 9 {
				digit -= 9
			}
			mul ^= 3
			sum += digit
			size++
		} else if !unicode.IsSpace(letter) {
			return false
		}
	}
	return sum%10 == 0 && size > 1
}
