package logging

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
)

// Middleware is a middleware handler that does request logging.
type Middleware struct {
	logDestination io.Writer
}

// New constructs a new Middleware middleware handler
func New(l io.Writer) *Middleware {
	return &Middleware{l}
}

// Serve implements the MiddleWareFunc type as defined in gorrila/mux.
func (lm *Middleware) Serve(next http.Handler) http.Handler {
	return handlers.LoggingHandler(lm.logDestination, next)
}
