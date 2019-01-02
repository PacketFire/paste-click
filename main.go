package main

import (
	"log"
	"os"

	"github.com/PacketFire/paste-click/handlers/health"
	"github.com/PacketFire/paste-click/handlers/upload"

	"net/http"

	cs "github.com/PacketFire/paste-click/lib/config/service"
	"github.com/PacketFire/paste-click/lib/middleware/logging"

	"github.com/gorilla/mux"
)

func main() {
	c := cs.New()

	// Router and health check handler
	mux := mux.NewRouter()
	mux.HandleFunc(`/healthcheck`, health.Handler).Methods("GET")

	// Setup Upload handling
	uh := upload.New()
	mux.Handle(`/`, uh).Methods("POST")

	if c.Logging {
		// standard logger
		sl := log.New(os.Stderr, "", log.LstdFlags)
		loggingMiddleware := logging.New(sl)
		mux.Use(loggingMiddleware.Serve)
	}

	log.Printf("Starting server on %s\n", c.Addr)
	if err := http.ListenAndServe(c.Addr, mux); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
