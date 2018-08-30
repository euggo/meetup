package main

import (
	"fmt"
	"time"
)

// BGN OMIT
func main() {
	c := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c <- "Eugene"
	}()

	fmt.Print("Hello, ")
	fmt.Println(<-c)
}

// END OMIT
