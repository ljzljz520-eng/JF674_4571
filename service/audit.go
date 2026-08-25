package service

import (
	"fmt"
	"galleryline/domain"
	"galleryline/storage"
	"sort"
	"strings"
)

type AuditEntry struct{ Action, Actor, Target, Detail string }
type AuditLog struct {
	store   *storage.Store
	entries []AuditEntry
}

func NewAuditLog(s *storage.Store) *AuditLog { return &AuditLog{store: s, entries: []AuditEntry{}} }
func (a *AuditLog) Record(actor, action, target, detail string) AuditEntry {
	e := AuditEntry{Action: action, Actor: actor, Target: target, Detail: detail}
	a.entries = append(a.entries, e)
	return e
}
func (a *AuditLog) All() []AuditEntry { return append([]AuditEntry(nil), a.entries...) }
func (a *AuditLog) Filter(actor string) []AuditEntry {
	out := []AuditEntry{}
	for _, e := range a.entries {
		if actor == "" || e.Actor == actor {
			out = append(out, e)
		}
	}
	return out
}
func (a *AuditLog) Summary() map[string]int {
	m := map[string]int{}
	for _, e := range a.entries {
		m[e.Action]++
	}
	return m
}
func (a *AuditLog) Describe(e AuditEntry) string {
	return fmt.Sprintf("%s:%s:%s:%s", e.Actor, e.Action, e.Target, e.Detail)
}
func (a *AuditLog) Search(term string) []AuditEntry {
	term = strings.ToLower(term)
	out := []AuditEntry{}
	for _, e := range a.entries {
		if strings.Contains(strings.ToLower(a.Describe(e)), term) {
			out = append(out, e)
		}
	}
	return out
}
func (a *AuditLog) Sorted() []AuditEntry {
	out := a.All()
	sort.SliceStable(out, func(i, j int) bool { return out[i].Action < out[j].Action })
	return out
}
func (a *AuditLog) AttachCall(c domain.CallSession) AuditEntry {
	return a.Record(c.CallerID, "call", c.ID, c.Status)
}
func (a *AuditLog) AttachExtension(e domain.Extension) AuditEntry {
	return a.Record(e.ID, "presence", e.ID, fmt.Sprint(e.Online))
}
