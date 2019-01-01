package objectstore

import "github.com/PacketFire/paste-click/lib/objectstore/objectid"

// ObjectStore provides the interface for persisting files posted to the api.
type ObjectStore interface {
	Init() error                             // Initializes a ObjectStore
	Read(objectid.ObjectID) (*Object, error) // Takes an object identifier as an argument and retrieves it from the store.
	Write(*Object) error                     // Attempts to write an Object to the object store.
	Close() error                            // Safely shutsdown a object store.
}
