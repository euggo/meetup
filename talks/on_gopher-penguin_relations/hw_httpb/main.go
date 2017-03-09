package main

import (
	"net"
	"os"
	"time"
)

// START1 OMIT
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

// END1 OMIT
var redundant bool

func init() {
	Init()
}

// Init ...
func Init() {
	_, err := net.Dial("tcp", ":8789")
	if err != nil {
		switch redundant {
		case true:
			os.Exit(42)
		case false:
			return
		}
	}
	redundant = true

	time.Sleep(333 * time.Millisecond)
	Init()
}
