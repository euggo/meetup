package main

import "fmt"

// START OMIT
func main() {
	s := make([]int, 10, 10) // builtin function
	s[0] = 42
	printEach(s)
}

// END OMIT

func printEach(s []int) {
	for _, v := range s {
		fmt.Println(v)
	}
}
