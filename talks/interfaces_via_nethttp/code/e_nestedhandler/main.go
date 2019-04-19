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

func addGreetings(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello.")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Goodbye.")
	})
}

func main() {
	h := requestHandler("This is a nested example.")
	http.ListenAndServe(":11142", addGreetings(h))
}

// END1 OMIT
