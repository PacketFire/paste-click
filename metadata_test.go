package main

import (
	"testing"
)

const (
	TestMP4Path     string = "testfiles/test.mp4"
	TestMP4MetaPath string = "testfiles/_test"
)

func TestFileMetaUnmarshal(t *testing.T) {
	fm, err := OpenMetaFile(TestMP4MetaPath)
	if err != nil {
		t.Error(err)
	}

	if fm.AccessCount != 10 {
		t.Error("FileMeta.Size doesn't match test file.")
	}
}

func TestFileMetaBandwidth(t *testing.T) {
	fm, err := OpenMetaFile(TestMP4MetaPath)
	if err != nil {
		t.Error(err)
	}

	if fm.bandwidth() != 105302080 {
		t.Error("fm.Bandwidth() calculation doesn't match estimate.")
	}
}

func TestFileMetaAtime(t *testing.T) {
	fm, err := OpenMetaFile(TestMP4MetaPath)
	if err != nil {
		t.Error(err)
	}

	_, err = fm.atime()
	if err != nil {
		t.Error("Unable to fetch atime.")
	}
}

func TestFileMetaJSON(t *testing.T) {
	fm, err := OpenMetaFile(TestMP4MetaPath)
	if err != nil {
		t.Error(err)
	}

	_, err = fm.JSON()
	if err != nil {
		t.Error(err)
	}
}
