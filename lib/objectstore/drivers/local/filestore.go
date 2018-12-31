package local

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/caarlos0/env"
)

// Store implements the ObjectStore interface to store Objects on the local
// filesystem.
type Store struct {
	// the base path for reading and writing objects to.
	BasePath string `env:"STORE_LOCAL_BASE_PATH,required"`
}

// Init initializes the local store.
func (s *Store) Init() error {
	return env.Parse(s)
}

// readMetadata takes an object ID and attempts to read the metadata for a file.
// On success a pointer to metadata is returned. On fail nil and a corresponding
// error is returned.
func (s *Store) readMetadata(id objectstore.ObjectID) (*objectstore.Metadata, error) {
	md := &objectstore.Metadata{}
	mdPath := filepath.Join(s.BasePath, "_"+string(id))
	rawMdData, err := ioutil.ReadFile(mdPath)
	if err != nil {
		return md, err
	}

	err = json.Unmarshal(rawMdData, md)
	if err != nil {
		return md, err
	}

	return md, nil
}

// Read takes an ObjectID as an argument and attempts to read the corresponding
// file from the filesystem. On success, a file is returned. On failure an
// error is returned with a nil Object pointer.
func (s *Store) Read(id objectstore.ObjectID) (*objectstore.Object, error) {
	metadata, err := s.readMetadata(id)
	if err != nil {
		return nil, err
	}

	return &objectstore.Object{
		Metadata: *metadata,
	}, nil
}

// Write takes an Object pointer as an argument. If the Object can be
// correctly written to disk, nil is returned. Otherwise an error is returned
// providing why the Object could not be correctly written.
func (s *Store) Write(obj *objectstore.Object) error {
	return nil
}

// Close returns nil to comply with the FileStore interface.
func (s *Store) Close() error {
	return nil
}
