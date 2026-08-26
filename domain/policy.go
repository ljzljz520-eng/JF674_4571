package domain

import (
	"fmt"
	"sort"
	"strings"
)

type Transition struct{ From, Event, To string }
type CallPolicy struct{ transitions map[string]Transition }

func NewCallPolicy() *CallPolicy {
	return &CallPolicy{transitions: map[string]Transition{
		"ringing:accept": {From: "ringing", Event: "accept", To: "connected"}, "ringing:reject": {From: "ringing", Event: "reject", To: "rejected"}, "ringing:end": {From: "ringing", Event: "end", To: "ended"}, "connected:end": {From: "connected", Event: "end", To: "ended"},
	}}
}
func (p *CallPolicy) Transition(status, event string) (string, error) {
	k := status + ":" + event
	v, ok := p.transitions[k]
	if !ok {
		return status, ErrInvalidTransition
	}
	return v.To, nil
}
func (p *CallPolicy) Events(status string) []string {
	out := []string{}
	for _, v := range p.transitions {
		if v.From == status {
			out = append(out, v.Event)
		}
	}
	sort.Strings(out)
	return out
}
func (p *CallPolicy) Can(status, event string) bool {
	_, e := p.Transition(status, event)
	return e == nil
}
func (p *CallPolicy) Explain(status, event string) string {
	to, e := p.Transition(status, event)
	if e != nil {
		return e.Error()
	}
	return fmt.Sprintf("%s -> %s", status, to)
}
func RoleMatrix() map[string][]string {
	return map[string][]string{"admin": {"dial", "accept", "reject", "end", "records", "presence"}, "guide": {"dial", "end"}, "desk": {"dial", "accept", "reject", "end"}, "device": {"dial", "accept", "end"}}
}
func AllowedRole(role, action string) bool {
	for _, v := range RoleMatrix()[role] {
		if v == action {
			return true
		}
	}
	return false
}
func AllRoles() []string {
	m := RoleMatrix()
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
func DescribeRole(role string) string {
	if !AllowedRole(role, "dial") {
		return "restricted"
	}
	return strings.Title(role) + " can place calls"
}
func SanitizeNickname(v string) string { return strings.TrimSpace(strings.ReplaceAll(v, "\n", " ")) }
func ExtensionSort(in []Extension) []Extension {
	out := append([]Extension(nil), in...)
	sort.Slice(out, func(i, j int) bool { return SanitizeNickname(out[i].Nickname) < SanitizeNickname(out[j].Nickname) })
	return out
}
func FindOnline(in []Extension) []Extension {
	out := []Extension{}
	for _, v := range in {
		if v.Available() {
			out = append(out, v)
		}
	}
	return ExtensionSort(out)
}
func MatchRole(in []Extension, role string) []Extension {
	out := []Extension{}
	for _, v := range in {
		if v.Role == role {
			out = append(out, v)
		}
	}
	return ExtensionSort(out)
}
func UniqueIDs(in []Extension) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, v := range in {
		if !seen[v.ID] {
			seen[v.ID] = true
			out = append(out, v.ID)
		}
	}
	sort.Strings(out)
	return out
}
func ContainsID(in []Extension, id string) bool {
	for _, v := range in {
		if v.ID == id {
			return true
		}
	}
	return false
}
func CountRole(in []Extension, role string) int {
	n := 0
	for _, v := range in {
		if v.Role == role {
			n++
		}
	}
	return n
}
func CountOnlineRole(in []Extension, role string) int {
	n := 0
	for _, v := range in {
		if v.Role == role && v.Online {
			n++
		}
	}
	return n
}
func StatusCounts(in []CallSession) map[string]int {
	m := map[string]int{}
	for _, v := range in {
		m[v.Status]++
	}
	return m
}
func FinishedCalls(in []CallSession) []CallSession {
	out := []CallSession{}
	for _, v := range in {
		if v.IsFinished() {
			out = append(out, v)
		}
	}
	return out
}
func ActiveCalls(in []CallSession) []CallSession {
	out := []CallSession{}
	for _, v := range in {
		if !v.IsFinished() {
			out = append(out, v)
		}
	}
	return out
}
func CallsBetween(in []CallSession, a, b string) []CallSession {
	out := []CallSession{}
	for _, v := range in {
		if (v.CallerID == a && v.CalleeID == b) || (v.CallerID == b && v.CalleeID == a) {
			out = append(out, v)
		}
	}
	return out
}
func HasActiveBetween(in []CallSession, a, b string) bool {
	for _, v := range ActiveCalls(in) {
		if (v.CallerID == a && v.CalleeID == b) || (v.CallerID == b && v.CalleeID == a) {
			return true
		}
	}
	return false
}
func RecordTotal(in []CallRecord) int64 {
	n := int64(0)
	for _, v := range in {
		n += v.Duration
	}
	return n
}
func RecordOutcomes(in []CallRecord) map[string]int {
	m := map[string]int{}
	for _, v := range in {
		m[v.Outcome]++
	}
	return m
}
func RecordsFor(in []CallRecord, id string) []CallRecord {
	out := []CallRecord{}
	for _, v := range in {
		if v.Caller == id || v.Callee == id {
			out = append(out, v)
		}
	}
	return out
}
