package timeout

import (
	"net/http"
	"time"
)

// BGN OMIT
// Handler should exceed the function configuration timeout and do nothing. OMIT
func Handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 3100)

	http.Error(w, "this should not be reached", http.StatusInternalServerError)
}

// END OMIT
