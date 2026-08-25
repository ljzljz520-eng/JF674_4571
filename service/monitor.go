package service

import (
	"galleryline/domain"
	"galleryline/storage"
	"sort"
)

type Monitor struct {
	store *storage.Store
	dir   *Directory
}

func NewMonitor(s *storage.Store, d *Directory) *Monitor { return &Monitor{store: s, dir: d} }
func (m *Monitor) Snapshot() ([]domain.Extension, []domain.CallSession, error) {
	e, er := m.dir.List()
	if er != nil {
		return nil, nil, er
	}
	c, er := m.store.ListCalls()
	return e, c, er
}
func (m *Monitor) StatusMap() map[string]bool {
	out := map[string]bool{}
	es, e := m.dir.List()
	if e != nil {
		return out
	}
	for _, v := range es {
		out[v.ID] = v.Online
	}
	return out
}
func (m *Monitor) ActiveCalls() []domain.CallSession {
	all, e := m.store.ListCalls()
	if e != nil {
		return nil
	}
	out := []domain.CallSession{}
	for _, c := range all {
		if c.Status == "ringing" || c.Status == "connected" {
			out = append(out, c)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
func (m *Monitor) Busy(id string) bool {
	for _, c := range m.ActiveCalls() {
		if c.CallerID == id || c.CalleeID == id {
			return true
		}
	}
	return false
}
func (m *Monitor) CountOnline() int {
	n := 0
	for _, v := range m.StatusMap() {
		if v {
			n++
		}
	}
	return n
}
