package handlers

import (
	"net/http"
	"strings"
	"testing"

	"github.com/AdamHutchison/flux/http/handlers"
)

func TestHomeHandler(t *testing.T) {
	path := getRoute("home", t)

	response := testHandler(http.MethodGet, path, handlers.HomeHandler{}, t)

	assertStatus(response, http.StatusOK, t)

    expected := `{"data":{"message":"Welcome to your new flux app"}}`

	body := strings.Replace(response.Body.String(), "\n", "", -1)

    if body != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            response.Body.String(), expected)
    }
}