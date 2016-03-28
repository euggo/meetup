// START1 OMIT
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type timeJSON struct {
	Time time.Time `json:"time"` // struct tags pass info to callers looking for it // HL
}

func localHandler(w http.ResponseWriter, r *http.Request) { // a valid http.HandlerFunc
	t := &timeJSON{time.Now()} // HL

	b, err := json.Marshal(t) // HL
	if err != nil {
		http.Error(w, http.StatusText(500), 500) // use http.StatusInternalServerError
		return
	}

	w.Write(b) // HL
}

// END1 OMIT

// START2 OMIT
func main() {
	mux := http.NewServeMux() // HL

	mux.HandleFunc("/api/local", localHandler) // HL

	http.ListenAndServe(":29876", mux) // mux is a valid http.Handler // HL

	// if a type has the method "ServeHTTP(http.ResponseWriter, *http.Request)",
	// it is an http.Handler
}

// END2 OMIT
