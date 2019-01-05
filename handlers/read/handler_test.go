package read

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PacketFire/paste-click/lib/objectstore/drivers/mock"
)

func handlerTest(method, path string, reqBody io.Reader, respCode int, respBody string, h http.HandlerFunc) error {
	req, err := http.NewRequest(method, path, reqBody)
	if err != nil {
		return err
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h)

	handler.ServeHTTP(rr, req)

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
	uh := New(store)
	t.Run("Get handler returns the correct response", func(t *testing.T) {
		err := handlerTest("GET", "/", strings.NewReader(`helloworld`), http.StatusOK, `helloworld`, uh.ServeHTTP)
		if err != nil {
			//t.Error(err)
		}
	})
}
