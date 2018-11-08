package main

import (
	"fmt"
	"strconv"
)

func content(num string) string {
	// BGN OMIT
	n, err := strconv.Atoi(num)
	if err != nil {
		if err, ok := err.(*strconv.NumError); ok {
			switch err.Err {
			case strconv.ErrSyntax:
				return "The value is not an integer."
			case strconv.ErrRange:
				return "The value is out of range."
			}
		}
		return "The value is wack, yo."
	}
	// END OMIT

	return "The integer is " + strconv.Itoa(n) + "."
}

func main() {
	fmt.Println(content(".333"))
}
