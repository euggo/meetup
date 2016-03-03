package main

import (
	"fmt"
	"time"
)

// START OMIT
func toy() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("almost done!")
	time.Sleep(time.Millisecond * 100)
}

func main() {
	t := time.Now()
	toy()
	toy()
	fmt.Println("blocking duration:", time.Now().Sub(t))

	tt := time.Now()
	go toy()
	go toy()
	fmt.Println("concurrent duration:", time.Now().Sub(tt))
}

// END OMIT
