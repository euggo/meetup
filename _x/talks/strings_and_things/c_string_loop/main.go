package main

import "fmt"

func main() {
	//BGN1 OMIT
	s := "å være midt i smørøyet."

	for i := 0; i < len(s); i++ {
		fmt.Printf("%2d: %3d %s\n", i, s[i], string(s[i]))
	}
	//END1 OMIT
}
