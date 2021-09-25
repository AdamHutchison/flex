package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AdamHutchison/flux/bootstrap"
)

func testHandler(httpMethod string, route string, handler http.Handler, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest(httpMethod, route, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	return rr
}

func getRoute(name string, t *testing.T) string {
	kernal := bootstrap.NewKernal()
	router := kernal.GetRouter()

	url, err := router.Get(name).URL()

	if err != nil {
		t.Fatal(err)
	}

	return url.Path
}

func assertStatus(response *httptest.ResponseRecorder, httpStatus int, t *testing.T) {
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
