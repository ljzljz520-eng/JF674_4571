package api

import (
	"encoding/json"
	"galleryline/auth"
	"net/http"
)

func decode[T any](r *http.Request, v *T) error { return json.NewDecoder(r.Body).Decode(v) }
func principal(r *http.Request) auth.Principal {
	return auth.Principal{ID: r.Header.Get("X-Extension"), Role: auth.NormalizeRole(r.Header.Get("X-Role"))}
}
