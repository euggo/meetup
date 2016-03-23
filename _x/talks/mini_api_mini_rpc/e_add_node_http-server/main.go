// START1 OMIT
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	n, err := newNode(":19876")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	http.ListenAndServe(":29876", n)
}

// END1 OMIT
