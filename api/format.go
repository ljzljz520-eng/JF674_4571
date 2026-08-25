package api

import (
	"encoding/json"
	"fmt"
	"galleryline/domain"
	"net/http"
)

type Envelope struct {
	Data  any    `json:"data"`
	Error string `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Envelope{Data: v})
}
func writeError(w http.ResponseWriter, status int, e error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Envelope{Error: e.Error()})
}
func statusForError(e error) int {
	switch e {
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrUnauthorized:
		return http.StatusForbidden
	case domain.ErrOffline:
		return http.StatusConflict
	default:
		return http.StatusBadRequest
	}
}
func extensionJSON(e domain.Extension) map[string]any {
	return map[string]any{"id": e.ID, "nickname": e.Nickname, "role": e.Role, "online": e.Online, "label": e.Label()}
}
func callJSON(c domain.CallSession) map[string]any {
	return map[string]any{"id": c.ID, "caller": c.CallerID, "callee": c.CalleeID, "status": c.Status, "parties": c.Parties()}
}
func recordJSON(r domain.CallRecord) map[string]any {
	return map[string]any{"id": r.ID, "caller": r.Caller, "callee": r.Callee, "outcome": r.Outcome, "duration": r.Duration}
}
func parseID(r *http.Request) string {
	if v := r.URL.Query().Get("id"); v != "" {
		return v
	}
	return r.Header.Get("X-Request-ID")
}
func requireMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		writeError(w, http.StatusMethodNotAllowed, fmt.Errorf("method %s required", method))
		return false
	}
	return true
}
