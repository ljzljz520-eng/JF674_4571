package auth

import (
	"galleryline/domain"
	"testing"
)

func TestRoles(t *testing.T) {
	if !CanDial(Principal{ID: "x", Role: "guide"}, domain.Extension{ID: "y", Online: true}) {
		t.Fatal("cannot dial")
	}
	if CanViewRecords(Principal{Role: "guide"}) {
		t.Fatal("guide sees records")
	}
}
