package upload

import (
	"net/http"

	"github.com/PacketFire/paste-click/lib/objectstore"
)

// Handler stores all required context for handing off upload requests to a
// storage driver.
type Handler struct {
	StorageDriver objectstore.ObjectStore
}

// New instantiates a new upload Handler.
func New(store objectstore.ObjectStore) *Handler {
	store.Init()
	return &Handler{
		StorageDriver: store,
	}
}

// ServeHTTP implements the http.Handler interface for serving responses.
func (uh *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//		write := savePost(w, r)
	//		w.Write(write)
	return
}
