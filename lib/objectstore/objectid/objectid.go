package objectid

import (
	"encoding/base64"
	"hash"
)

// ObjectID contains the string representation of an objects identifier.
type ObjectID string

// New takes a Hash as an argument and uses that to generate a new ObjectID.
func New(h hash.Hash) ObjectID {
	// Generate checksum and truncate it to 6 characters
	return ObjectID(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}
