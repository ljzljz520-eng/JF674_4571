package domain

import "testing"

func TestModels(t *testing.T) {
	e := Extension{ID: "x", Nickname: "X", Online: true}
	if !e.Available() {
		t.Fatal("unavailable")
	}
	if NormalizeStatus("CONNECTED") != "connected" {
		t.Fatal("status")
	}
}
