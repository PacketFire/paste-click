package upload

import (
	"io/ioutil"
	"net/http"
	"strings"

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
	contentType := r.Header.Get("Content-type")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read post body.", http.StatusInternalServerError)
	}

	if len(body) == 0 {
		http.Error(w, "Content body cannot be empty", http.StatusBadRequest)
	}

	scheme := r.Header.Get("X-Scheme")
	if scheme == "" {
		scheme = "http"
	}

	obj := objectstore.New(contentType, body)
	id := string(obj.Metadata.Object)

	err = uh.StorageDriver.Write(obj)
	if err != nil {
		http.Error(w, "Unable to save object", http.StatusInternalServerError)
	}

	fileURL := strings.Join([]string{scheme, "://", r.Host, r.RequestURI, id, "\n"}, "")
	w.Write([]byte(fileURL))
}
