package main

import "fmt"

func main() {
	//BGN1 OMIT
	s := "å være midt i smørøyet."

	for i := 0; i < len(s); i++ {
		fmt.Printf("%2d: %v\n", i, s[i])
	}
	//END1 OMIT
}
