package health

import (
	"fmt"
	"net/http"
)

// Handler takes a GET request and returns a 200 response to simulate a
// health check.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "Ok"}`)
}
