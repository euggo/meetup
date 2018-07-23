package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	//BGN1 OMIT
	in := "The quick bròwn 狐 jumped over the lazy 犬"
	rs := []rune(norm.NFC.String(in))

	n := len(rs)
	for i := 0; i < n/2; i++ {
		rs[i], rs[n-1-i] = rs[n-1-i], rs[i]
	}

	fmt.Println(string(rs))
	//END1 OMIT
}
