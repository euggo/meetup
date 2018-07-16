package caselib

import (
	"strings"
)

// ToWaveCase uppercases every other letter in a string.
func ToWaveCase(s string) string {
	r, uc := "", true
	for i := 0; i < len(s); i++ {
		c := s[i : i+1]
		if c == " " {
			r += " "
			continue
		}
		if !uc {
			r += c
			uc = true
			continue
		}
		r += strings.ToUpper(c)
		uc = false
	}
	return r
}
