package caselib

import (
	"strings"
	"unicode"
)

// ToWaveCase uppercases every other letter in a string.
func ToWaveCase(s string) string { //OMIT
	n, p := unicode.ToUpper, unicode.ToLower
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			r = n(r)
			n, p = p, n
		}
		return r
	}, s)
} //OMIT
