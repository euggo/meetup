package main

import "fmt"

type gopher struct {
	name string
}

// START OMIT
func main() {
	gi0 := gopher{
		name: "Innie",
	}

	var gi1 gopher
	gi1.name = "Outie"

	gp0 := &gopher{
		name: "Pointy",
	}

	gp1 := new(gopher) // type *gopher
	gp1.name = "Poindexter"

	fmt.Printf("%T, %T, %T, %T\n", gi0, gi1, gp0, gp1)
	fmt.Printf("%v, %v, %v, %v\n", gi0, gi1, gp0, gp1)
}

// END OMIT
