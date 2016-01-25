package main

import (
	"fmt"
	"time"
)

// START OMIT
func nums(t time.Time) (int, int, int) {
	return t.Hour(), t.Minute(), t.Second()
}

func numsDub(t time.Time) (int, int, int) {
	return t.Hour() * 2, t.Minute() * 2, t.Second() * 2
}

func main() {
	h, m, s := nums(time.Now())
	fmt.Println(h, m, s)

	h, _, s = numsDub(time.Now())
	fmt.Println(h, m, s)
}

// END OMIT
