package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { // HLhandler
	fmt.Fprintln(w, "Hello, Eugene") // HLhandler
} // HLhandler

func main() { // HLmain
	mux := http.NewServeMux() // HLmain

	mux.HandleFunc("/", helloHandler) // HLmain

	http.ListenAndServe(":8789", mux) // HLmain
} // HLmain
