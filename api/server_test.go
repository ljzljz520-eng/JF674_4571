package api

import (
	"galleryline/service"
	"galleryline/signaling"
	"galleryline/storage"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealth(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "d"))
	defer s.Close()
	h := NewServer(service.NewDirectory(s), service.NewCallManager(s, service.NewDirectory(s), signaling.NewRouter())).Handler()
	r := httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 204 {
		t.Fatal(r.Code)
	}
}
