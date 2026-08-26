package storage

import (
	"galleryline/domain"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) TouchExtension(id string, at time.Time) error {
	e, er := s.GetExtension(id)
	if er != nil {
		return er
	}
	e.LastSeen = at
	return s.SaveExtension(e)
}
func (s *Store) UpdateCallStatus(id, status string, at time.Time) error {
	c, er := s.GetCall(id)
	if er != nil {
		return er
	}
	c.Status = domain.NormalizeStatus(status)
	if c.Status == "connected" {
		c.ConnectedAt = at
	}
	if c.Status == "ended" {
		c.EndedAt = at
	}
	return s.SaveCall(c)
}
func (s *Store) Delete(kind, id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Delete([]byte(key(kind, id))) })
}
