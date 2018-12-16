package logging

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggerReturnsCorrectFormat(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	rr := httptest.NewRecorder()

	logBuffer := new(bytes.Buffer)
	sl := log.New(logBuffer, "", log.LstdFlags)
	handler := New(sl, testHandler)

	handler.ServeHTTP(rr, req)

	if !bytes.Contains(logBuffer.Bytes(), []byte("GET /")) {
		t.Error(`Log should contain: "GET /"`)
	}
}
