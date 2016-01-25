package main

import "fmt"

// START1 OMIT
type greeter interface {
	greeting() string
}

type person struct {
	message string
}

func (p *person) greeting() string {
	return p.message
}

type dog struct {
	bark string
}

func (d *dog) greeting() string {
	return d.bark
}

// END1 OMIT

// START2 OMIT
func speak(g greeter) {
	fmt.Println(g.greeting())
}

func main() {
	p := &person{
		message: "Hello World",
	}

	d := &dog{
		bark: "Woof Woof",
	}

	speak(p)
	speak(d)
}

// END2 OMIT
