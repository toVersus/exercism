package romannumerals

import (
	"bytes"
	"fmt"
)

type roman struct {
	symbol  string
	numeral int
}

var romans = []roman{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// ToRomanNumeral converts from nomal numbers to Roman Numerals.
func ToRomanNumeral(num int) (string, error) {
	if num < 1 || num > 3000 {
		return "", fmt.Errorf("Number must be positive integer from 1 to 3000, got: %d", num)
	}

	var letters bytes.Buffer
	for _, r := range romans {
		for num != 0 && num >= r.numeral {
			letters.WriteString(r.symbol)
			num -= r.numeral
		}
	}

	return letters.String(), nil
}
