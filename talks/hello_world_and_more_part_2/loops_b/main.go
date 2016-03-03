package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	for {
		fmt.Println("ding")
		if time.Now().Second()%3 == 0 {
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}

// END OMIT
