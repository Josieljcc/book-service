package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestBookHandler_MissingISBN(t *testing.T) {
	req := httptest.NewRequest("GET", "/book/", nil)
	rr := httptest.NewRecorder()
	BookHandler(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "ISBN not provided") {
		t.Errorf("expected error message about missing ISBN, got: %s", rr.Body.String())
	}
}

func TestBookHandler_NoAPIKey(t *testing.T) {
	os.Unsetenv("GOOGLE_BOOKS_API_KEY")
	req := httptest.NewRequest("GET", "/book/1234567890", nil)
	rr := httptest.NewRecorder()
	BookHandler(rr, req)
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "API key not configured") {
		t.Errorf("expected error message about missing API key, got: %s", rr.Body.String())
	}
}

// Real integration or httpmock tests can be added for more coverage.
