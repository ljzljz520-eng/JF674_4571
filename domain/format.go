package domain

import (
	"fmt"
	"strings"
)

func FormatExtension(e Extension) string {
	return fmt.Sprintf("%s [%s] %s", e.ID, e.Role, onlineWord(e.Online))
}
func onlineWord(v bool) string {
	if v {
		return "online"
	}
	return "offline"
}
func FormatCall(c CallSession) string {
	return fmt.Sprintf("%s %s %s", c.ID, c.Parties(), c.StatusLabel())
}
func FormatRecord(r CallRecord) string { return fmt.Sprintf("%s %ds", r.Summary(), r.Duration) }
func FormatPermission(n PermissionNotice) string {
	state := "pending"
	if n.Granted {
		state = "granted"
	}
	return n.Summary() + " (" + state + ")"
}
func ParseRole(v string) string {
	v = strings.ToLower(strings.TrimSpace(v))
	if AllowedRole(v, "dial") || v == "admin" {
		return v
	}
	return "guest"
}
func ParseOutcome(v string) string {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "completed", "rejected", "missed":
		return strings.ToLower(strings.TrimSpace(v))
	default:
		return "unknown"
	}
}
func IsValidOutcome(v string) bool { return ParseOutcome(v) != "unknown" }
func IsValidRole(v string) bool    { return ParseRole(v) != "guest" }
func IsValidKind(v string) bool {
	switch strings.ToLower(v) {
	case "offer", "answer", "hangup", "candidate":
		return true
	default:
		return false
	}
}
func NormalizeID(v string) string { return strings.ToLower(strings.TrimSpace(v)) }
func SameParty(a, b string) bool  { return NormalizeID(a) == NormalizeID(b) }
