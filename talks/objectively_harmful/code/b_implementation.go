package main

import "fmt"

// BGN1 OMIT
type human struct {
	name string
}

func (h human) greeting(name string) string {
	return fmt.Sprintf("Hello, %s. I'm %s.", name, h.name)
}

// END1 OMIT
