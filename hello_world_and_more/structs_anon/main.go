package main

import "fmt"

func main() {
	// START OMIT
	a := struct {
		name string
		age  int
	}{
		name: "Hombre",
		age:  42,
	}
	// END OMIT

	fmt.Println(a)
}
