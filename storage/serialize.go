package storage

import (
	"encoding/json"
	"galleryline/domain"
	"go.etcd.io/bbolt"
)

func encode(v any) ([]byte, error) { return json.Marshal(v) }
func decode(b []byte, v any) error { return json.Unmarshal(b, v) }
func (s *Store) PutRaw(kind, id string, v any) error {
	b, e := encode(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Put([]byte(key(kind, id)), b) })
}
func (s *Store) GetRaw(kind, id string, v any) error {
	var b []byte
	e := s.db.View(func(tx *bbolt.Tx) error {
		x := tx.Bucket(bucket).Get([]byte(key(kind, id)))
		if x == nil {
			return domain.ErrNotFound
		}
		b = append([]byte(nil), x...)
		return nil
	})
	if e != nil {
		return e
	}
	return decode(b, v)
}
func (s *Store) ReplaceExtensions(v []domain.Extension) error {
	for _, x := range v {
		if e := s.SaveExtension(x); e != nil {
			return e
		}
	}
	return nil
}
