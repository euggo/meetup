package main

import (
	"fmt"
	"strconv"
)

func content(num string) string {
	// BGN OMIT
	n, err := strconv.Atoi(num)
	if err != nil {
		return "The value is not a valid integer."
	}

	return "The integer is " + strconv.Itoa(n) + "."
	// END OMIT
}

func main() {
	fmt.Println(content(".333"))
}
