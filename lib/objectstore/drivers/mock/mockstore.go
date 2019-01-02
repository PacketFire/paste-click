package mock

import (
	"errors"
	"sync"

	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
)

// Store implements the ObjectStore interface to store Objects in memory for
// mocking.
type Store struct {
	sync.RWMutex
	MemStore map[objectid.ObjectID]*objectstore.Object
}

// Init initializes the mock store.
func (s *Store) Init() error {
	s.MemStore = make(map[objectid.ObjectID]*objectstore.Object)
	return nil
}

// Read takes an ObjectID as an argument and attempts to lookup the object in
// it's MemStore. If the object is present it is returned. If not, an error is
// returned.
func (s *Store) Read(id objectid.ObjectID) (*objectstore.Object, error) {
	s.RLock()
	defer s.RUnlock()

	if obj, prs := s.MemStore[id]; prs == true {
		return obj, nil
	}

	return nil, errors.New("object not present")
}

// Write takes an Object pointer as an argument. If the object can be assigned
// to the MemStore nil is returned otherwise an error is returned.
func (s *Store) Write(obj *objectstore.Object) error {
	s.Lock()
	defer s.Unlock()

	id := obj.Metadata.Object
	if _, prs := s.MemStore[id]; prs == false {
		s.MemStore[id] = obj
		return nil
	}

	return errors.New("Object already defined")
}

// Close returns nil to comply with the FileStore interface.
func (s *Store) Close() error {
	return nil
}
