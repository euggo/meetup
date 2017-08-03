package main

import (
	"flag"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

// START OMIT
func newMux(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/hello", helloHandler)

	return mux
}

func main() {
	flag.Parse() // assuming flags have been declared in the application

	m := newMux(http.NewServeMux())
	http.ListenAndServe(":8789", m)
}

// END OMIT
