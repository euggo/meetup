package main

import (
	"encoding/json" // OMIT
	"net/http"      // OMIT
	"time"          // OMIT
	// "encoding/json", "net/http", "time"
)

// START1 OMIT
type timeJSON struct {
	Time time.Time `json:"time"`
}

func (n *node) localHandler(w http.ResponseWriter, r *http.Request) {
	t := &timeJSON{time.Now()}

	b, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(500), 500) // use http.StatusInternalServerError
		return
	}

	w.Write(b)
}

func (n *node) remoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("remote"))
}

func (n *node) statsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("stats"))
}

// END1 OMIT
