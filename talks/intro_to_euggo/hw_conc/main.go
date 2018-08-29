package main

import (
	"fmt"
	"time"
)

func main() {
	// BGN OMIT
	c := make(chan string, 2)

	go func() {
		time.Sleep(time.Second)
		c <- "World"
		close(c)
	}()

	c <- "Hello"

	for s := range c {
		fmt.Println(s)
	}
	// END OMIT
}
