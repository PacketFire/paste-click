package local

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/PacketFire/paste-click/lib/objectstore"
)

const (
	testingBasePath = "/tmp/"
)

var (
	ObjectStoragePath         = fmt.Sprintf("%s%s", testingBasePath, "abcdef.asc")
	ObjectMetadataStoragePath = fmt.Sprintf("%s%s", testingBasePath, "_abcdef")
	Object                    = objectstore.Object{
		Metadata: objectstore.Metadata{
			Size:     10,
			Mimetype: "text/plain",
			Uploaded: time.Now().String(),
			Object:   "abcdef",
		},
		Data: []byte("helloworld"),
	}
)

func initializeStoreForTesting() Store {
	return Store{
		BasePath: testingBasePath,
	}
}

func TestInit(t *testing.T) {
	os.Setenv("STORE_LOCAL_BASE_PATH", testingBasePath)
	defer os.Unsetenv("STORE_LOCAL_BASE_PATH")

	expectedStore := Store{
		BasePath: testingBasePath,
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

func TestStoreRead(t *testing.T) {
	s := initializeStoreForTesting()
	defer s.Close()

	t.Run("Should be able to read a file.", func(t *testing.T) {
		err := ioutil.WriteFile(ObjectStoragePath, Object.Data, 0644)
		if err != nil {
			t.Errorf("Unable to write temporary file, %v.", ObjectStoragePath)
		}
		defer os.Remove(ObjectStoragePath)

		metadata, _ := Object.Metadata.JSON()
		err = ioutil.WriteFile(ObjectMetadataStoragePath, metadata, 0644)
		if err != nil {
			t.Errorf("Unable to write temporary metadata file, %v.", ObjectMetadataStoragePath)
		}
		defer os.Remove(ObjectMetadataStoragePath)

		if obj, err := s.Read(Object.Metadata.Object); err != nil && reflect.DeepEqual(obj, Object) {
			t.Errorf("Unable to read the file from disk, got %v want nil", err)
		}
	})

	t.Run("Reading a non-existent file should return an error.", func(t *testing.T) {
		if _, err := s.Read(Object.Metadata.Object); err == nil {
			t.Error("File doesn't exist and should error, got nil want error")
		}
	})

	t.Run("Invalid json in metadata file should return an error.", func(t *testing.T) {
		metadata := []byte(`{"hello: "world}`)

		err := ioutil.WriteFile(ObjectMetadataStoragePath, metadata, 0644)
		if err != nil {
			t.Errorf("Unable to write temporary metadata file, %v.", ObjectMetadataStoragePath)
		}
		defer os.Remove(ObjectMetadataStoragePath)
		if _, err := s.Read(`abcdef`); err == nil {
			t.Error("Invalid json should throw an error, got nil want error")
		}
	})
}

func TestStoreWrite(t *testing.T) {
	s := initializeStoreForTesting()
	defer s.Close()

	t.Run("an unallocated ID should successfully write", func(t *testing.T) {
		if err := s.Write(&Object); err != nil {
			t.Errorf("Unable to write the file to disk, got %v want nil", err)
		}
	})
}

func TestClose(t *testing.T) {
	s := initializeStoreForTesting()
	if err := s.Close(); err != nil {
		t.Errorf("close should aways return nil, got %v want nil", err)
	}
}
