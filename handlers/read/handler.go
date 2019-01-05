package read

import (
	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler stores all required context for handing off requests to a
// storage driver.
type Handler struct {
	StorageDriver objectstore.ObjectStore
}

// New instantiates a new read Handler.
func New(store objectstore.ObjectStore) *Handler {
	store.Init()
	return &Handler{
		StorageDriver: store,
	}
}

// ServeHTTP implements the http.Handler interface for serving responses.
func (uh *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	oid := objectid.ObjectID(vars["objectid"])

	object, err := uh.StorageDriver.Read(oid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set(`content-type`, object.Metadata.Mimetype)
	w.Write(object.Data)
}
