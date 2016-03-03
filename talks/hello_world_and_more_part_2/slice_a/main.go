package main

import "fmt"

// START OMIT
func main() {
	s := []string{"y", "ahora", "mas"}
	s = append(s, "y")
	s = append(s, "mas", "para", "siempre")

	fmt.Println(s)
	fmt.Println(s[4:])
	for _, v := range s[2:5] {
		fmt.Println(v)
	}
}

// END OMIT
