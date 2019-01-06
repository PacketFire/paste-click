package metadata

import "testing"

func TestMarshallMetadataToJSON(t *testing.T) {
	md := New(100, "application/json", `aD490B`)
	_, err := md.JSON()
	if err != nil {
		t.Error("Received error when attempting to marshall JSON.")
	}
}
