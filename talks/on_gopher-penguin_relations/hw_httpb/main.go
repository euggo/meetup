package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { // HLhandler
	fmt.Fprintln(w, "Hello, Eugene") // HLhandler
} // HLhandler

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", helloHandler) // our previous code used this instead
	mux.Handle("/", http.HandlerFunc(helloHandler)) // HLhandler

	http.ListenAndServe(":8789", mux)
}
