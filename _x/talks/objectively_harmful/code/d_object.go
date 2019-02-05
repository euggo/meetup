package main

import "fmt"

// BGN1 OMIT
// structure declares a simple data structure
type structure struct {
	field string
}

// method prints the field contained in the structure.
func (s *structure) method() {
	fmt.Println(s.field)
}

func example() {
	s := structure{field: "data"} // s is an instance of "structure"
	s.method()                    // s is an object
}

// END1 OMIT

func main() {
	example()
}
