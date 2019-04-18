package main

import (
	"fmt"
	"net/http"
)

// BGN1 OMIT
func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is an example.")
}

func main() {
	h := http.HandlerFunc(handleRequest)
	http.ListenAndServe(":11142", h)
}

// END1 OMIT
