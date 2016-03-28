package main

import (
	"testing"
)

func TestNewMimeMapValid(t *testing.T) {
	mimeTypes := map[string]string{
		"video/ogg":  ".ogv",
		"audio/midi": ".mid",
		"image/jpeg": ".jpeg",
	}

	mMap := new(MimeMap)
	mMap.New()
	for k, v := range mimeTypes {
		if mMap.m[k] != v {
			t.Errorf("Extension %v not initialized in MimeMap.", k)
		}
	}
}

func TestNewMimeMapInvalid(t *testing.T) {
	mimeNonExists := map[string]string{
		"text/nonexist": ".non",
	}

	mMap := new(MimeMap)
	mMap.New()
	for k, v := range mimeNonExists {
		if mMap.m[k] == v {
			t.Errorf("Extension %v should not correctly return.", k)
		}
	}
}

func TestExtensionValid(t *testing.T) {
	mimeTypes := map[string]string{
		"video/ogg":  ".ogv",
		"audio/midi": ".mid",
		"image/jpeg": ".jpeg",
	}

	mMap := new(MimeMap)
	mMap.New()
	for k, v := range mimeTypes {
		if s, e := mMap.Extension(k); e != nil || s != v {
			t.Errorf("Unable to fetch extension %v from MimeMap using Extension method.", k)
		}
	}
}

func TestExtensionInvalid(t *testing.T) {
	mimeNonExists := map[string]string{
		"text/nonexist": ".non",
	}

	mMap := new(MimeMap)
	mMap.New()
	for k, v := range mimeNonExists {
		if s, e := mMap.Extension(k); e == nil || s == v {
			t.Errorf("Incorrectly fetched extension %v from MimeMap using Extension method.", k)
		}
	}
}

func TestNewExtensionValid(t *testing.T) {
	newExtensions := map[string]string{
		"text/nonexist": ".non",
	}

	mMap := new(MimeMap)
	mMap.New()
	for k, v := range newExtensions {
		if e := mMap.NewExtension(k, v); e != nil {
			t.Errorf("Attempt to add new valid extension %v failed with an error")
		}
	}

	for k, v := range newExtensions {
		if s, e := mMap.Extension(k); e != nil || s != v {
			t.Errorf("Unable to add new extension %v to MimeMap.", k)
		}
	}
}

func TestNewExtensionInvalid(t *testing.T) {
	newInvalidExtensions := map[string]string{
		"text/vtt": ".vtt",
	}

	mMap := new(MimeMap)
	mMap.New()
	for k, v := range newInvalidExtensions {
		if e := mMap.NewExtension(k, v); e == nil {
			t.Errorf("Adds duplicate record %v to MimeMap.", k)
		}
	}
}

func generateRandSeq(t *testing.T) {
	seq := randSeq(6, "")

	if seq[0] != '_' {
		t.Errorf("Doesn't contain underscore in Rand Sequence")
	}

	if len(seq) != 7 {
		t.Errorf("RandSeq returns the wrong sized sequence.")
	}
}
