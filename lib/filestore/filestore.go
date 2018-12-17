package filestore

// FileStore provides the interface for persisting files posted to the api.
type FileStore interface {
	Init() error         // Initializes a Filestore
	Write(*Object) error // Attempts to write an Object to the filestore.
	Read(string) *Object // Takes an object identifier as an argument and retrieves it from the store.
	Close() error        // Safely shutsdown a filestore.
}
