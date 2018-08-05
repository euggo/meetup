package main

import (
	"net"
	"os"
)

// BGN1 OMIT
import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello HTTP")
	fmt.Println(r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	http.ListenAndServe(":8789", mux)
}

// END1 OMIT

func init() {
	_, err := net.Dial("tcp", ":8789")
	if err == nil {
		fmt.Println("line in use")
		os.Exit(42)
	}
}
