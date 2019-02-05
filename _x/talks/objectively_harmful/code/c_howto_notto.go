package main

import "fmt"

// BGN1 OMIT
// concreteMeet is bound to "human" (exposing unnecessary fields/methods)
func concreteMeet(name string, h *human) {
	fmt.Println(h.greeting(name))
}

// polymorphicMeet takes any "greeter" (requiring abstractly by need)
func polymorphicMeet(name string, g greeter) {
	fmt.Println(g.greeting(name))
}

// END1 OMIT
