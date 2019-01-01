package logging

import (
	"bytes"
	"log"
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
	rr := httptest.NewRecorder()

	logBuffer := new(bytes.Buffer)
	sl := log.New(logBuffer, "", log.LstdFlags)
	middleware := New(sl)

	router.Use(middleware.Serve)

	router.ServeHTTP(rr, req)

	if !bytes.Contains(logBuffer.Bytes(), []byte("GET /")) {
		t.Error(`Log should contain: "GET /"`)
	}
}
