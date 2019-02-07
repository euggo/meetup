package main

import (
	"fmt"
	"strings"
)

type greeter interface {
	greeting(name string) string
}

func meet(name string, gs ...greeter) {
	for _, g := range gs {
		fmt.Println(g.greeting(name))
	}
}

type human struct {
	name string
}

func (h human) greeting(name string) string {
	return fmt.Sprintf("Hello, %s. I'm %s.", name, h.name)
}

type wolf struct {
	freq int
}

func (w wolf) greeting(string) string {
	msg := strings.Repeat("woof ", w.freq)
	return fmt.Sprintf("%s!", strings.TrimSpace(msg))
}

type werewolf struct {
	human
	wolf
}

func newWerewolf(name string, freq int) *werewolf {
	return &werewolf{
		human: human{name},
		wolf:  wolf{freq},
	}
}

// BGN1 OMIT
func (w *werewolf) greeting(name string) string {
	return w.wolf.greeting(name) + " " + w.human.greeting(name)
}

func main() {
	a := human{"Alice"}
	b := wolf{3}
	c := newWerewolf("Carlos", 1)
	meet("Dan", a, b, c)
}

// END1 OMIT
