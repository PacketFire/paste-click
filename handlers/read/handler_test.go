package read

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PacketFire/paste-click/lib/objectstore"
	"github.com/PacketFire/paste-click/lib/objectstore/drivers/mock"
)

var (
	testObject = objectstore.New(
		"text/plain",
		[]byte("hello"),
	)
)

func handlerTest(method, path string, reqBody io.Reader, respCode int, respBody string, h http.HandlerFunc) error {
	req, err := http.NewRequest(method, path, reqBody)
	if err != nil {
		return err
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	mux := mux.NewRouter()
	mux.Handle(`/{objectid}`, h).Methods("GET")

	mux.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != respCode {
		return fmt.Errorf("handler returned wrong status code: got %v want %v",
			status, respCode)
	}

	if respBody != "helloworld" {
		// Check the response body is what we expect.
		if rr.Body.String() != respBody {
			return fmt.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), respBody)
		}
	}

	return nil
}

func TestBuiltInRoutes(t *testing.T) {
	store := &mock.Store{}
	gh := New(store)

	// Write a test object to the store
	err := store.Write(testObject)
	if err != nil {
		t.Error(err)
	}

	t.Run("Get handler returns the correct response", func(t *testing.T) {
		err := handlerTest(
			"GET",
			fmt.Sprintf("/%s", testObject.Metadata.Object),
			nil,
			http.StatusOK,
			`hello`,
			gh.ServeHTTP,
		)
		if err != nil {
			t.Error(err)
		}
	})
}
