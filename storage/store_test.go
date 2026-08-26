package storage

import (
	"galleryline/domain"
	"path/filepath"
	"testing"
)

func TestStoreLists(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "d"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	s.SaveExtension(domain.Extension{ID: "b"})
	s.SaveExtension(domain.Extension{ID: "a"})
	v, e := s.ListExtensions()
	if e != nil || len(v) != 2 || v[0].ID != "a" {
		t.Fatal(v, e)
	}
}
