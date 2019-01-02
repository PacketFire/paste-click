package upload

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/rakyll/magicmime"
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

	// lookup content-type of the data
	contentType, err := getMimeString(body)
	if err != nil {
		http.Error(w, "Couldn't assertain mimetype from post body", http.StatusBadRequest)
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

func getMimeString(data []byte) (string, error) {
	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_ERROR); err != nil {
		log.Fatal(err)
	}
	defer magicmime.Close()

	mimetype, err := magicmime.TypeByBuffer(data)
	if err != nil {
		log.Fatalf("error occured during type lookup: %v", err)
	}
	return mimetype, nil
}
