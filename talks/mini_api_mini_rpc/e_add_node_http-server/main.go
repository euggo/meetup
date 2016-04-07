// START1 OMIT
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	n, err := newNode(":19876") // HL
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	http.ListenAndServe(":29876", n) // HL
}

// END1 OMIT
