package main

import (
	"fmt"
	"time"
)

// START OMIT

// taken from https://tour.golang.org/concurrency/6

func main() {
	tick := time.Tick(200 * time.Millisecond)
	boom := time.After(600 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// END OMIT
