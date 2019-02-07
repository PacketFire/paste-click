package gcs

import (
	"github.com/fsouza/fake-gcs-server/fakestorage"

	"bytes"
	"context"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/PacketFire/paste-click/lib/objectstore"
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

func TestStoreWrite(t *testing.T) {
	data := []byte(`write the world`)
	object := objectstore.New(`text/plain`, data)

	t.Run("An unallocated ID should successfully write", func(t *testing.T) {
		runTestWithTemporaryObject(object, bucketName, func(s *fakestorage.Server) {
			store := initMockStore(s.Client())
			store.Write(object)

			o, err := s.GetObject(bucketName, string(object.Metadata.Object))
			if err != nil {
				t.Errorf("Unable to write the object to the store, got error \"%v\" want nil", err)
			}

			if o.Name != string(object.Metadata.Object) {
				t.Error("Object wasn't written to store.")
			}

			if bytes.Compare(o.Content, object.Data) != 0 {
				t.Errorf("object data doesn't match, want %v got %v", o.Content, object.Data)
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
