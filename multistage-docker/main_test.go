package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	// Mock request
	req := httptest.NewRequest(http.MethodGet, "/minnie", nil)
	w := httptest.NewRecorder()

	// Expected values
	expectedStatusCode := 200
	expectedBody := "Hello, minnie!"

	// Invoke the handler
	HelloServer(w, req)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error("Unexpected error during body parsing", err)
	}

	// Check status code and response body
	if res.StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, res.StatusCode)
	}

	if string(body) != expectedBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedBody, body)
	}
}
