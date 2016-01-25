package main

import "fmt"

// START OMIT
func main() {
	c := make(chan string)

	c <- "test"

	fmt.Println("completed")
}

// END OMIT
