package galleryline

import (
	"galleryline/auth"
	"galleryline/domain"
	"galleryline/lifecycle"
	"galleryline/service"
	"galleryline/signaling"
	"galleryline/storage"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "a"))
	defer s.Close()
	d := service.NewDirectory(s)
	d.Register(domain.Extension{ID: "g", Nickname: "Guide"})
	d.Register(domain.Extension{ID: "desk", Nickname: "Desk"})
	m := service.NewCallManager(s, d, signaling.NewRouter())
	c, e := m.Dial(auth.Principal{ID: "g", Role: "guide"}, "g", "desk", "c1")
	if e != nil {
		t.Fatal(e)
	}
	if c.Status != "ringing" {
		t.Fatal(c)
	}
	if e = m.Accept("c1"); e != nil {
		t.Fatal(e)
	}
	if e = m.End("c1", "desk"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "a"))
	defer s.Close()
	d := service.NewDirectory(s)
	d.Register(domain.Extension{ID: "desk", Nickname: "Desk"})
	d.Register(domain.Extension{ID: "offline", Nickname: "Offline", Online: false})
	d.SetPresence("offline", false)
	m := service.NewCallManager(s, d, signaling.NewRouter())
	_, e := m.Dial(auth.Principal{ID: "desk", Role: "desk"}, "desk", "offline", "c2")
	if e == nil {
		t.Fatal("self dial accepted")
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "a"))
	defer s.Close()
	d := service.NewDirectory(s)
	d.Register(domain.Extension{ID: "g", Nickname: "Guide"})
	d.Register(domain.Extension{ID: "d", Nickname: "Device"})
	m := service.NewCallManager(s, d, signaling.NewRouter())
	_, e := m.Dial(auth.Principal{ID: "g", Role: "guide"}, "g", "d", "c3")
	if e != nil {
		t.Fatal(e)
	}
	if e = m.Reject("c3"); e != nil {
		t.Fatal(e)
	}
	if e = m.Accept("c3"); e == nil {
		t.Fatal("accepted rejected call")
	}
}
func TestBusinessChain21(t *testing.T) {
	tr := lifecycle.NewTracker()
	if e := lifecycle.CloseSession(tr); e != nil {
		t.Fatal(e)
	}
	got := tr.Events()
	want := []string{"transport", "audio"}
	if len(got) != len(want) {
		t.Fatalf("close order %v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("close order %v", got)
		}
	}
}
