package main

import "fmt"

// START OMIT
func main() {
	fmt.Println(0)
	defer fmt.Println(1)
	fmt.Println(2)
	defer fmt.Println(3)
}

// END OMIT
