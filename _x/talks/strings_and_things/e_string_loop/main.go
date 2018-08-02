package main

import "fmt"

func main() {
	//BGN1 OMIT
	s := "å være"

	for i, r := range s {
		fmt.Printf("%2d: %T - %3d / string - %s\n", i, r, r, string(r))
	}
	//END1 OMIT
}
