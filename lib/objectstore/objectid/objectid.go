package objectid

import (
	"hash"

	"github.com/lytics/base62"
)

// ObjectID contains the string representation of an objects identifier.
type ObjectID string

// New takes a Hash as an argument and uses that to generate a new ObjectID.
func New(h hash.Hash) ObjectID {
	// Generate checksum and truncate it to 6 characters
	//checksum := fmt.Sprintf("%x", h.Sum(nil))

	return ObjectID(base62.StdEncoding.EncodeToString(h.Sum(nil)))
}
