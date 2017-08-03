package main

import "fmt"

// START OMIT
func main() {
	f()
	fmt.Println("returned normally from f")
}

func f() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered in f (defer):", r)
		}
	}()

	fmt.Println("calling g")
	g()

	fmt.Println("returned normally from g")
}

func g() {
	panic("who wrote this lib?")
}

// END OMIT
