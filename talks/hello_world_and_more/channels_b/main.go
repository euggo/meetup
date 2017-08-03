package main

import "fmt"

// START OMIT
func toy(c chan string) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan string)

	go toy(c)

	c <- "test"
	close(c)

	fmt.Println("completed")
}

// END OMIT
