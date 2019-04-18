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

func addFirstLine(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "handlers.")
	})
}

func main() {
	h := requestHandler("nested")
	http.ListenAndServe(":11142", addFirstLine(h))
}

// END1 OMIT
