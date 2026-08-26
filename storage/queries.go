package storage

import (
	"encoding/json"
	"galleryline/domain"
	"sort"
)

func (s *Store) ListExtensions() ([]domain.Extension, error) {
	raw, e := s.List("extension:")
	if e != nil {
		return nil, e
	}
	out := make([]domain.Extension, 0, len(raw))
	for _, b := range raw {
		var v domain.Extension
		if json.Unmarshal(b, &v) == nil {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}
func (s *Store) ListCalls() ([]domain.CallSession, error) {
	raw, e := s.List("call:")
	if e != nil {
		return nil, e
	}
	out := []domain.CallSession{}
	for _, b := range raw {
		var v domain.CallSession
		if json.Unmarshal(b, &v) == nil {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].StartedAt.Before(out[j].StartedAt) })
	return out, nil
}
func (s *Store) ListRecords() ([]domain.CallRecord, error) {
	raw, e := s.List("record:")
	if e != nil {
		return nil, e
	}
	out := []domain.CallRecord{}
	for _, b := range raw {
		var v domain.CallRecord
		if json.Unmarshal(b, &v) == nil {
			out = append(out, v)
		}
	}
	return out, nil
}
