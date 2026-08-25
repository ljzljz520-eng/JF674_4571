package api

import (
	"fmt"
	"net/http"
	"strings"
)

type CallRequest struct {
	CallerID string `json:"caller_id"`
	CalleeID string `json:"callee_id"`
	CallID   string `json:"call_id"`
}

func validateCallRequest(v CallRequest) error {
	if strings.TrimSpace(v.CallerID) == "" {
		return fmt.Errorf("caller required")
	}
	if strings.TrimSpace(v.CalleeID) == "" {
		return fmt.Errorf("callee required")
	}
	if v.CallerID == v.CalleeID {
		return fmt.Errorf("caller and callee must differ")
	}
	return nil
}
func requestID(r *http.Request) string {
	v := r.Header.Get("X-Request-ID")
	if v == "" {
		v = r.URL.Query().Get("request")
	}
	return strings.TrimSpace(v)
}
func wantsJSON(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "application/json") || r.Header.Get("Accept") == ""
}
func allowOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Role, X-Extension")
}
func methodAllowed(r *http.Request, methods ...string) bool {
	for _, m := range methods {
		if r.Method == m {
			return true
		}
	}
	return false
}
