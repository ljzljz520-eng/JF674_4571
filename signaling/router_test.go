package signaling

import (
	"galleryline/domain"
	"testing"
)

func TestRouter(t *testing.T) {
	r := NewRouter()
	r.Send(domain.SignalMessage{From: "a", To: "b", Kind: "offer"})
	if r.Pending("b") != 1 {
		t.Fatal("pending")
	}
	if len(r.Receive("b")) != 1 {
		t.Fatal("receive")
	}
}
