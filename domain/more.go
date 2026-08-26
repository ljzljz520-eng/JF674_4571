package domain

import (
	"sort"
	"strings"
)

func SortCalls(in []CallSession) []CallSession {
	out := append([]CallSession(nil), in...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].StartedAt.Before(out[j].StartedAt) })
	return out
}
func SortRecords(in []CallRecord) []CallRecord {
	out := append([]CallRecord(nil), in...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].StartedAt.Before(out[j].StartedAt) })
	return out
}
func FilterStatus(in []CallSession, status string) []CallSession {
	out := []CallSession{}
	for _, v := range in {
		if v.Status == status {
			out = append(out, v)
		}
	}
	return out
}
func FilterOutcome(in []CallRecord, outcome string) []CallRecord {
	out := []CallRecord{}
	for _, v := range in {
		if v.Outcome == outcome {
			out = append(out, v)
		}
	}
	return out
}
func NormalizeAction(v string) string {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "dial", "accept", "reject", "end", "presence", "records":
		return strings.ToLower(strings.TrimSpace(v))
	default:
		return "unknown"
	}
}
func ActionRequiresPermission(v string) bool            { return NormalizeAction(v) != "unknown" }
func IsTerminalStatus(v string) bool                    { return v == "ended" || v == "rejected" }
func IsRingingStatus(v string) bool                     { return v == "ringing" }
func IsConnectedStatus(v string) bool                   { return v == "connected" }
func SessionParticipants(c CallSession) []string        { return []string{c.CallerID, c.CalleeID} }
func IncludesParticipant(c CallSession, id string) bool { return c.CallerID == id || c.CalleeID == id }
func OppositeParticipant(c CallSession, id string) string {
	if c.CallerID == id {
		return c.CalleeID
	}
	if c.CalleeID == id {
		return c.CallerID
	}
	return ""
}
func ValidTransition(status, event string) bool { return NewCallPolicy().Can(status, event) }
