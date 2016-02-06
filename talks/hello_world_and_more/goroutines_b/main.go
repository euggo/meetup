package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func toy(wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("almost done!")
	time.Sleep(time.Millisecond * 100)

	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	t := time.Now()
	go toy(wg)
	go toy(wg)
	fmt.Println("concurrent duration A:", time.Now().Sub(t))

	wg.Wait()
	fmt.Println("concurrent duration B:", time.Now().Sub(t))
}

// END OMIT
