package filestore

// Object stores contains both the data and metadata for an object written to
// or read from the filestore.
type Object struct {
	Metadata Metadata `json:"metadata"`
	Data     []byte   `json:"data"`
}
