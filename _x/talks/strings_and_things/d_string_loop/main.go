package main

import "fmt"

func main() {
	//BGN1 OMIT
	s := "å være"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %T - %v\n", i, s[i], s[i])
	}
	//END1 OMIT
}
