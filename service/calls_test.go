package service

import (
	"galleryline/auth"
	"galleryline/domain"
	"galleryline/signaling"
	"galleryline/storage"
	"path/filepath"
	"testing"
)

func TestCallReject(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "d"))
	defer s.Close()
	d := NewDirectory(s)
	d.Register(domain.Extension{ID: "a", Nickname: "A"})
	d.Register(domain.Extension{ID: "b", Nickname: "B"})
	m := NewCallManager(s, d, signaling.NewRouter())
	m.Dial(auth.Principal{ID: "a", Role: "guide"}, "a", "b", "x")
	if e := m.Reject("x"); e != nil {
		t.Fatal(e)
	}
}
