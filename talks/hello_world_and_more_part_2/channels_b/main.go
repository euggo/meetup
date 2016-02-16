package main

import "fmt"

// START OMIT
func toy(mc chan string) {
	fmt.Println(<-mc)
}

func main() {
	c := make(chan string)
	go toy(c)

	c <- "test"

	fmt.Println("completed")
}

// END OMIT
