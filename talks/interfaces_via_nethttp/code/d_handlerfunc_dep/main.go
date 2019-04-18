package main

import (
	"fmt"
	"net/http"
)

// BGN1 OMIT
func requestHandler(msg string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	})
}

func main() {
	h := requestHandler("This is another example.")
	http.ListenAndServe(":11142", h)
}

// END1 OMIT
