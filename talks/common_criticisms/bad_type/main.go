package main

import (
	"fmt"

	"github.com/euggo/meetup/talks/common_criticisms/bad_type/p1"
)

// START OMIT
func main() {
	h := p1.NewHidden()
	fmt.Println(h.Field)
	h.Field = 9000
	fmt.Println(h.Field)
}

// END OMIT
