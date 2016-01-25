package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

// START OMIT
func multiplexer(f string) *http.ServeMux {
	// ignore f, it is used for example only
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	return mux
}

func main() {
	var profCPU = flag.String("profcpu", "", "Run CPU profile and write to file.")
	flag.Parse()

	mux := multiplexer(*profCPU)
	http.ListenAndServe(":8789", mux)
	log.Println("waiting for connections")
}

// END OMIT
