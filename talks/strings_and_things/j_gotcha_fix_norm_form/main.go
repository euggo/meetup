package main

import (
	"fmt"
	"unicode"
)

func main() {
	//BGN1 OMIT
	in := "The quick bròwn 狐 jumped o⃐⃐⃐ver the lazy 犬"
	rs := []rune(in)

	n := len(rs)
	for i := 0; i < n/2; i++ {
		rs[i], rs[n-1-i] = rs[n-1-i], rs[i]
		if unicode.IsMark(rs[n-1-i]) {
			rs[n-1-i], rs[n-i] = rs[n-i], rs[n-1-i]
		}
	}
	for i := n / 2; i >= 0; i-- {
		if i > 0 && unicode.IsMark(rs[i-1]) {
			rs[i], rs[i-1] = rs[i-1], rs[i]
		}
	}

	fmt.Println(string(rs))
	//END1 OMIT
}
