package gcs

import (
	"github.com/fsouza/fake-gcs-server/fakestorage"

	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"github.com/PacketFire/paste-click/lib/objectstore"
	"testing"
	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
)

const (
	bucketName = `test_bucket`
)

func runTestWithTemporaryObject(obj *objectstore.Object, bucketName string, callback func(*fakestorage.Server)) error {
	server := fakestorage.NewServer([]fakestorage.Object{
		{
			BucketName: bucketName,
			Name:       string(obj.Metadata.Object),
			Content:    obj.Data,
		},
	})
	defer server.Stop()

	callback(server)
	return nil
}

func initMockStore(c *storage.Client) *Store {
	return &Store{
		BucketName: bucketName,
		client:     c,
		bucket:     c.Bucket(bucketName),
		ctx:        context.Background(),
	}
}

func TestStoreRead(t *testing.T) {
	data := []byte(`helloworld`)
	object := objectstore.New(`text/plain`, data)

	t.Run("Should be able to read a valid object", func(t *testing.T) {
		runTestWithTemporaryObject(object, bucketName, func(s *fakestorage.Server) {
			store := initMockStore(s.Client())
			so, err := store.Read(object.Metadata.Object)
			if err != nil {
				t.Error(err)
			}
			if bytes.Compare(so.Data, data) != 0 {
				t.Errorf("Data field doesn't match, got %v, want %v", so.Data, data)
			}
		})
	})

	t.Run("Reading a non-existent object should return an error", func(t *testing.T) {
		runTestWithTemporaryObject(object, bucketName, func(s *fakestorage.Server) {
			store := initMockStore(s.Client())
			_, err := store.Read(objectid.ObjectID("InvalidID"))
			if err == nil {
				t.Error("Object should throw an error if it doesn't exist")
			}
		})
	})
}

func TestStoreClose(t *testing.T) {
	data := []byte(`helloworld`)
	object := objectstore.New(`text/plain`, data)

	runTestWithTemporaryObject(object, bucketName, func(s *fakestorage.Server) {
		store := initMockStore(s.Client())
		if store.Close() != nil {
			t.Error("Store should successfully close session with server.")
		}
	})
}

