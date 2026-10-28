package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeFavicon(t *testing.T) {
	recorder := httptest.NewRecorder()
	serveFavicon(recorder, httptest.NewRequest(http.MethodGet, "/favicon.svg", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	if got := recorder.Header().Get("Content-Type"); got != "image/svg+xml; charset=utf-8" {
		t.Fatalf("Content-Type = %q", got)
	}
	if !strings.Contains(recorder.Body.String(), "<svg") {
		t.Fatal("favicon response is not SVG")
	}
}
