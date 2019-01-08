package logging

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestLoggerReturnsCorrectFormat(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")

	logBuffer := new(bytes.Buffer)
	middleware := New(logBuffer)

	router.Use(middleware.Serve)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if !bytes.Contains(logBuffer.Bytes(), []byte("GET /")) {
		t.Errorf(`Log doesn't contain expected string got '%s' want 'GET /'`, logBuffer.Bytes())
	}
}

