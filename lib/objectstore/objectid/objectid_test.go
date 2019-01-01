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
		expectedHash := `5d41402abc4b2a76b9719d911017c592`
		sum := md5.New()

		sum.Write(data)

		oid := New(sum)

		if string(oid) == expectedHash {
			t.Errorf("The md5 hasher should generate a consistent MD5 hash, got %v want %v", sum, expectedHash)
		}
	})
}
