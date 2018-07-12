package main

import (
	"log"
	"os"
)

func main() {
	// START1 OMIT
	// the specified file does not exist
	f, _ := os.Open("file.txt")
	f.Chmod(777) // this kills the program
	// END1 OMIT
}

func other() {
	// START2 OMIT
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// use the opened file safely ...
	_ = f // END2 OMIT
}
