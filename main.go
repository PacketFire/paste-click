package main

import (
	"log"
	"os"

	"github.com/PacketFire/paste-click/handlers/health"
	"github.com/PacketFire/paste-click/handlers/upload"

	"net/http"

	cs "github.com/PacketFire/paste-click/lib/config/service"
	"github.com/PacketFire/paste-click/lib/middleware/logging"
	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/PacketFire/paste-click/lib/objectstore/drivers/fs"
	"github.com/PacketFire/paste-click/lib/objectstore/drivers/mock"

	"github.com/gorilla/mux"
)

// store takes a string representing drivers and attempts to return a
// corresponding driver if there is no match, nil is returned.
func store(driverName string) objectstore.ObjectStore {
	switch driverName {
	case `fs`:
		return &fs.Store{}
	case `mock`:
		return &mock.Store{}
	default:
		return nil
	}
}

func main() {
	c := cs.New()

	// Router and health check handler
	mux := mux.NewRouter()
	mux.HandleFunc(`/healthcheck`, health.Handler).Methods("GET")

	s := store(c.StorageDriver)
	// Setup Upload handling
	uh := upload.New(s)
	mux.Handle(`/`, uh).Methods("POST")

	if c.Logging {
		// standard logger
		loggingMiddleware := logging.New(os.Stdout)
		mux.Use(loggingMiddleware.Serve)
	}

	log.Printf("Starting server on %s\n", c.Addr)
	if err := http.ListenAndServe(c.Addr, mux); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
