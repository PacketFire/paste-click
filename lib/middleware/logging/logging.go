package logging

import (
	"log"
	"net/http"
)

// Middleware is a middleware handler that does request
type Middleware struct {
	logger  *log.Logger
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (lm *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lm.handler.ServeHTTP(w, r)
	lm.logger.Printf("%s %s", r.Method, r.URL.Path)
}

// New constructs a new Middleware middleware handler
func New(l *log.Logger, handlerToWrap http.Handler) *Middleware {
	return &Middleware{l, handlerToWrap}
}
