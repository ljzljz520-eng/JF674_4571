package service

import (
	"galleryline/domain"
	"sort"
)

type PresenceEvent struct {
	ExtensionID string
	Online      bool
	Sequence    int
}
type PresenceTracker struct {
	current map[string]bool
	events  []PresenceEvent
}

func NewPresenceTracker() *PresenceTracker {
	return &PresenceTracker{current: map[string]bool{}, events: []PresenceEvent{}}
}
func (p *PresenceTracker) Set(id string, online bool) {
	p.current[id] = online
	p.events = append(p.events, PresenceEvent{ExtensionID: id, Online: online, Sequence: len(p.events) + 1})
}
func (p *PresenceTracker) Get(id string) bool { return p.current[id] }
func (p *PresenceTracker) Online() []string {
	out := []string{}
	for id, v := range p.current {
		if v {
			out = append(out, id)
		}
	}
	sort.Strings(out)
	return out
}
func (p *PresenceTracker) Offline() []string {
	out := []string{}
	for id, v := range p.current {
		if !v {
			out = append(out, id)
		}
	}
	sort.Strings(out)
	return out
}
func (p *PresenceTracker) Events() []PresenceEvent { return append([]PresenceEvent(nil), p.events...) }
func (p *PresenceTracker) Count() int              { return len(p.current) }
func (p *PresenceTracker) ActiveCount() int        { return len(p.Online()) }
func (p *PresenceTracker) Apply(e PresenceEvent) {
	p.current[e.ExtensionID] = e.Online
	p.events = append(p.events, e)
}
func (p *PresenceTracker) Merge(events []PresenceEvent) {
	for _, e := range events {
		p.Apply(e)
	}
}
func (p *PresenceTracker) Extensions() []domain.Extension {
	out := []domain.Extension{}
	for id, on := range p.current {
		out = append(out, domain.Extension{ID: id, Online: on})
	}
	return domain.ExtensionSort(out)
}
