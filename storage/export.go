package storage

import (
	"encoding/json"
	"galleryline/domain"
)

type Snapshot struct {
	Extensions []domain.Extension
	Calls      []domain.CallSession
	Records    []domain.CallRecord
}

func (s *Store) Snapshot() (Snapshot, error) {
	e, er := s.ListExtensions()
	if er != nil {
		return Snapshot{}, er
	}
	c, er := s.ListCalls()
	if er != nil {
		return Snapshot{}, er
	}
	r, er := s.ListRecords()
	return Snapshot{e, c, r}, er
}
func (s *Store) Export() ([]byte, error) {
	v, e := s.Snapshot()
	if e != nil {
		return nil, e
	}
	return json.MarshalIndent(v, "", "  ")
}
func (s *Store) Import(data []byte) error {
	var v Snapshot
	if e := json.Unmarshal(data, &v); e != nil {
		return e
	}
	for _, x := range v.Extensions {
		if e := s.SaveExtension(x); e != nil {
			return e
		}
	}
	for _, x := range v.Calls {
		if e := s.SaveCall(x); e != nil {
			return e
		}
	}
	for _, x := range v.Records {
		if e := s.SaveRecord(x); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) Count(kind string) int {
	raw, e := s.List(kind + ":")
	if e != nil {
		return 0
	}
	return len(raw)
}
