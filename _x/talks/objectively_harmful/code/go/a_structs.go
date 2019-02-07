package main

import (
	"fmt"
	"strings"
)

// BGN1 OMIT
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

func main() {
	a := human{"Alice"}
	b := wolf{3}
	username := "Dan"
	fmt.Println(a.greeting(username))
	fmt.Println(b.greeting(username))
}

// END1 OMIT
