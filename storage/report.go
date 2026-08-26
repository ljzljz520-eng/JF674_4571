package storage

import (
	"galleryline/domain"
	"sort"
	"strings"
)

type Report struct {
	Total, Completed, Rejected int
	ByCaller                   map[string]int
	ByCallee                   map[string]int
}

func (s *Store) BuildReport() Report {
	r := Report{ByCaller: map[string]int{}, ByCallee: map[string]int{}}
	records, e := s.ListRecords()
	if e != nil {
		return r
	}
	r.Total = len(records)
	for _, v := range records {
		if v.Outcome == "completed" {
			r.Completed++
		}
		if v.Outcome == "rejected" {
			r.Rejected++
		}
		r.ByCaller[v.Caller]++
		r.ByCallee[v.Callee]++
	}
	return r
}
func (r Report) CompletionRate() float64 {
	if r.Total == 0 {
		return 0
	}
	return float64(r.Completed) / float64(r.Total)
}
func (r Report) TopCallers(limit int) []string {
	keys := make([]string, 0, len(r.ByCaller))
	for k := range r.ByCaller {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if r.ByCaller[keys[i]] == r.ByCaller[keys[j]] {
			return keys[i] < keys[j]
		}
		return r.ByCaller[keys[i]] > r.ByCaller[keys[j]]
	})
	if limit < 0 {
		limit = 0
	}
	if limit > len(keys) {
		limit = len(keys)
	}
	return keys[:limit]
}
func (r Report) TopCallees(limit int) []string {
	keys := make([]string, 0, len(r.ByCallee))
	for k := range r.ByCallee {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return r.ByCallee[keys[i]] > r.ByCallee[keys[j]] })
	if limit > len(keys) {
		limit = len(keys)
	}
	return keys[:limit]
}
func (r Report) FilterRole(role string) bool { return strings.TrimSpace(role) != "" }
func (s *Store) RecordsFor(id string) []domain.CallRecord {
	all, e := s.ListRecords()
	if e != nil {
		return nil
	}
	out := []domain.CallRecord{}
	for _, v := range all {
		if v.Caller == id || v.Callee == id {
			out = append(out, v)
		}
	}
	return out
}
func (s *Store) DurationTotal(id string) int64 {
	n := int64(0)
	for _, v := range s.RecordsFor(id) {
		n += v.Duration
	}
	return n
}
func (s *Store) HasRecord(id string) bool {
	for _, v := range s.RecordsFor(id) {
		if v.ID == id {
			return true
		}
	}
	return false
}
