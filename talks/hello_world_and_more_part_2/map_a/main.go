package main

import "fmt"

// START OMIT
type info struct {
	age   int
	catch string
}

func main() {
	m := map[string]info{
		"Rick":    info{80, "Wubba lubba dub-dub!"},
		"Morty":   info{15, "Oh, man! Oh, oh geez! Oh..."},
		"Mr. Pbh": info{0, "Ooh-wee!"},
	}

	x := make(map[string]bool)
	x["Rick"] = true
	x["Morty"] = true
	x["Mr. Pbh"] = false

	for k, v := range m {
		if x[k] {
			fmt.Println(k, v)
		}
	}
}

// END OMIT
