package storage

import "galleryline/domain"

func (s *Store) GetDevice(id string) (domain.Device, error) {
	var v domain.Device
	e := s.get(key("device", id), &v)
	return v, e
}
