package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type timeJSON struct {
	Time time.Time `json:"time"`
}

func localHandler(w http.ResponseWriter, r *http.Request) {
	t := &timeJSON{time.Now()}

	b, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(500), 500) // use http.StatusInternalServerError
		return
	}

	w.Write([]byte(b))
}

func remoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("remote"))
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("stats"))
}

// START1 OMIT
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/local", localHandler)
	mux.HandleFunc("/api/remote/", remoteHandler)
	mux.HandleFunc("/api/stats", statsHandler)

	http.ListenAndServe(":29876", mux)
}

// END1 OMIT
