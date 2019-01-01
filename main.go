package main

import (
	"log"
	"os"

	"net/http"

	cs "github.com/PacketFire/paste-click/lib/config/service"
	"github.com/PacketFire/paste-click/lib/middleware/logging"

	"github.com/gorilla/mux"
)

func main() {
	c := cs.New()

	mux := mux.NewRouter()
	mux.HandleFunc(`/healthcheck`, healthHandler).Methods("GET")
	mux.HandleFunc(`/`, uploadHandler).Methods("POST")

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
