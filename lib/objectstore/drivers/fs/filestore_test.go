package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/PacketFire/paste-click/lib/objectstore/metadata"
	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
)

const (
	BasePath = "/tmp/"
)

func initializeStoreForTesting() Store {
	return Store{
		BasePath: BasePath,
	}
}

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

func runTestWithTemporaryObject(obj *objectstore.Object, extension string, callback func()) error {
	dataPath := filepath.Join(BasePath, string(obj.Metadata.Object)+"."+extension)
	metadataPath := filepath.Join(BasePath, "_"+string(obj.Metadata.Object))
	err := ioutil.WriteFile(dataPath, obj.Data, 0644)
	if err != nil {
		return err
	}
	defer os.Remove(dataPath)

	metadata, _ := obj.Metadata.JSON()
	err = ioutil.WriteFile(metadataPath, metadata, 0644)
	if err != nil {
		return err
	}
	defer os.Remove(metadataPath)

	callback()
	return nil
}

func TestInit(t *testing.T) {
	os.Setenv("STORE_FS_BASE_PATH", BasePath)
	defer os.Unsetenv("STORE_FS_BASE_PATH")

	expectedStore := Store{
		BasePath: BasePath,
	}

	s := Store{}
	if err := s.Init(); err != nil {
		t.Errorf("Init should aways return nil, got %v want nil", err)
	}

	if !reflect.DeepEqual(s, expectedStore) {
		t.Errorf("Initialized store doesn't match expected value, got %v want %v",
			s,
			expectedStore)
	}
	defer s.Close()
}

func TestGetExtension(t *testing.T) {
	store := initializeStoreForTesting()

	t.Run("Validate that a valid mimetype returns the expected string.", func(t *testing.T) {
		if extension := store.getExtensionFromMimetype(`text/plain`); !(extension == `.asc` || extension == `.txt`) {
			t.Errorf("Extension didn't match mimetype, expected .asc got %s", extension)
		}
	})
	t.Run("Attempt to read an unknown mimetype should return an error.", func(t *testing.T) {
		if extension := store.getExtensionFromMimetype(`text/invalid1234`); extension != `` {
			t.Errorf("An invalid mimetype should return an empty string, got %v  want \"\"", extension)
		}
	})
}

func TestStoreRead(t *testing.T) {
	s := initializeStoreForTesting()
	defer s.Close()

	objectMetadataStoragePath := filepath.Join(BasePath, "_abcdef")
	object := generateObject(`abcdef`, `text/plain`, `helloworld`)

	t.Run("Should be able to read a file.", func(t *testing.T) {
		runTestWithTemporaryObject(&object, `asc`, func() {
			if obj, err := s.Read(object.Metadata.Object); err != nil && reflect.DeepEqual(obj, object) {
				t.Errorf("Unable to read the file from disk, got %v want nil", err)
			}
		})
	})

	t.Run("Reading a non-existent file should return an error.", func(t *testing.T) {
		if _, err := s.Read(object.Metadata.Object); err == nil {
			t.Error("File doesn't exist and should error, got nil want error")
		}
	})

	t.Run("Invalid json in metadata file should return an error.", func(t *testing.T) {
		metadata := []byte(`{"hello: "world}`)

		err := ioutil.WriteFile(objectMetadataStoragePath, metadata, 0644)
		if err != nil {
			t.Errorf("Unable to write temporary metadata file, %v.", objectMetadataStoragePath)
		}
		defer os.Remove(objectMetadataStoragePath)

		if _, err := s.Read(`abcdef`); err == nil {
			t.Error("Invalid json should throw an error, got nil want error")
		}
	})

	t.Run("Attempt to read an unknown mimetype should return an error.", func(t *testing.T) {
		invalidObject := generateObject(`invalidobjecttest`, `text/invalidmimetype1234`, `helloworld`)
		runTestWithTemporaryObject(&invalidObject, "invalid", func() {
			if _, err := s.Read(`invalidobjecttest`); err == nil {
				t.Error(`An invalid mimetype should return an error, got nil want an "unknown mimetype extension" error`)
			}
		})
	})
}

func TestStoreWrite(t *testing.T) {
	s := initializeStoreForTesting()
	defer s.Close()

	object := generateObject(`abcdef`, `text/plain`, `helloworld`)

	t.Run("an unallocated ID should successfully write", func(t *testing.T) {
		if err := s.Write(&object); err != nil {
			t.Errorf("Unable to write the file to disk, got %v want nil", err)
		}
	})

	os.Remove(filepath.Join(BasePath, "_abcdef"))
	os.Remove(filepath.Join(BasePath, "abcdef.asc"))
}

func TestClose(t *testing.T) {
	s := initializeStoreForTesting()
	if err := s.Close(); err != nil {
		t.Errorf("close should aways return nil, got %v want nil", err)
	}
}
