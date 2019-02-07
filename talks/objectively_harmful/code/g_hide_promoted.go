package main

import (
	"fmt"
)

type greeter interface {
	greeting(string) string
}

func meet(name string, gs ...greeter) {
	for _, g := range gs {
		fmt.Println(g.greeting(name))
	}
}

type messager interface {
	message() string
}

func talk(ms ...messager) {
	for _, m := range ms {
		fmt.Println(m.message())
	}
}

type human struct {
	name string
}

func (h *human) greeting(name string) string {
	return fmt.Sprintf("Hello, %s. My name is %s.", name, h.name)
}

func (h *human) message() string { return "Nice to meet you!" }

// BGN1 OMIT
type wolf struct {
	human
	message struct{} // can be any type; empty struct has zero cost
}

func main() {
	a := &human{"Alice"}
	b := &wolf{human: human{"Bob"}}
	meet("Dan", a, b)
	talk(a, b)
}

// END1 OMIT
