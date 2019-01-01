package objectstore

import (
	"crypto/md5"

	"github.com/PacketFire/paste-click/lib/objectstore/metadata"
	"github.com/PacketFire/paste-click/lib/objectstore/objectid"
)

// Object stores contains both the data and metadata for an object written to
// or read from the objectstore.
type Object struct {
	Metadata metadata.Metadata `json:"metadata"`
	Data     []byte            `json:"data"`
}

// New Instantiates a new Object.
func New(mimetype string, data []byte) *Object {
	sum := md5.New()
	sum.Write(data)
	oid := objectid.New(sum)

	md := metadata.New(int64(len(data)), mimetype, oid)
	return &Object{
		Metadata: *md,
		Data:     data,
	}
}
