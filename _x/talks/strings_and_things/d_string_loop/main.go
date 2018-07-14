package main

import "fmt"

func main() {
	//BGN1 OMIT
	s := "å være midt i smørøyet."
	// TODO: decide on a good-ish reason for needing to control iteration count
	for k, v := range s {
		fmt.Printf("%2d: %3d %s\n", k, v, string(v))
	}
	//END1 OMIT
}
