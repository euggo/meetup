package main

import "fmt"

func main() {
	//BGN1 OMIT
	s := "å være midt i smørøyet."

	for i := 0; i < len(s); i++ {
		fmt.Printf("%2d: %T - %v\n", i, s[i:i+1], s[i:i+1])
	}
	//END1 OMIT
}
