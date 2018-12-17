package local

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	fs "github.com/PacketFire/paste-click/lib/filestore"
)

var (
	UnusedObjectStoragePath = "/tmp/abcdef.txt"
	UnusedObject            = fs.Object{
		Metadata: fs.Metadata{
			Size:     10,
			Mimetype: "text/plain",
			Uploaded: time.Now().String(),
			Object:   "abcdef",
		},
		Data: []byte("helloworld"),
	}
)

func TestInit(t *testing.T) {
	s := Store{}
	if err := s.Init(); err != nil {
		t.Errorf("Init should aways return nil, got %v want nil", err)
	}
	defer s.Close()
}

func TestStoreRead(t *testing.T) {
	s := Store{}
	s.Init()
	defer s.Close()

	err := ioutil.WriteFile(UnusedObjectStoragePath, UnusedObject.Data, 0644)
	if err != nil {
		t.Errorf("Unable to write temporary file, %v.", UnusedObjectStoragePath)
	}
	defer os.Remove(UnusedObjectStoragePath)

	t.Run("an unallocated ID should successfully write", func(t *testing.T) {
		if obj, err := s.Read(UnusedObject.Metadata.Object); err != nil && obj != nil {
			t.Errorf("unable to write the file to disk, got %v want nil", err)
		}
	})
}

func TestStoreWrite(t *testing.T) {
	s := Store{}
	s.Init()
	defer s.Close()

	t.Run("an unallocated ID should successfully write", func(t *testing.T) {
		if err := s.Write(&UnusedObject); err != nil {
			t.Errorf("unable to write the file to disk, got %v want nil", err)
		}
	})
}

func TestClose(t *testing.T) {
	s := Store{}
	s.Init()
	if err := s.Close(); err != nil {
		t.Errorf("close should aways return nil, got %v want nil", err)
	}
}
