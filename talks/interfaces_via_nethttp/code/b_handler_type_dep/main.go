package main

import (
	"fmt"
	"net/http"
)

// BGN1 OMIT
type myScope struct {
	msg string
}

func (s *myScope) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s.msg)
}

func main() {
	s := myScope{
		msg: "This is another test.",
	}
	http.ListenAndServe(":11142", &s)
}

// END1 OMIT
