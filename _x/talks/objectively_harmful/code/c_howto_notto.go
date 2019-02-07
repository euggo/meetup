package main

import "fmt"

// BGN1 OMIT
func concreteMeet(name string, h human) {
	fmt.Println(h.greeting(name))
}

func polymorphicMeet(name string, g greeter) {
	fmt.Println(g.greeting(name))
}

// END1 OMIT
