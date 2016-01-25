package main

import (
	"fmt"
	"time"
)

// START OMIT
func toy(mc chan string) {
	fmt.Println("toy is running and waiting for test message")
	fmt.Println(<-mc)
}

func main() {
	c := make(chan string)
	go toy(c)

	time.Sleep(time.Second)
	c <- "test"

	fmt.Println("completed")
}

// END OMIT
