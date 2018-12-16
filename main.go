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
	var router http.Handler
	c := cs.New()

	mux := mux.NewRouter()
	mux.HandleFunc(`/healthcheck`, healthHandler).Methods("GET")
	router = mux

	if c.Logging {
		// standard logger
		sl := log.New(os.Stderr, "", log.LstdFlags)
		router = logging.New(sl, router)
	}

	log.Printf("Starting server on %s\n", c.Addr)
	if err := http.ListenAndServe(c.Addr, router); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
