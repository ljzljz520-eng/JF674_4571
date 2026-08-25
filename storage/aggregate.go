package storage

import (
	"galleryline/domain"
	"sort"
)

type PresenceSummary struct{ Total, Online, Guides, Desks, Devices int }

func (s *Store) Presence() PresenceSummary {
	v := PresenceSummary{}
	all, e := s.ListExtensions()
	if e != nil {
		return v
	}
	v.Total = len(all)
	for _, x := range all {
		if x.Online {
			v.Online++
		}
		switch x.Role {
		case "guide":
			v.Guides++
		case "desk":
			v.Desks++
		case "device":
			v.Devices++
		}
	}
	return v
}
func (s *Store) CallsByStatus() map[string][]domain.CallSession {
	out := map[string][]domain.CallSession{}
	all, e := s.ListCalls()
	if e != nil {
		return out
	}
	for _, v := range all {
		out[v.Status] = append(out[v.Status], v)
	}
	for k := range out {
		sort.Slice(out[k], func(i, j int) bool { return out[k][i].ID < out[k][j].ID })
	}
	return out
}
func (s *Store) CallsForExtension(id string) []domain.CallSession {
	all, e := s.ListCalls()
	if e != nil {
		return nil
	}
	out := []domain.CallSession{}
	for _, v := range all {
		if v.CallerID == id || v.CalleeID == id {
			out = append(out, v)
		}
	}
	return out
}
func (s *Store) RecentCalls(limit int) []domain.CallSession {
	all, e := s.ListCalls()
	if e != nil {
		return nil
	}
	sort.Slice(all, func(i, j int) bool { return all[i].StartedAt.After(all[j].StartedAt) })
	if limit < 0 {
		limit = 0
	}
	if limit > len(all) {
		limit = len(all)
	}
	return all[:limit]
}
func (s *Store) RecentRecords(limit int) []domain.CallRecord {
	all, e := s.ListRecords()
	if e != nil {
		return nil
	}
	sort.Slice(all, func(i, j int) bool { return all[i].StartedAt.After(all[j].StartedAt) })
	if limit < 0 {
		limit = 0
	}
	if limit > len(all) {
		limit = len(all)
	}
	return all[:limit]
}
