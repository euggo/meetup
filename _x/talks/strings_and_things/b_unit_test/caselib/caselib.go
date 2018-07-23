package caselib

import (
	"strings"
)

// ToWaveCase uppercases every other letter in a string.
func ToWaveCase(s string) string {
	out, uc := "", true
	for i := 0; i < len(s); i++ {
		c := s[i : i+1]
		if c == " " {
			out += c
			continue
		}
		if !uc {
			out += c
			uc = true
			continue
		}
		out += strings.ToUpper(c)
		uc = false
	}
	return out
}
