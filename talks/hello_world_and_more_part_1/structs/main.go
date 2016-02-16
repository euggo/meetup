package main

import (
	"fmt"
)

// START OMIT
type gopher struct {
	name string
}

func roleCall(g gopher) {
	fmt.Printf("%s is present.\n", g.name)
}

func main() {
	g := gopher{
		name: "Novyem",
	}

	roleCall(g)
}

// END OMIT
