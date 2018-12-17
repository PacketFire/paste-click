package filestore

import "testing"

func TestMarshallMetadataToJSON(t *testing.T) {
	md := NewMetadata(100, "application/json", "test.json", `aD490B`)
	_, err := md.JSON()
	if err != nil {
		t.Error("Received error when attempting to marshall JSON.")
	}
}
