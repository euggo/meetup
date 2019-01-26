package fn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type request struct {
	Name string `json:"name"`
}

type response struct {
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

// Handler prints a hello message. OMIT
func Handler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	defer r.Body.Close() // nolint

	req := request{}
	if err := d.Decode(&req); err != nil {
		s := http.StatusInternalServerError
		if _, ok := err.(*json.SyntaxError); ok {
			s = http.StatusBadRequest
		}
		http.Error(w, http.StatusText(s), s)
		return
	}

	msg := fmt.Sprintf("Hello from GCP, %s. Have a wonderful day.", req.Name)
	resp := response{
		Msg:  msg,
		Time: time.Now().String(),
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, http.StatusText(srvrError), srvrError)
	}
}

// END OMIT

var (
	srvrError = http.StatusInternalServerError
)
