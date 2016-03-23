// START1 OMIT
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type timeJSON struct {
	Time time.Time `json:"time"` // struct tags pass info to callers looking for it
}

func localHandler(w http.ResponseWriter, r *http.Request) { // a valid http.HandlerFunc
	t := &timeJSON{time.Now()}

	b, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(500), 500) // use http.StatusInternalServerError
		return
	}

	w.Write(b)
}

// END1 OMIT

// START2 OMIT
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/local", localHandler)

	http.ListenAndServe(":29876", mux) // mux is a valid http.Handler

	// if a type has the method "ServeHTTP(http.ResponseWriter, *http.Request)",
	// it is an http.Handler
}

// END2 OMIT
