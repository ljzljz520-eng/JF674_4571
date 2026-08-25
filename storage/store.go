package storage

import (
	"encoding/json"
	"fmt"
	"galleryline/domain"
	"go.etcd.io/bbolt"
	"sync"
)

var bucket = []byte("galleryline")

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	if err = db.Update(func(tx *bbolt.Tx) error { _, e := tx.CreateBucketIfNotExists(bucket); return e }); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) put(key string, v any) error {
	b, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Put([]byte(key), b) })
}
func (s *Store) get(key string, v any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket).Get([]byte(key))
		if b == nil {
			return domain.ErrNotFound
		}
		return json.Unmarshal(b, v)
	})
}
func key(kind, id string) string { return fmt.Sprintf("%s:%s", kind, id) }
func (s *Store) SaveExtension(v domain.Extension) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.put(key("extension", v.ID), v)
}
func (s *Store) GetExtension(id string) (domain.Extension, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var v domain.Extension
	e := s.get(key("extension", id), &v)
	return v, e
}
func (s *Store) SaveCall(v domain.CallSession) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.put(key("call", v.ID), v)
}
func (s *Store) GetCall(id string) (domain.CallSession, error) {
	var v domain.CallSession
	e := s.get(key("call", id), &v)
	return v, e
}
func (s *Store) SavePermission(v domain.PermissionNotice) error {
	return s.put(key("permission", v.ID), v)
}
func (s *Store) SaveDevice(v domain.Device) error     { return s.put(key("device", v.ID), v) }
func (s *Store) SaveRecord(v domain.CallRecord) error { return s.put(key("record", v.ID), v) }
func (s *Store) List(prefix string) ([][]byte, error) {
	var out [][]byte
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucket).ForEach(func(k, v []byte) error {
			if len(k) >= len(prefix) && string(k[:len(prefix)]) == prefix {
				out = append(out, append([]byte(nil), v...))
			}
			return nil
		})
	})
	return out, e
}
