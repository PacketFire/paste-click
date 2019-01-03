package objectstore

import (
	"testing"
)

func TestObjectInitialization(t *testing.T) {
	obj := New(`text/plain`, []byte(`helloworld`))
	t.Log(obj.Metadata.Object)
}
