package objectid

import "testing"

const (
	oidString = `abcdef`
)

func TestObjectIDToString(t *testing.T) {
	var oid ObjectID = oidString

	if oid != oidString {
		t.Errorf("ObjectID didn't correctly convert to a string, expected %v got %v", oidString, oid)
	}
}
