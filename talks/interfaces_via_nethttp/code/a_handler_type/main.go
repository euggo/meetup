// BGN1 OMIT
package main

import (
	"fmt"
	"net/http"
)

type myScope struct{}

func (s *myScope) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a test.")
}

func main() {
	s := myScope{}
	http.ListenAndServe(":11142", &s)
}

// END1 OMIT
