package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"syscall"
	"time"
)

// FileMetadata stores file specific data for individual pastes.
type FileMetadata struct {
	AccessCount int64  `json:"access_count"`
	Size        int64  `json:"size"`
	Bandwidth   int64  `json:"bandwidth,omitempty"`
	Mimetype    string `json:"mime_type"`
	Filename    string `json:"filename,omitempty"`
	Uploaded    string `json:"uploaded"`
	Path        string `json:"path"`
	Atime       string `json:"atime,omitempty"`
}

func NewFileMetadata(size int64, mime, filename, path string) *FileMetadata {
	return &FileMetadata{
		AccessCount: 0,
		Size:        size,
		Bandwidth:   0,
		Mimetype:    mime,
		Filename:    filename,
		Uploaded:    time.Now().String(),
		Path:        path,
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

	fm.Bandwidth = fm.bandwidth()

	return fm, nil
}

// bandwidth calculates and returns the bandwidth used by multiplying
// AccessCount by Size to estimate the total used bandwidth. This value is
// returned as a int64.
func (this *FileMetadata) bandwidth() int64 {
	return this.Size * this.AccessCount
}

// atime returns that access time of a FileMeta.
func (this *FileMetadata) atime() (time.Time, error) {
	fi, err := os.Stat(this.Path)
	if err != nil {
		return time.Time{}, err
	}

	stat := fi.Sys().(*syscall.Stat_t)
	atime := time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))

	return atime, nil
}

// JSON attempts to render FileMetaData to json.
func (this *FileMetadata) JSON() ([]byte, error) {
	return json.Marshal(this)
}
