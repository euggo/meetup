package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func toy(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 100)
	fmt.Println("almost done!")
	time.Sleep(time.Millisecond * 100)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	t := time.Now()
	go toy(wg)
	go toy(wg)
	fmt.Println("concurrent funcs call duration:", time.Now().Sub(t))

	wg.Wait()
	fmt.Println("concurrent funcs completion duration:", time.Now().Sub(t))
}

// END OMIT
