package local

import (
	fs "github.com/PacketFire/paste-click/lib/filestore"
	"github.com/caarlos0/env"
)

// Store implements the FileStore interface to store Objects on the local
// filesystem.
type Store struct {
	// the base path for reading and writing objects to.
	BasePath string `env:"STORE_LOCAL_BASE_PATH,required"`
}

// Init initializes the local store.
func (s *Store) Init() error {
	return env.Parse(s)
}

// Read takes an ObjectID as an argument and attempts to read the corresponding
// file from the filesystem. On success, a file is returned. On failure an
// error is returned with a nil Object pointer.
func (s *Store) Read(id fs.ObjectID) (*fs.Object, error) {
	return &fs.Object{}, nil
}

// Write takes an Object pointer as an argument. If the Object can be
// correctly written to disk, nil is returned. Otherwise an error is returned
// providing why the Object could not be correctly written.
func (s *Store) Write(obj *fs.Object) error {
	return nil
}

// Close returns nil to comply with the FileStore interface.
func (s *Store) Close() error {
	return nil
}
