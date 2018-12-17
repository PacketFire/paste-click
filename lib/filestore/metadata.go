package filestore

import (
	"encoding/json"
	"time"
)

// ObjectID contains the string representation of an objects identifier.
type ObjectID string

// Metadata stores file specific data for individual pastes.
type Metadata struct {
	Size     int64    `json:"size"`
	Mimetype string   `json:"mime_type"`
	Filename string   `json:"filename,omitempty"`
	Uploaded string   `json:"uploaded"`
	Object   ObjectID `json:"object"`
}

// NewMetadata instantiates takes the size, mimetype, filename and object ID
// as arguments and uses this to instantiate a new instance of Metadata,
// returning a pointer to it.
func NewMetadata(size int64, mime, filename string, object ObjectID) *Metadata {
	return &Metadata{
		Size:     size,
		Mimetype: mime,
		Filename: filename,
		Uploaded: time.Now().String(),
		Object:   object,
	}
}

// JSON attempts to marshall MetaData to json.
func (meta *Metadata) JSON() ([]byte, error) {
	return json.Marshal(meta)
}
