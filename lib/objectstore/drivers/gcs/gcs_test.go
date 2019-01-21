package gcs

import (
	"github.com/fsouza/fake-gcs-server/fakestorage"

	"cloud.google.com/go/storage"
	"context"
	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/PacketFire/paste-click/lib/objectstore/metadata"
	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
	"testing"
	"time"
)

const (
	bucketName = `test_bucket`
)

func generateObject(objectID, mimetype, data string) objectstore.Object {
	return objectstore.Object{
		Metadata: metadata.Metadata{
			Size:     int64(len(data)),
			Mimetype: mimetype,
			Uploaded: time.Now().String(),
			Object:   objectid.ObjectID(objectID),
		},
		Data: []byte(data),
	}
}

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
	t.Run("Should be able to read a file.", func(t *testing.T) {
		object := generateObject(`abcdef`, `text/plain`, `helloworld`)

		runTestWithTemporaryObject(&object, bucketName, func(s *fakestorage.Server) {})
	})
}

