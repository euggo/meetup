package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func toy(c chan string, wg *sync.WaitGroup) {
	fmt.Println("toy is running and waiting for test message")
	fmt.Println(<-c)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	c := make(chan string)

	wg.Add(1)
	go toy(c, wg)

	time.Sleep(time.Second)
	c <- "test"

	wg.Wait()
	fmt.Println("completed")
}

// END OMIT
