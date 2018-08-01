package caselib

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// ToWaveCase uppercases every other letter in a string.
func ToWaveCase(s string) string { //OMIT
	bs := make([]byte, 0, len(s))
	uc := true
	for i, r := range s {
		l := utf8.RuneLen(r)
		if !unicode.IsLetter(r) {
			bs = append(bs, s[i:i+l]...)
			continue
		}
		if !uc {
			bs = append(bs, strings.ToLower(s[i:i+l])...)
			uc = true
			continue
		}
		bs = append(bs, strings.ToUpper(s[i:i+l])...)
		uc = false
	}
	return string(bs)
} //OMIT
