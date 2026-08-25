package galleryline

import (
	"galleryline/domain"
	"galleryline/storage"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "data.db")
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveExtension(domain.Extension{ID: "g1", Nickname: "Guide"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	v, e := s.GetExtension("g1")
	if e != nil || v.Nickname != "Guide" {
		t.Fatalf("%v %#v", e, v)
	}
}
