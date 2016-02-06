package main

import "os"

// START OMIT
func main() {
	f, _ := os.Open("file.txt")
	f.Chmod(777)
}

// END OMIT
