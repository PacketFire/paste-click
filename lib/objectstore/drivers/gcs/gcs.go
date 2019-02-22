package gcs

import (
	"bytes"
	"fmt"
	"io"

	"github.com/PacketFire/paste-click/lib/objectstore/metadata"

	"context"

	"cloud.google.com/go/storage"
	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
	"github.com/caarlos0/env"
)

// Store implements the ObjectStore interface to store Objects in GCS
type Store struct {
	BucketName string `env:"STORE_GCS_BUCKET_NAME,required"`
	client     *storage.Client
	ctx        context.Context
}

// Init initializes the gcs store.
func (s *Store) Init() error {
	s.ctx = context.Background()
	client, err := storage.NewClient(s.ctx)
	if err != nil {
		return fmt.Errorf("Unable to initialize client: %s", err)
	}

	if err = env.Parse(s); err != nil {
		return err
	}

	s.client = client

	return nil
}

// Read takes an ObjectID as an argument and attempts to read the corresponding
// object from the GCS bucket. On success, a file is returned. On failure an
// error is returned with a nil Object pointer.
func (s *Store) Read(id objectid.ObjectID) (*objectstore.Object, error) {
	bucket := s.client.Bucket(s.BucketName)
	obj := bucket.Object(string(id))

	r, err := obj.NewReader(s.ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	bodyBuf := new(bytes.Buffer)
	if _, err := io.Copy(bodyBuf, r); err != nil {
		return nil, err
	}

	attrs, err := obj.Attrs(s.ctx)
	if err != nil {
		return nil, err
	}

	metadata := &metadata.Metadata{
		Object:   id,
		Mimetype: attrs.ContentType,
		Uploaded: attrs.Created.String(),
		Size:     attrs.Size,
	}

	return &objectstore.Object{
		Metadata: *metadata,
		Data:     bodyBuf.Bytes(),
	}, nil
}

// Write takes an Object pointer as an argument. If the Object can be
// correctly written to the GCS bucket, nil is returned. Otherwise an error is
// returned providing why the Object could not be correctly written.
func (s *Store) Write(obj *objectstore.Object) error {
	oid := string(obj.Metadata.Object)

	bucket := s.client.Bucket(s.BucketName)
	objHandler := bucket.Object(oid)
	w := objHandler.NewWriter(s.ctx)
	if w == nil {
		return nil
	}

	w.ContentType = obj.Metadata.Mimetype

	buf := bytes.NewBuffer(obj.Data)

	if _, err := io.Copy(w, buf); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	// Set Attrs
	uAttrs := storage.ObjectAttrsToUpdate{
		ContentType: obj.Metadata.Mimetype,
	}

	if _, err := objHandler.Update(s.ctx, uAttrs); err != nil {
	}

	return nil
}

// Close cleans up any connections with the remote store.
func (s *Store) Close() error {
	return s.client.Close()
}
