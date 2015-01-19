package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if w.Header().Get("Cache-Control") == "" {
			t.Errorf("got %v want a value", w.Header().Get("Cache-Control"))
		}

		if w.Header().Get("Expires") == "" {
			t.Errorf("got %v want a value", w.Header().Get("Expires"))
		}

		w.Write([]byte("Hello"))

	})

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	Cache(0, h).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got %v want 200", w.Code)
	}
}

