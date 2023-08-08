package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandler(t *testing.T) {
	srv := server{
		router: chi.NewRouter(),
	}
	srv.routes()
	rec := httptest.NewRecorder()

	srv.ServeHTTP(
		rec,
		httptest.NewRequest(
			http.MethodGet,
			"/route",
			strings.NewReader(``),
		),
	)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != `OK` {
		t.Errorf("expected body %q, got %q", `OK`, rec.Body.String())
	}
}
