package upload

import (
	"net/http"
)

// Handler stores all required context for handing off upload requests to a
// storage driver.
type Handler struct{}

// New instantiates a new uplad Handler.
func New() *Handler {
	return &Handler{}
}

// ServeHTTP implements the http.Handler interface for serving responses.
func (uh *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//		write := savePost(w, r)
	//		w.Write(write)
	return
}
