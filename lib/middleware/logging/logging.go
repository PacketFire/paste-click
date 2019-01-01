package logging

import (
	"log"
	"net/http"
)

// Middleware is a middleware handler that does request logging.
type Middleware struct {
	logger *log.Logger
}

// Serve implements the MiddleWareFunc type as defined in gorrila/mux.
func (lm *Middleware) Serve(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		lm.logger.Printf("%s %s", r.Method, r.URL.Path)

	})
}

// New constructs a new Middleware middleware handler
func New(l *log.Logger) *Middleware {
	return &Middleware{l}
}
