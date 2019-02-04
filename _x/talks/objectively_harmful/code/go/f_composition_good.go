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

// BGN1 OMIT
type messager interface {
	message() string
}

func talk(ms ...messager) {
	for _, m := range ms {
		fmt.Println(m.message())
	}
}

// END1 OMIT

type human struct{ name string }

func newHuman(name string) *human {
	return &human{name: name}
}

func (h *human) greeting(name string) string {
	return fmt.Sprintf("Hello, %s. I'm %s.", name, h.name)
}

// BGN2 OMIT
func (h *human) message() string {
	return "Nice to meet you."
}

// END2 OMIT

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

func (w *werewolf) greeting(name string) string {
	return w.wolf.greeting(name) + " " + w.human.greeting(name)
}

// BGN3 OMIT
func main() {
	a := newHuman("Alice")
	b := newWolf(3)
	c := newWerewolf(newHuman("Carlos"), newWolf(1))
	meet("Dan", a, b, c)
	// cannot use b (type *wolf) as type messager in argument to talk:
	// *wolf does not implement messager
	talk(a, c)
}

// END3 OMIT
