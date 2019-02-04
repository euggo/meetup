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

type human struct{ name string }

func newHuman(name string) *human {
	return &human{name: name}
}

func (h *human) greeting(name string) string {
	return fmt.Sprintf("Hello, %s. I'm %s.", name, h.name)
}

type wolf struct{ freq int }

func newWolf(freq int) *wolf {
	return &wolf{freq: freq}
}

func (w *wolf) greeting(string) string {
	msg := strings.Repeat("woof ", w.freq)
	return fmt.Sprintf("%s!", strings.TrimSpace(msg))
}

type werewolf struct {
	human
	wolf
}

func newWerewolf(h *human, w *wolf) *werewolf {
	return &werewolf{human: *h, wolf: *w}
}

// BGN1 OMIT
func (w *werewolf) greeting(name string) string {
	return w.wolf.greeting(name) + " " + w.human.greeting(name)
}

// END1 OMIT

func main() {
	a := newHuman("Alice")
	b := newWolf(3)
	c := newWerewolf(newHuman("Carlos"), newWolf(1))
	meet("Dan", a, b, c)
}
