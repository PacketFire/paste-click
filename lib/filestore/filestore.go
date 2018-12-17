package filestore

// FileStore provides the interface for persisting files posted to the api.
type FileStore interface {
	Init() error                    // Initializes a Filestore
	Read(ObjectID) (*Object, error) // Takes an object identifier as an argument and retrieves it from the store.
	Write(*Object) error            // Attempts to write an Object to the filestore.
	Close() error                   // Safely shutsdown a filestore.
}
