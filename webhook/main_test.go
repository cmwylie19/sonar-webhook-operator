package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cmwylie19/sonar-webhook-operator/webhook/handlers"
)

func TestHealthEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/webhook/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetRemote)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"x-forwarded-for":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
