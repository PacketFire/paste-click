package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"time"
)

// FileMetadata stores file specific data for individual pastes.
type FileMetadata struct {
	Size     int64  `json:"size"`
	Mimetype string `json:"mime_type"`
	Filename string `json:"filename,omitempty"`
	Uploaded string `json:"uploaded"`
}

func NewFileMetadata(size int64, mime, filename, object string) *FileMetadata {
	return &FileMetadata{
		Size:     size,
		Mimetype: mime,
		Filename: filename,
		Uploaded: time.Now().String(),
	}
}

// OpenMetaFile attemps to open an unmarshal a meta file into a FileMetadata
// struct. On success a *FileMetadata and nil is returned. On failure a nil
// pointer and a corresponding error is returned.
func OpenMetaFile(path string) (*FileMetadata, error) {
	var fb bytes.Buffer
	fm := new(FileMetadata)

	f, err := os.Open(path)
	if err != nil {
		return fm, err
	}
	defer f.Close()

	_, err = io.Copy(&fb, f)
	if err != nil {
		return fm, err
	}

	if err = json.Unmarshal(fb.Bytes(), &fm); err != nil {
		return fm, err
	}

	return fm, nil
}

// JSON attempts to render FileMetaData to json.
func (this *FileMetadata) JSON() ([]byte, error) {
	return json.Marshal(this)
}
