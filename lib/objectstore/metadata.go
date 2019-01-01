package objectstore

import (
	"encoding/json"
	"time"

	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
)

// Metadata stores file specific data for individual pastes.
type Metadata struct {
	Size     int64             `json:"size"`
	Mimetype string            `json:"mime_type"`
	Uploaded string            `json:"uploaded"`
	Object   objectid.ObjectID `json:"object"`
}

// NewMetadata instantiates takes the size, mimetype and object ID as arguments
// and uses this to instantiate a new instance of Metadata, returning a pointer
// to it.
func NewMetadata(size int64, mime string, object objectid.ObjectID) *Metadata {
	return &Metadata{
		Size:     size,
		Mimetype: mime,
		Uploaded: time.Now().String(),
		Object:   object,
	}
}

// JSON attempts to marshall MetaData to json.
func (meta *Metadata) JSON() ([]byte, error) {
	return json.Marshal(meta)
}
