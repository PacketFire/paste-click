package objectstore

import "github.com/PacketFire/paste-click/lib/objectstore/metadata"

// Object stores contains both the data and metadata for an object written to
// or read from the objectstore.
type Object struct {
	Metadata metadata.Metadata `json:"metadata"`
	Data     []byte            `json:"data"`
}

// New Instantiates a new Object.
func New(data []byte) *Object {
	return &Object{}
}
