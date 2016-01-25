package main

import "fmt"

// START OMIT
func main() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

// END OMIT
