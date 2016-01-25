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

// START OMIT
func (g gopher) obfuscate() {
	g.name = secretName
}

func roleCall(g gopher) {
	fmt.Printf("%s is present.\n", g.name)
}

func main() {
	g := &gopher{
		name: "Novyem",
	}

	g.obfuscate()

	roleCall(g)
}

// END OMIT
