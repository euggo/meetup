package main

import (
	"fmt"
)

var (
	secretName = "Your Mom"
)

type gopher struct {
	name string
}

// START1 OMIT
func (g gopher) obfuscate() {
	g.name = secretName

	if g.name != "Novyem" {
		fmt.Println("field changed")
	}
}

// END1 OMIT

func roleCall(g *gopher) {
	fmt.Printf("%s is present.\n", g.name)
}

// START2 OMIT
func main() {
	g := &gopher{
		name: "Novyem",
	}

	g.obfuscate()

	roleCall(g)
}

// END2 OMIT
