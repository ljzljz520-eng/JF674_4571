package storage

import (
	"encoding/json"
	"fmt"
	"galleryline/domain"
	"sort"
)

func (s *Store) EnsureExtension(e domain.Extension) error {
	if e.Validate() != nil {
		return domain.ErrNotFound
	}
	if _, er := s.GetExtension(e.ID); er == nil {
		return nil
	}
	return s.SaveExtension(e)
}
func (s *Store) EnsureDevice(d domain.Device) error {
	if d.Validate() != nil {
		return domain.ErrNotFound
	}
	if _, er := s.GetDevice(d.ID); er == nil {
		return nil
	}
	return s.SaveDevice(d)
}
func (s *Store) ValidateCall(id string) error {
	c, e := s.GetCall(id)
	if e != nil {
		return e
	}
	return c.Validate()
}
func (s *Store) Require(kind, id string) error {
	switch kind {
	case "extension":
		_, e := s.GetExtension(id)
		return e
	case "call":
		_, e := s.GetCall(id)
		return e
	case "device":
		_, e := s.GetDevice(id)
		return e
	default:
		return fmt.Errorf("unknown entity")
	}
}
func (s *Store) CountOnline() int {
	n := 0
	all, e := s.ListExtensions()
	if e != nil {
		return 0
	}
	for _, v := range all {
		if v.Online {
			n++
		}
	}
	return n
}
func (s *Store) CountActiveCalls() int {
	n := 0
	all, e := s.ListCalls()
	if e != nil {
		return 0
	}
	for _, v := range all {
		if !v.IsFinished() {
			n++
		}
	}
	return n
}
func (s *Store) IDs(kind string) []string {
	raw, e := s.List(kind + ":")
	if e != nil {
		return nil
	}
	out := []string{}
	for _, b := range raw {
		var v map[string]any
		if json.Unmarshal(b, &v) == nil {
			if id, ok := v["id"].(string); ok {
				out = append(out, id)
			}
		}
	}
	sort.Strings(out)
	return out
}
