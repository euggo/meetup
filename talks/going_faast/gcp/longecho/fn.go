package longecho

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Handler returns a message that sometimes exceeds a deadline. OMIT
func Handler(w http.ResponseWriter, r *http.Request) {
	var toc <-chan time.Time
	if dl, ok := r.Context().Deadline(); ok {
		toc = time.After(time.Until(dl.Add(-10 * time.Millisecond)))
	}

	req, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close() // nolint
	if err != nil {
		httpError(w, http.StatusInternalServerError)
		return
	}

	select {
	case <-toc:
		httpError(w, http.StatusRequestTimeout)
	case resp := <-randomlySlowMessage(string(req)):
		if _, err := fmt.Fprintln(w, resp); err != nil {
			httpError(w, http.StatusInternalServerError)
		}
	}
}

// END OMIT

func httpError(w http.ResponseWriter, code int) {
	st := http.StatusText(code)
	if code == http.StatusRequestTimeout {
		st = "Sometimes, if you wait too long, it's too late. — Sarah Strohmeyer"
	}

	http.Error(w, st, code)
}

func randomlySlowMessage(in string) <-chan string {
	c := make(chan string)

	go func() {
		if time.Now().UnixNano()%2 == 0 {
			time.Sleep(3001 * time.Millisecond) // should be just above default lambda timeout
		}

		c <- fmt.Sprintf("Your input: %s", in)
	}()

	return c
}
