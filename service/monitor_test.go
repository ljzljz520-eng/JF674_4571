package service

import (
	"galleryline/domain"
	"galleryline/storage"
	"path/filepath"
	"testing"
)

func TestMonitor(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "d"))
	defer s.Close()
	d := NewDirectory(s)
	d.Register(domain.Extension{ID: "a", Nickname: "A"})
	if NewMonitor(s, d).CountOnline() != 1 {
		t.Fatal("online")
	}
}
