package objectid

import (
	"crypto/md5"
	"testing"
)

const (
	oidString = `abcdef`
)

func TestObjectIDToString(t *testing.T) {
	var oid ObjectID = oidString

	if oid != oidString {
		t.Errorf("ObjectID didn't correctly convert to a string, expected %v got %v", oidString, oid)
	}
}

func TestObjectIDInstantiation(t *testing.T) {
	t.Run("ObjectID generates a consistent hash from a md5 hasher.", func(t *testing.T) {
		data := []byte("hello")
		expectedHash := "XUFAKrxLKna4cZ1REBfFkg++"
		sum := md5.New()

		sum.Write(data)

		oid := New(sum)

		if string(oid) != expectedHash {
			t.Errorf("The md5 hasher should generate a consistent MD5 hash, got %v want %v", oid, expectedHash)
		}
	})
}
