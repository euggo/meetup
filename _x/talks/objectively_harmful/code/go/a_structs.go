package main

import (
	"fmt"
	"strings"
)

// BGN1 OMIT
type human struct{ name string }

func newHuman(name string) *human {
	return &human{name: name}
}

func (h *human) greeting(name string) string {
	return fmt.Sprintf("Hello, %s. I'm %s.", name, h.name)
}

// END1 OMIT

// BGN2 OMIT
type wolf struct{ freq int }

func newWolf(freq int) *wolf {
	return &wolf{freq: freq}
}

func (w *wolf) greeting(string) string {
	msg := strings.Repeat("woof ", w.freq)
	return fmt.Sprintf("%s!", strings.TrimSpace(msg))
}

// END2 OMIT

// BGN3 OMIT
func main() {
	a := newHuman("Alice")
	b := newWolf(3)
	username := "Dan"
	fmt.Println(a.greeting(username))
	fmt.Println(b.greeting(username))
}

// END3 OMIT
