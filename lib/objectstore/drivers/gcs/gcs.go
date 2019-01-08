package gcs

import (
	"errors"

	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
	"github.com/caarlos0/env"
	"cloud.google.com/go/storage"
	"context"
)

// Store implements the ObjectStore interface to store Objects in GCS
type Store struct {
	// the base path for reading and writing objects to.
	BasePath  string `env:"STORE_FS_BASE_PATH,required"`
	BucketName string `env:"STORE_GCS_BUCKET_NAME,required"`
	client *storage.Client
	bucket *storage.BucketHandle
	gcsContext context.Context
}

// Init initializes the gcs store.
func (s *Store) Init() error {
	s.gcsContext = context.Background()
	client, err := storage.NewClient(s.gcsContext)
	if err != nil {
		return errors.New("Unable to initialize client")
	}

	if err = env.Parse(s); err != nil {
		return err
	}

	s.client = client
	s.bucket = client.Bucket(s.BucketName)

	return nil 
}

// Read takes an ObjectID as an argument and attempts to read the corresponding
// object from the GCS bucket. On success, a file is returned. On failure an
// error is returned with a nil Object pointer.
func (s *Store) Read(id objectid.ObjectID) (*objectstore.Object, error) {
	return nil, nil
}

// Write takes an Object pointer as an argument. If the Object can be
// correctly written to the GCS bucket, nil is returned. Otherwise an error is
// returned providing why the Object could not be correctly written.
func (s *Store) Write(obj *objectstore.Object) error {
	return nil
}

// Close cleans up any connections with the remote store.
func (s *Store) Close() error {
	return nil
}

